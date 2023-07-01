package storagecore

import (
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/storage"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Manager struct {
	DB         *PostgresDB
	AwsSession *session.Session
}

func NewStorageManager(options config.Options) storage.StorageManager {
	db := NewPostgresDB(options)

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(options.AwsRegion)},
	)
	if err != nil {
		log.Fatalf("Can't create AWS Session: %v", err)
	}

	return &Manager{
		DB:         db,
		AwsSession: awsSession,
	}
}

func (m *Manager) GetDB() storage.DB {
	return m.DB
}

func (m *Manager) GetBlobStorage(bucketName string) storage.BlobStorage {
	return &AwsBlobStorage{
		s3Client:   s3.New(m.AwsSession),
		bucketName: bucketName,
	}
}
