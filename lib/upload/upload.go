package upload

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UploadRequestDTO struct {
	ID                  uuid.UUID `db:"id"`
	UserID              string    `db:"user_id"`
	SystemID            string    `db:"system_id"`
	VersionTag          string    `db:"version_tag"`
	CreateTime          time.Time `db:"create_time"`
	ExpireTime          time.Time `db:"expire_time"`
	CodeFiles           string    `db:"code_files"`
	ModelFiles          string    `db:"model_files"`
	CodeFilesSignedURL  string    `db:"code_files_signed_url"`
	ModelFilesSignedURL string    `db:"model_files_signed_url"`
	Schema              string    `db:"schema"`
}

type UploadManager interface {
	//  Given a system name (and a user ID in ctx), this function generates a unique
	// time limited signed URL that can be used to upload the files comprising the system to cloud storage.
	// Once started, the upload needs to be Completed - if not anything written gets deleted
	CreateUploadRequest(ctx context.Context, systemId uuid.UUID, versionTag string,
		codeFiles []string, modelFiles []string, schema string) (*UploadRequestDTO, error)

	// Complete a previously started upload.
	TryCompleteUpload(ctx context.Context, reqId uuid.UUID) (*UploadRequestDTO, error)

	// Get the list of all upload requests for the specified user
	GetUploadRequestsForUser(ctx context.Context, userId uuid.UUID) []UploadRequestDTO
}
