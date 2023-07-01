package api

import (
	"context"
	"encoding/json"
	"feather-ai/service-core/api/generated/models"
	"feather-ai/service-core/api/generated/restapi/operations"
	"feather-ai/service-core/lib/aisystems"
	"feather-ai/service-core/lib/aisystems/definition"
	"feather-ai/service-core/lib/gatekeeper"
	"io/ioutil"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

/*
* Install all the handlers needed to manage uploads of models
 */
func InstallAiSystemHandlers(api *operations.CoreAPI) {
	api.GetV1PublicSystemHandler = operations.GetV1PublicSystemHandlerFunc(LookupSystemDetails)
	api.GetV1PublicSystemsHandler = operations.GetV1PublicSystemsHandlerFunc(ListSystems)
	api.GetV1PublicSystemSystemIDHandler = operations.GetV1PublicSystemSystemIDHandlerFunc(GetSystemDetails)

	api.PutV1PublicSystemSystemIDStepStepIndexHandler = operations.PutV1PublicSystemSystemIDStepStepIndexHandlerFunc(RunSystemStep)
	api.GetV1DebugExecuteRequestSchemaHandler = operations.GetV1DebugExecuteRequestSchemaHandlerFunc(DebugGetExecuteRequestSchema)
}

/*
* Get a list of all available systems
 */
func ListSystems(params operations.GetV1PublicSystemsParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()

	var err error
	var systems []aisystems.SystemDTO
	if params.Username != nil {
		username := *params.Username
		userInfo, err := gUserManager.GetUserByName(ctx, username)
		if err != nil {
			return operations.NewGetV1PublicSystemSystemIDBadRequest().WithPayload(models.GenericError(err.Error()))
		}
		if userInfo == nil {
			return operations.NewGetV1PublicSystemsNotFound()
		}

		systems, err = gSystemsManager.GetSystemsByUser(ctx, userInfo.UserID)
	} else {
		systems, err = gSystemsManager.GetSystems(ctx)
	}
	if err != nil {
		return operations.NewGetV1PublicSystemsNotFound()
	}

	// Lookup User Info. This returns a map of userId -> userDto
	userIds := []uuid.UUID{}
	for _, system := range systems {
		userIds = append(userIds, system.UserID)
	}
	userMap, err := gUserManager.LookUpUsersByIDs(ctx, userIds)
	if err != nil {
		return operations.NewGetV1PublicSystemSystemIDBadRequest().WithPayload(models.GenericError(err.Error()))
	}

	// Build the response
	response := []*models.SystemInfo{}
	for _, system := range systems {
		created := strfmt.DateTime(system.Created)
		id := system.SystemID.String()
		ownerId := system.UserID.String()
		ownerDto := userMap[ownerId]

		response = append(response, &models.SystemInfo{
			Created:     created,
			Description: string(system.Description),
			ID:          id,
			Name:        system.Name,
			Slug:        system.Slug,
			OwnerID:     ownerId,
			OwnerName:   ownerDto.Name,
		})
	}

	return operations.NewGetV1PublicSystemsOK().WithPayload(response)
}

func getSystemDetailsByID(ctx context.Context, systemId uuid.UUID) middleware.Responder {
	system, version, files, err := gSystemsManager.GetSystemDetails(ctx, systemId)
	if err != nil {
		return operations.NewGetV1PublicSystemSystemIDBadRequest().WithPayload(models.GenericError(err.Error()))
	}

	response := models.SystemDetails{}

	response.Name = system.Name
	response.Description = string(system.Description)
	response.LastUpdated = strfmt.DateTime(version.Created)
	response.Created = strfmt.DateTime(system.Created)
	response.OwnerID = system.UserID.String()
	response.SystemID = systemId.String()
	response.ID = systemId.String()

	err = json.Unmarshal([]byte(version.Schema), &response.Schema)
	if err != nil {
		logrus.Warnf("GetSystemDetails: Could not unmarshal json: %v", err)
		return operations.NewGetV1PublicSystemSystemIDBadRequest().WithPayload(models.GenericError(err.Error()))
	}

	rawDef, err := definition.Parse(version.Schema)
	if err != nil {
		logrus.Warnf("GetSystemDetails: Could not unmarshal json: %v", err)
		return operations.NewGetV1PublicSystemSystemIDBadRequest().WithPayload(models.GenericError(err.Error()))
	}

	response.NumSteps = int64(rawDef.NumSteps())

	if len(files) > 0 {
		response.Files = make([]*models.SystemDetailsFilesItems0, len(files))
		for k, f := range files {
			finfo := models.SystemDetailsFilesItems0{
				Created: strfmt.DateTime(f.Created),
				Name:    f.FileName,
				Type:    f.FileType,
			}
			response.Files[k] = &finfo
		}
	}
	return operations.NewGetV1PublicSystemSystemIDOK().WithPayload(&response)
}

/*
* Lookup system details and return full information on a system
 */
func LookupSystemDetails(params operations.GetV1PublicSystemParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if params.Username == nil || params.Systemname == nil {
		return operations.NewGetV1PublicSystemBadRequest()
	}

	username := *params.Username
	systemname := *params.Systemname

	userInfo, err := gUserManager.GetUserByName(ctx, username)
	if err != nil {
		return operations.NewGetV1PublicSystemBadRequest().WithPayload(models.GenericError(err.Error()))
	}

	if userInfo == nil {
		return operations.NewGetV1PublicSystemNotFound()
	}

	systemInfo := gSystemsManager.GetSystemByUserAndSlug(ctx, userInfo.UserID, systemname)
	if systemInfo == nil {
		return operations.NewGetV1PublicSystemNotFound()
	}

	return getSystemDetailsByID(ctx, systemInfo.SystemID)
}

/*
* Retrieve full information on a system
 */
func GetSystemDetails(params operations.GetV1PublicSystemSystemIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	systemId := uuid.FromStringOrNil(params.SystemID.String())
	return getSystemDetailsByID(ctx, systemId)
}

func DebugGetExecuteRequestSchema(params operations.GetV1DebugExecuteRequestSchemaParams, auth interface{}) middleware.Responder {

	systemId := uuid.FromStringOrNil(params.SystemID)
	ctx := params.HTTPRequest.Context()
	result, err := gSystemsManager.DebugGetRunSystemSchema(ctx, systemId, 0)

	if err != nil {
		return operations.NewGetV1DebugExecuteRequestSchemaBadRequest()
	}

	return operations.NewGetV1DebugExecuteRequestSchemaOK().WithPayload(result)
}

/*
* Called to run the step of a system
 */
func RunSystemStep(params operations.PutV1PublicSystemSystemIDStepStepIndexParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()

	err := gatekeeper.CanPerhormHTTPRequest(ctx, params.HTTPRequest, "RunSystem", gAnonymousGateKeeper)
	if err != nil {
		return operations.NewGetV1PublicSystemsTooManyRequests().WithPayload(models.GenericError(err.Error()))
	}

	inputBytes, err := ioutil.ReadAll(params.InputData)
	if err != nil {
		return operations.NewPutV1PublicSystemSystemIDStepStepIndexBadRequest().WithPayload(&models.RunSystemError{
			Error: err.Error(),
			Tty:   "",
		})
	}

	if len(inputBytes) > 1024*1024*10 {
		return operations.NewPutV1PublicSystemSystemIDStepStepIndexRequestEntityTooLarge()
	}

	systemId := uuid.FromStringOrNil(params.SystemID.String())
	stepIndex := params.StepIndex

	outputs, output_location, tty, err := gSystemsManager.RunSystem(ctx, systemId, int(stepIndex), string(inputBytes))
	if err != nil {
		return operations.NewPutV1PublicSystemSystemIDStepStepIndexBadRequest().WithPayload(&models.RunSystemError{
			Error: err.Error(),
			Tty:   tty,
		})
	}

	// Convert the outputs
	ret := models.RunSystemResponse{
		Outputs:        make([]interface{}, len(outputs)),
		OutputLocation: output_location,
		Tty:            tty,
	}
	for k, v := range outputs {
		ret.Outputs[k] = v
	}

	gatekeeper.RecordHTTPRequest(ctx, params.HTTPRequest, "RunSystem", gAnonymousGateKeeper)
	return operations.NewPutV1PublicSystemSystemIDStepStepIndexOK().WithPayload(&ret)
}
