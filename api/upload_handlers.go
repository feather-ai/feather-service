package api

import (
	"errors"
	"feather-ai/service-core/api/generated/models"
	"feather-ai/service-core/api/generated/restapi/operations"
	"fmt"
	"strings"

	"feather-ai/service-core/lib/aisystems/definition"
	"feather-ai/service-core/lib/config"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

/*
* Install all the handlers needed to manage uploads of models
 */
func InstallUploadHandlers(api *operations.CoreAPI) {
	api.PutV1APISystemPreparePublishHandler = operations.PutV1APISystemPreparePublishHandlerFunc(PreparePublish)
	api.PutV1APISystemCompletePublishHandler = operations.PutV1APISystemCompletePublishHandlerFunc(CompletePublish)
}

func ValidateFilename(filename string) (string, error) {
	//  :TODO:
	return filename, nil
}

// HELPER - Parse the filenames from the payload
func loadFiles(params operations.PutV1APISystemPreparePublishParams) ([]string, []string, error) {
	codeFiles := make([]string, 0, 8)
	modelFiles := make([]string, 0, 8)
	for _, file := range params.Definition.Files {
		filename, err := ValidateFilename(file.Filename)
		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("Invalid filename %s", file.Filename))
		}

		if file.Filetype == "python" {
			codeFiles = append(codeFiles, filename)
		} else if file.Filetype == "model" {
			modelFiles = append(modelFiles, filename)
		} else {
			return nil, nil, errors.New(fmt.Sprintf("Unknown filetype %s", file.Filetype))
		}
	}
	return codeFiles, modelFiles, nil
}

/*
* Prepare to Publish a system to feather
 */
func PreparePublish(params operations.PutV1APISystemPreparePublishParams, authenticatedUserId interface{}) middleware.Responder {

	defSpec := gSpec.Definitions["PreparePublishRequest"]

	if err := validate.AgainstSchema(&defSpec, params.Definition, strfmt.Default); err != nil {
		return operations.NewPutV1APISystemPreparePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	ctx := params.HTTPRequest.Context()
	reqCtx := config.GetRequestContext(ctx)
	reqCtx.UserID = authenticatedUserId.(uuid.UUID)

	log := logrus.WithContext(ctx)

	// Extract  and validate the params

	log.Infof("Prepare upload: user=%v", reqCtx.UserID.String())

	name := *params.Definition.Name
	slug := *params.Definition.Slug
	desc := params.Definition.Description
	schema := *params.Definition.Schema

	// Validate the schema
	def, err := definition.Parse(schema)
	if err != nil || def.NumSteps() == 0 {
		return operations.NewPutV1APISystemPreparePublishBadRequest().
			WithPayload(models.GenericError("Invalid Schema"))
	}

	systemId, err := gSystemsManager.GetOrCreateNewSystem(ctx, reqCtx.UserID, name, slug, desc)
	if err != nil {
		return operations.NewPutV1APISystemPreparePublishInternalServerError().
			WithPayload(models.GenericError(err.Error()))
	}

	codeFiles, modelFiles, err := loadFiles(params)
	if err != nil {
		return operations.NewPutV1APISystemPreparePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	// Send the request
	req, err := gUploadManager.CreateUploadRequest(ctx, systemId, "0.0.0", codeFiles, modelFiles, schema)
	if err != nil {
		return operations.NewPutV1APISystemPreparePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	// Convert the response

	id := req.ID.String()
	expiry := strfmt.DateTime(req.ExpireTime)

	files := make([]*models.PreparePublishResponseFilesItems0, 0, 8)
	codeSignedUrls := strings.Split(req.CodeFilesSignedURL, "\n")
	codeFiles = strings.Split(req.CodeFiles, "\n")
	for idx, signedFile := range codeSignedUrls {
		if len(signedFile) > 0 {
			files = append(files, &models.PreparePublishResponseFilesItems0{
				Filename:  codeFiles[idx],
				UploadURL: signedFile,
			})
		}
	}
	modelSignedUrls := strings.Split(req.ModelFilesSignedURL, "\n")
	modelFiles = strings.Split(req.ModelFiles, "\n")
	for idx, signedFile := range modelSignedUrls {
		if len(signedFile) > 0 {
			files = append(files, &models.PreparePublishResponseFilesItems0{
				Filename:  modelFiles[idx],
				UploadURL: signedFile,
			})
		}
	}

	log.Infof("Successfull prepareUpload. systemId=%s", systemId.String())
	return operations.NewPutV1APISystemPreparePublishOK().WithPayload(&models.PreparePublishResponse{
		Files:      files,
		ID:         id,
		ExpiryTime: expiry,
	})
}

/*
* Finish the publishing of a system.
 */
func CompletePublish(params operations.PutV1APISystemCompletePublishParams, authenticatedUserId interface{}) middleware.Responder {
	defSpec := gSpec.Definitions["CompletePublishRequest"]
	if err := validate.AgainstSchema(&defSpec, params.Definition, strfmt.Default); err != nil {
		return operations.NewPutV1APISystemCompletePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	ctx := params.HTTPRequest.Context()
	reqCtx := config.GetRequestContext(ctx)
	reqCtx.UserID = authenticatedUserId.(uuid.UUID)

	id := uuid.FromStringOrNil(*params.Definition.ID)

	req, err := gUploadManager.TryCompleteUpload(ctx, id)
	if err != nil {
		return operations.NewPutV1APISystemCompletePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	systemId := req.SystemID
	userId := req.UserID

	systemInfo, err := gSystemsManager.GetSystemByID(ctx, uuid.FromStringOrNil(systemId))
	if err != nil {
		return operations.NewPutV1APISystemCompletePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	userInfo, err := gUserManager.GetUserByID(ctx, uuid.FromStringOrNil(userId))
	if err != nil {
		return operations.NewPutV1APISystemCompletePublishBadRequest().
			WithPayload(models.GenericError(err.Error()))
	}

	return operations.NewPutV1APISystemCompletePublishOK().WithPayload(&models.CompletePublishResponse{
		User:   userInfo.Name,
		System: systemInfo.Slug,
	})
}

/*
* Get the list of upload requests for the logged in user
 */
/*func GetUserUploadRequests(params operations.GetV1ModelUploadsParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()
	userId := config.GetRequestContext(ctx).UserID

	result := gUploadManager.GetUploadRequestsForUser(ctx, userId)
	response := []*models.UploadRequest{}

	if result != nil {
		for _, v := range result {
			id := v.ID.String()
			expiry := strfmt.DateTime(v.ExpireTime)
			r := &models.UploadRequest{
				ID:         &id,
				ExpiryTime: &expiry,
				UploadURL:  &v.URL,
			}

			response = append(response, r)
		}
	}

	return operations.NewGetV1ModelUploadsOK().WithPayload(response)
}*/
