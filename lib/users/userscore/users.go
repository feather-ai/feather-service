package userscore

import (
	"context"
	"encoding/json"
	"errors"
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/storage"
	"feather-ai/service-core/lib/users"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gosimple/slug"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var gUnauthorizedError = errors.New("Unauthorized")
var gExpiredError = errors.New("Expired")

type Auth0IdentityResponseDto struct {
	Id       string `json:"sub"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Verified bool   `json:"email_verified"`
}

type Auth0Dto struct {
	Auth0Id   string    `db:"auth0_id"`
	FeatherId uuid.UUID `db:"feather_id"`
}

type usersCoreManager struct {
	options    config.Options
	httpClient *resty.Client
	DB         storage.DB
}

func New(manager storage.StorageManager, opts config.Options) users.UserManager {
	prepareDBQueries(manager.GetDB())

	return &usersCoreManager{
		options:    opts,
		httpClient: resty.New(),
		DB:         manager.GetDB(),
	}
}

/*
* Dev only function to perform a debug login of a user - ie. login without authenticating with the identity provider
 */
func (m *usersCoreManager) debugUserLogin(ctx context.Context, accessToken string) *Auth0IdentityResponseDto {
	id := strings.TrimPrefix(accessToken, "debug ")

	info := Auth0IdentityResponseDto{
		Id:       id,
		Email:    id,
		Name:     id,
		Nickname: id,
		Verified: true,
	}
	return &info
}

/*
* Call the auth0 userinfo endpoint
 */
func (m *usersCoreManager) auth0UserInfo(ctx context.Context, accessToken string) (*Auth0IdentityResponseDto, error) {
	log := logrus.WithContext(ctx)

	if m.options.DebugUser {
		if strings.HasPrefix(accessToken, "debug ") {
			log.Infof("Login: DebugLogin for user: %s", accessToken)
			return m.debugUserLogin(ctx, accessToken), nil
		}
	}

	// Query Auth0 /userinfo Endpoint
	url := m.options.Auth0IdentityURL
	response, err := m.httpClient.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get(url)

	if err != nil {
		log.Errorf("Login: Netlify error: %v", err)
		return nil, gUnauthorizedError
	}
	if response.StatusCode() != 200 {
		log.Errorf("Login: Netlify error code: %v, %v", response.StatusCode(), response.String())
		return nil, gUnauthorizedError
	}

	info := Auth0IdentityResponseDto{}
	if err = json.Unmarshal(response.Body(), &info); err != nil {
		log.Errorf("Login: Cannot parse netlify response: %v", err)
		return nil, gUnauthorizedError
	}

	return &info, nil
}

func (m *usersCoreManager) LoginAuthenticate(ctx context.Context, token string) (*users.EncryptedFeatherToken, error) {
	log := logrus.WithContext(ctx)
	info, err := m.auth0UserInfo(ctx, token)
	if err != nil {
		return nil, err
	}
	log.Printf("Login: Auth0 ID=%s, email=%s", info.Id, info.Email)

	// Check if we have a record for this user

	auth0Record := []Auth0Dto{}
	args := map[string]interface{}{"auth0_id": info.Id}
	err = gQueryAuth0UserIDStmt.QueryMany(ctx, args, &auth0Record)
	if err != nil {
		log.Printf("Login: Failed to lookup auth0 table: %v\n", err)
		return nil, err
	}

	var userId uuid.UUID

	if len(auth0Record) == 0 {
		userId = uuid.NewV4()
		// Create a new user
		newAuth0Record := Auth0Dto{
			Auth0Id:   info.Id,
			FeatherId: userId,
		}

		tx, err := m.DB.NewTransaction(ctx)
		if err != nil {
			log.Errorf("Login: Failed to create DB Transaction: %v", err)
			return nil, err
		}
		// In a transaction, insert entry in "users" table and the mapping from auth0 to featherid in "auth0" table
		tx.Execute(ctx, gInsertAuth0UserStmt, &newAuth0Record)
		insertUserArgs := map[string]interface{}{
			"user_id": userId,
			"name":    slug.Make(info.Name),
		}
		tx.Execute(ctx, gInsertUserStmt, insertUserArgs)
		tx.Commit(ctx)
		if tx.Error() != nil {
			log.Errorf("Login: Failed to update user tables: %v", tx.Error())
			return nil, tx.Error()
		}
	} else {
		userId = auth0Record[0].FeatherId
	}

	log.Printf("Login: Mapped Auth0 id %s to feather userId: %s", info.Id, userId.String())

	// Generate a Feather Token
	featherToken := GenerateEncryptedFeatherToken(userId)
	if featherToken == nil {
		return nil, gUnauthorizedError
	}
	return featherToken, nil
}

func (m *usersCoreManager) ExtendFeatherToen(ctx context.Context, token users.EncryptedFeatherToken) (*users.EncryptedFeatherToken, error) {
	log := logrus.WithContext(ctx)
	decrypted := DecryptFeatherToken(token)
	if decrypted == nil {
		log.Info("ExtendedFeatherToken: Cannot decrypt")
		return nil, gUnauthorizedError
	}

	if decrypted.Secret != gFeatherTokenSecret {
		log.Info("ExtendedFeatherToken: Secret mismatch")
		return nil, gUnauthorizedError
	}

	if decrypted.Expiry.After(time.Now().UTC()) {
		log.Infof("ExtendedFeatherToken: Attempt to extend token that hasn't expired for userId %v", decrypted.UserID.String())
		return nil, gUnauthorizedError
	}
	// If user hasn't logged in for 48 hours (TODO: configure), force a re-login
	if decrypted.Expiry.Before(time.Now().UTC().Add(-time.Hour * 24 * 2)) {
		log.Infof("ExtendedFeatherToken: Attempt to extend token that is too old %v", decrypted.UserID.String())
		return nil, gUnauthorizedError
	}

	// :TODO: Ideally, we re-validate with Netlify here

	newFeatherToken := GenerateEncryptedFeatherToken(decrypted.UserID)
	if newFeatherToken == nil {
		return nil, gUnauthorizedError
	}

	return newFeatherToken, nil
}

func (m *usersCoreManager) ValidateFeatherToken(ctx context.Context, token users.EncryptedFeatherToken) (*users.FeatherTokenData, error) {
	log := logrus.WithContext(ctx)
	decrypted := DecryptFeatherToken(token)
	if decrypted == nil {
		return nil, gUnauthorizedError
	}

	if decrypted.Expiry.Before(time.Now().UTC()) {
		log.Infof("Token expired - userId=%v", decrypted.UserID.String())
		return nil, gExpiredError
	}

	if decrypted.Secret != gFeatherTokenSecret {
		return nil, gUnauthorizedError
	}

	return decrypted, nil
}

func (m *usersCoreManager) GetUserByName(ctx context.Context, name string) (*users.UserDTO, error) {
	result := []users.UserDTO{}

	args := map[string]interface{}{"name": name}
	err := gQueryUserByNameStmt.QueryMany(ctx, args, &result)
	if err != nil {
		log.Printf("GetUserByName: Failed to query: %v\n", err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (m *usersCoreManager) GetUserByID(ctx context.Context, userId uuid.UUID) (*users.UserDTO, error) {
	result := []users.UserDTO{}

	args := map[string]interface{}{"user_id": userId.String()}
	err := gQueryUserByIDStmt.QueryMany(ctx, args, &result)
	if err != nil {
		log.Printf("GetUserByID: Failed to query: %v\n", err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (m *usersCoreManager) LookUpUsersByIDs(ctx context.Context, userIds []uuid.UUID) (map[string]*users.UserDTO, error) {
	uniqueIds := map[string]*users.UserDTO{}
	for _, id := range userIds {
		uniqueIds[id.String()] = nil
	}

	// Build the query string to fetch all users at once
	sb := strings.Builder{}
	sb.WriteString(`SELECT * FROM users WHERE user_id in (`)
	first := true
	for k, _ := range uniqueIds {
		if first == false {
			sb.WriteRune(',')
		}
		first = false
		sb.WriteRune('\'')
		sb.WriteString(k)
		sb.WriteRune('\'')
	}
	sb.WriteRune(')')

	queryString := sb.String()

	results := make([]users.UserDTO, 0, len(uniqueIds))
	err := m.DB.DynamicQuery(ctx, queryString, &results)
	if err != nil {
		return nil, err
	}

	for _, item := range results {
		captureVar := item // Need to copy the iterator else, we will capture a changing var
		uniqueIds[item.UserID.String()] = &captureVar
	}

	return uniqueIds, nil
}

func (m *usersCoreManager) CreateAPIKey(ctx context.Context, userId uuid.UUID, name string) (uuid.UUID, error) {
	keyID := uuid.NewV4()
	dto := users.ApiKeyDTO{
		ApiKey:  keyID.String(),
		UserID:  userId.String(),
		KeyName: name,
		Created: time.Now().UTC(),
		Revoked: false,
	}

	err := gInsertApiKeyStmt.Write(ctx, &dto)
	if err != nil {
		return uuid.Nil, err
	}

	return keyID, nil
}

func (m *usersCoreManager) DeleteAPIKey(ctx context.Context, userId uuid.UUID, key uuid.UUID) error {
	args := map[string]interface{}{"api_key": key.String()}
	err := gDeleteApiKeyStmt.Write(ctx, args)
	if err != nil {
		log.Printf("DeleteAPIKey: db error: %v\n", err)
		return err
	}
	return nil
}

func (m *usersCoreManager) GetAllAPIKeys(ctx context.Context, userId uuid.UUID) ([]users.ApiKeyDTO, error) {
	result := []users.ApiKeyDTO{}

	args := map[string]interface{}{"user_id": userId.String()}
	err := gQueryGetAllApiKeyStmt.QueryMany(ctx, args, &result)
	if err != nil {
		log.Printf("GetAllAPIKeys: Failed to query for api key: %v\n", err)
		return nil, err
	}

	return result, nil
}

func (m *usersCoreManager) ValidateAPIKey(ctx context.Context, key uuid.UUID) (uuid.UUID, error) {
	result := []users.ApiKeyDTO{}

	args := map[string]interface{}{"api_key": key.String()}
	err := gQueryGetApiKeyStmt.QueryMany(ctx, args, &result)
	if err != nil {
		log.Printf("Failed to query for api key: %v\n", err)
		return uuid.Nil, err
	}

	if len(result) != 1 || result[0].Revoked {
		return uuid.Nil, errors.New("Revoked")
	}

	return uuid.FromString(result[0].UserID)
}
