package api

import (
	"context"
	"errors"
	"feather-ai/service-core/api/generated/models"
	"feather-ai/service-core/api/generated/restapi/operations"
	"time"

	"feather-ai/service-core/lib/aisystems"
	"feather-ai/service-core/lib/gatekeeper"
	"feather-ai/service-core/lib/upload"
	"feather-ai/service-core/lib/users"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var gUploadManager upload.UploadManager
var gSystemsManager aisystems.AISystemsManager
var gUserManager users.UserManager
var gSpec *spec.Swagger
var gAnonymousGateKeeper gatekeeper.AnonGateKeeper

/*
* Install all the handlers needed to manage uploads of models
 */
func InstallCoreHandlers(api *operations.CoreAPI,
	spec *spec.Swagger,
	uploadManager upload.UploadManager,
	userManager users.UserManager,
	systemsManager aisystems.AISystemsManager,
	anonGateKeeper gatekeeper.AnonGateKeeper) {

	gUploadManager = uploadManager
	gSystemsManager = systemsManager
	gUserManager = userManager
	gAnonymousGateKeeper = anonGateKeeper
	gSpec = spec

	api.APIKeyAuthAuth = AuthAPIKey
	api.FeatherTokenAuth = AuthFeatherToken

	api.PutV1ClientLoginHandler = operations.PutV1ClientLoginHandlerFunc(LoginHandler)
	api.PutV1ClientRefreshHandler = operations.PutV1ClientRefreshHandlerFunc(RefreshHandler)
	api.GetV1HealthHandler = operations.GetV1HealthHandlerFunc(Health)

	api.PutV1ClientApikeyHandler = operations.PutV1ClientApikeyHandlerFunc(CreateAPIKey)
	api.GetV1ClientApikeyHandler = operations.GetV1ClientApikeyHandlerFunc(ListAPIKeys)
	api.DeleteV1ClientApikeyHandler = operations.DeleteV1ClientApikeyHandlerFunc(DeleteAPIKey)
}

func Health(params operations.GetV1HealthParams) middleware.Responder {
	return operations.NewGetV1HealthOK()
}

// API-KEY Auth handler for endpoints protected by 'ApiKeyAuth' - ie. Client SDK
func AuthAPIKey(token string) (interface{}, error) {
	key, err := uuid.FromString(token)
	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	userId, err := gUserManager.ValidateAPIKey(ctx, key)

	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	return userId, nil
}

// API-KEY Auth handler for endpoints protected by 'FeatherToken' - Ie browser
func AuthFeatherToken(encryptedToken string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	token, err := gUserManager.ValidateFeatherToken(ctx, users.EncryptedFeatherToken{
		Payload: encryptedToken,
	})
	if err != nil {
		return nil, err
	}

	return token.UserID, nil
}

/*
* Login the user by performing Netlify authentication - we return a feather token which the
* user can use for all subsequent API calls with feather
 */
func LoginHandler(params operations.PutV1ClientLoginParams) middleware.Responder {
	info, err := gUserManager.LoginAuthenticate(params.HTTPRequest.Context(), params.XAUTH0TOKEN)
	if err != nil {
		return operations.NewPutV1ClientLoginForbidden()
	}

	expiry := strfmt.DateTime(info.Expiry)
	return operations.NewPutV1ClientLoginOK().WithPayload(&models.LoginResponse{
		FeatherToken: info.Payload,
		ExpireAt:     expiry,
	})
}

/*
* Refresh a Feather Token
 */
func RefreshHandler(params operations.PutV1ClientRefreshParams) middleware.Responder {
	encryptedToken := params.XFEATHERTOKEN

	// :TODO: Revalidate with Auth0

	newToken, err := gUserManager.ExtendFeatherToen(params.HTTPRequest.Context(), users.EncryptedFeatherToken{
		Payload: encryptedToken,
	})

	if err != nil {
		return operations.NewPutV1ClientRefreshForbidden()
	}

	expiry := strfmt.DateTime(newToken.Expiry)
	return operations.NewPutV1ClientRefreshOK().WithPayload(&models.LoginResponse{
		FeatherToken: newToken.Payload,
		ExpireAt:     expiry,
	})
}

/*
* Create a new API key for this user
 */
func CreateAPIKey(params operations.PutV1ClientApikeyParams, authenticatedUserId interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userId := authenticatedUserId.(uuid.UUID)

	key, err := gUserManager.CreateAPIKey(ctx, userId, params.Name)
	if err != nil {
		logrus.WithContext(ctx).Infof("CreateApiKey error: %v", err)
		return operations.NewPutV1ClientApikeyBadRequest()
	}
	return operations.NewPutV1ClientApikeyOK().WithPayload(key.String())
}

/*
* Delete an API key for this user
 */
func DeleteAPIKey(params operations.DeleteV1ClientApikeyParams, authenticatedUserId interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userId := authenticatedUserId.(uuid.UUID)
	keyId := uuid.FromStringOrNil(params.Key.String())
	if gUserManager.DeleteAPIKey(ctx, userId, keyId) != nil {
		return operations.NewDeleteV1ClientApikeyBadRequest()
	}
	return operations.NewDeleteV1ClientApikeyOK()
}

/*
* List all the API keys for this user
 */
func ListAPIKeys(params operations.GetV1ClientApikeyParams, authenticatedUserId interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userId := authenticatedUserId.(uuid.UUID)

	keys, err := gUserManager.GetAllAPIKeys(ctx, userId)
	if err != nil {
		logrus.WithContext(ctx).Infof("ListAPIKeys error: %v", err)
		return operations.NewGetV1ClientApikeyBadRequest()
	}

	ret := make([]*operations.GetV1ClientApikeyOKBodyItems0, 0, len(keys))
	for _, v := range keys {
		dt := strfmt.DateTime(v.Created)
		id := strfmt.UUID(v.ApiKey)
		item := operations.GetV1ClientApikeyOKBodyItems0{
			Created: dt,
			Key:     id,
			Name:    v.KeyName,
		}
		ret = append(ret, &item)
	}
	return operations.NewGetV1ClientApikeyOK().WithPayload(ret)
}
