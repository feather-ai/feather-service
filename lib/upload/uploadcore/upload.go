package uploadcore

import (
	"context"
	"errors"
	"feather-ai/service-core/lib/aisystems"
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/storage"
	"feather-ai/service-core/lib/upload"
	"fmt"
	"sort"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type uploadCoreManager struct {
	BlobStorage    storage.BlobStorage // Where we send uploads to
	DB             storage.DB
	SystemsManager aisystems.AISystemsManager
	options        config.Options
}

func New(manager storage.StorageManager,
	systemsManager aisystems.AISystemsManager,
	opts config.Options) upload.UploadManager {

	prepareDBQueries(manager.GetDB())

	return &uploadCoreManager{
		BlobStorage:    manager.GetBlobStorage(opts.AwsUploadBucketName),
		DB:             manager.GetDB(),
		SystemsManager: systemsManager,
	}
}

func getLoadingBayBasePath(userId string, systemId string, reqId string) string {
	return fmt.Sprintf("%s/%s/%s/", userId, systemId, reqId)
}

func (m *uploadCoreManager) CreateUploadRequest(ctx context.Context,
	systemId uuid.UUID, versionTag string,
	codeFiles []string, modelFiles []string, schema string) (*upload.UploadRequestDTO, error) {

	log := logrus.WithContext(ctx)
	req := config.GetRequestContext(ctx)

	dto := &upload.UploadRequestDTO{
		UserID:     req.UserID.String(),
		SystemID:   systemId.String(),
		VersionTag: versionTag,
		ID:         uuid.NewV4(),
		Schema:     schema,
	}

	// Create the signed URL and record

	duration := time.Hour * 1 // :TODO: Get this from config

	//if semver.IsValid(systemVersion) == false {
	//	return nil, errors.New("Version is not a semver")
	//}

	// :TODO: Check all files paths

	basePath := getLoadingBayBasePath(req.UserID.String(), systemId.String(), dto.ID.String())

	for _, file := range codeFiles {
		url, err := m.BlobStorage.GenerateUploadURL(
			ctx,
			basePath+file,
			duration,
		)
		if err != nil {
			log.Errorf("CreateUploadRequest: Generate Code URL(%s): %v", basePath+file, err)
			return nil, err
		}
		dto.CodeFiles += file + "\n"
		dto.CodeFilesSignedURL += url + "\n"
	}
	for _, file := range modelFiles {
		url, err := m.BlobStorage.GenerateUploadURL(
			ctx,
			basePath+file,
			duration,
		)
		if err != nil {
			log.Errorf("CreateUploadRequest: Generate Model URL(%s): %v", basePath+file, err)
			return nil, err
		}
		dto.ModelFiles += file + "\n"
		dto.ModelFilesSignedURL += url + "\n"
	}

	dto.CreateTime = time.Now().UTC()
	dto.ExpireTime = dto.CreateTime.Add(duration)

	// Save to DB

	err := gInsertUploadRequestStmt.Write(ctx, &dto)
	if err != nil {
		log.Errorf("CreateUploadRequest: %v", err)
		return nil, err
	}

	return dto, nil
}

/*
*  To complete an upload/publish, we first check that all the files that should be, are in fact in S3.
* Once confirmed, we then remove the upload request and create a new system_version and records in the file
* table.
 */
func (m *uploadCoreManager) TryCompleteUpload(ctx context.Context, reqId uuid.UUID) (*upload.UploadRequestDTO, error) {

	userId := config.GetRequestContext(ctx).UserID
	log := logrus.WithContext(ctx)

	result := []upload.UploadRequestDTO{}

	uploadRequestArgs := map[string]interface{}{"id": reqId}
	err := gQueryUploadRequestByIdStmt.QueryMany(ctx, uploadRequestArgs, &result)
	if err != nil {
		log.Errorf("CanCompleteUpload: Failed to query request(id=%s): %v\n", reqId.String(), err)
		return nil, errors.New("Cannot find upload request")
	}

	if len(result) == 0 {
		log.Errorf("CanCompleteUpload: Nothing to query request(id=%s): %v\n", reqId.String(), err)
		return nil, errors.New("Cannot find upload request")
	}

	if time.Now().UTC().After(result[0].ExpireTime) {
		return nil, errors.New("Request expired")
	}

	// Check all the files exist on S3

	basePath := getLoadingBayBasePath(userId.String(), result[0].SystemID, reqId.String())

	blobs, err := m.BlobStorage.GetFilesInFolder(ctx, basePath)
	if err != nil {
		return nil, err
	}

	codeFiles := ExtractPackedFilenames(result[0].CodeFiles)
	modelFiles := ExtractPackedFilenames(result[0].ModelFiles)

	files := make([]string, 0)
	files = append(files, codeFiles...)
	files = append(files, modelFiles...)

	if len(blobs) != len(files) {
		log.Infof("CanCompleteUpload: Files in blob storage (n=%v) don't match expected files (n=%v)", len(blobs), len(files))
		return nil, errors.New("Uploaded files don't match expected count")
	}

	sort.Sort(sort.StringSlice(files))
	sort.Sort(sort.StringSlice(blobs))

	for i := 0; i < len(files); i++ {
		if strings.ToLower(files[i]) != strings.ToLower(blobs[i]) {
			log.Infof("File mismatch. %s does not match expected: %s", blobs[i], files[i])
			return nil, errors.New("Uploaded files don't match expected!")
		}
	}

	// Create a new system version
	// :TODO: Put in a transaction...
	systemId := uuid.FromStringOrNil(result[0].SystemID)
	versionId, err := m.SystemsManager.CreateNewVersion(ctx, systemId, result[0].VersionTag, result[0].Schema)
	if err != nil {
		log.Errorf("Failed to create new version for system: %v", err)
		return nil, err
	}

	// All files are where they should be. Create the system version
	dtos := make([]aisystems.FileDTO, 0, len(files))
	timestamp := time.Now().UTC()

	for _, file := range codeFiles {
		dtos = append(dtos, aisystems.FileDTO{
			SystemVersionID: versionId,
			FileName:        file,
			FileType:        "python",
			URL:/*m.BlobStorage.GetRootPath() + */ basePath + file,
			Created: timestamp,
		})
	}
	for _, file := range modelFiles {
		dtos = append(dtos, aisystems.FileDTO{
			SystemVersionID: versionId,
			FileName:        file,
			FileType:        "model",
			URL:/*m.BlobStorage.GetRootPath() + */ basePath + file,
			Created: timestamp,
		})
	}

	log.Infof("Commiting new version. system=%v, version=%v", systemId.String(), versionId)
	tx, err := m.DB.NewTransaction(ctx)
	{
		if err != nil {
			log.Errorf("Failed to create DB Transaction: %v", err)
			return nil, err
		}
		defer tx.Commit(ctx)

		for _, dto := range dtos {
			tx.Execute(ctx, gInsertFileRequestStmt, &dto)
		}
		tx.Execute(ctx, gDeleteUploadRequestByIdStmt, uploadRequestArgs)
	}

	return &result[0], tx.Error()
}

func (m *uploadCoreManager) GetUploadRequestsForUser(ctx context.Context, userId uuid.UUID) []upload.UploadRequestDTO {
	result := []upload.UploadRequestDTO{}

	args := map[string]interface{}{"userId": userId}
	err := gQueryUserUploadRequestsStmt.QueryMany(ctx, args, &result)
	if err != nil {
		logrus.Errorf("GetUploadRequestsForUser: Failed to query requests: %v\n", err)
		return nil
	}
	return result
}
