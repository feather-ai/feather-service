package storage

import (
	"context"
	"time"
)

type StorageManager interface {
	GetDB() DB
	GetBlobStorage(bucketName string) BlobStorage
}

type BlobStorage interface {
	GenerateUploadURL(ctx context.Context, name string, expiry time.Duration) (string, error)
	GetFilesInFolder(ctx context.Context, folder string) ([]string, error)
	GetRootPath() string
}

type PreparedStatement interface {
	Write(ctx context.Context, args interface{}) error
	WriteReturnID(ctx context.Context, args interface{}) (int64, error)

	QueryMany(ctx context.Context, args interface{}, results interface{}) error
}

type DBTransaction interface {
	Execute(ctx context.Context, statement PreparedStatement, args interface{})
	ExecuteMany(ctx context.Context, statement PreparedStatement, args []interface{})
	Commit(ctx context.Context)
	Rollback(ctx context.Context)
	Error() error
}

type DB interface {
	PrepareQuery(query string) PreparedStatement

	NewTransaction(ctx context.Context) (DBTransaction, error)
	DynamicQuery(ctx context.Context, query string, results interface{}) error
}
