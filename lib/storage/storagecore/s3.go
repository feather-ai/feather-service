package storagecore

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsBlobStorage struct {
	s3Client   *s3.S3
	bucketName string
}

func (m *AwsBlobStorage) GetRootPath() string {
	return m.bucketName
}

func (m *AwsBlobStorage) GenerateUploadURL(
	ctx context.Context, name string, expiry time.Duration) (string, error) {

	req, _ := m.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(m.bucketName),
		Key:    aws.String(name),
	})
	return req.Presign(expiry)
}

// Given a 'folder' in S3, list all the 'files' in it.
// Folder is full path from bucket root. Eg. <userid>/<systemid>/<version>/
func (m *AwsBlobStorage) GetFilesInFolder(ctx context.Context, folder string) ([]string, error) {

	response, err := m.s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(m.bucketName),
		Prefix: aws.String(folder),
	})

	if err != nil {
		return nil, err
	}

	// S3 will return a list like this:
	//
	// 	00000000-0000-0000-0000-000000000000/b65ad795-5c8b-401a-91aa-fe7b599d3db3/
	//	00000000-0000-0000-0000-000000000000/b65ad795-5c8b-401a-91aa-fe7b599d3db3/file1.txt
	// 	00000000-0000-0000-0000-000000000000/b65ad795-5c8b-401a-91aa-fe7b599d3db3/models/
	//	00000000-0000-0000-0000-000000000000/b65ad795-5c8b-401a-91aa-fe7b599d3db3/models/file2.txt
	//
	// Remove folders and prefix to return this when prefix=<userid>/<uploadid>
	//
	//	file1.txt
	//	models/file2.txt

	ret := make([]string, 0, len(response.Contents))
	for _, v := range response.Contents {
		str := *v.Key
		// Skip folders. S3 returns folders as objects
		if strings.HasSuffix(str, "/") {
			continue
		}

		// Strip the root folder
		str = strings.TrimPrefix(str, folder)
		ret = append(ret, str)
	}
	return ret, nil
}
