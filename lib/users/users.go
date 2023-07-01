package users

import (
	"context"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

var ErrorInvalidKey = errors.New("Invalid API Key")

type UserDTO struct {
	UserID uuid.UUID `db:"user_id"`
	Name   string    `db:"name"`
}

type ApiKeyDTO struct {
	ApiKey  string    `db:"api_key"`
	UserID  string    `db:"user_id"`
	KeyName string    `db:"key_name"`
	Created time.Time `db:"created"`
	Revoked bool      `db:"revoked"`
}

type EncryptedFeatherToken struct {
	Payload string
	Expiry  time.Time
}

type FeatherTokenData struct {
	UserID uuid.UUID `json:"userId"`
	Expiry time.Time `json:"expiry"`
	Secret string    `json:"secret"`
}

type UserManager interface {
	// Called to 'login' a user - we take a netifly token, validate it and
	// generate a feather token if successful. From this point on, we use feather token to access
	// out APIs
	LoginAuthenticate(ctx context.Context, netlifyToken string) (*EncryptedFeatherToken, error)

	// Given a feather token that is valid but expired, extend the token (generate a new one)
	ExtendFeatherToen(ctx context.Context, token EncryptedFeatherToken) (*EncryptedFeatherToken, error)

	// Validate that a provided feather token is valid and not expired
	ValidateFeatherToken(ctx context.Context, token EncryptedFeatherToken) (*FeatherTokenData, error)

	GetUserByName(ctx context.Context, name string) (*UserDTO, error)
	GetUserByID(ctx context.Context, userId uuid.UUID) (*UserDTO, error)
	LookUpUsersByIDs(ctx context.Context, userIds []uuid.UUID) (map[string]*UserDTO, error)

	CreateAPIKey(ctx context.Context, userId uuid.UUID, name string) (uuid.UUID, error)
	DeleteAPIKey(ctx context.Context, userId uuid.UUID, key uuid.UUID) error
	GetAllAPIKeys(ctx context.Context, userId uuid.UUID) ([]ApiKeyDTO, error)
	ValidateAPIKey(ctx context.Context, key uuid.UUID) (uuid.UUID, error)
}
