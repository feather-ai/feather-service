package aisystemscore

import (
	"context"
	"encoding/json"
	"errors"
	"feather-ai/service-core/lib/aisystems"
	"feather-ai/service-core/lib/aisystems/definition"
	"feather-ai/service-core/lib/config"
	"feather-ai/service-core/lib/feather"
	"feather-ai/service-core/lib/storage"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type aiSystemsCoreManager struct {
	options config.Options

	// Aws Resources
	awsSession *session.Session
	awsLambda  *lambda.Lambda
}

func New(manager storage.StorageManager, opts config.Options) aisystems.AISystemsManager {
	prepareDBQueries(manager.GetDB())

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(opts.AwsRegion)},
	)
	if err != nil {
		log.Fatalf("Can't create AWS Session: %v", err)
	}

	awsLambda := lambda.New(awsSession, &aws.Config{Region: aws.String(opts.AwsRegion)})
	return &aiSystemsCoreManager{
		options:    opts,
		awsSession: awsSession,
		awsLambda:  awsLambda,
	}
}

/*
* Look for a system by name and if we don't find anything, create one. Either way, return the SystemID
 */
func (m *aiSystemsCoreManager) GetOrCreateNewSystem(ctx context.Context, userId uuid.UUID, name string,
	slug string, description string) (uuid.UUID, error) {
	log := logrus.WithContext(ctx)
	dto := m.GetSystemByUserAndSlug(ctx, userId, slug)
	if dto != nil {
		log.Infof("GetOrCreateNewSystem - Found existing system slug=%s id=%v", slug, dto.SystemID.String())
		return dto.SystemID, nil
	}
	// Not found, create it
	log.Infof("Creating new system for user %s, with name=%s slug=%s", userId.String(), name, slug)

	systemID := uuid.NewV4()
	systemDto := aisystems.SystemDTO{
		SystemID:    systemID,
		UserID:      userId,
		Name:        name,
		Slug:        slug,
		Description: feather.String(description),
		Keywords:    "",
		Created:     time.Now().UTC(),
	}

	err := gInsertSystemStmt.Write(ctx, &systemDto)
	if err != nil {
		log.Errorf("CreateSystem: %v", err)
		return uuid.Nil, err
	}

	return systemID, nil
}

/*
* DEBUG function to return all the systems, unsorted, with no pagination - likely to be slow
 */
func (m *aiSystemsCoreManager) GetSystems(ctx context.Context) ([]aisystems.SystemDTO, error) {
	result := []aisystems.SystemDTO{}

	args := map[string]interface{}{}
	err := gQueryAllSystemsStmt.QueryMany(ctx, args, &result)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetSystems: %v\n", err)
		return nil, err
	}

	return result, nil
}

/*
* Get all the systems for a user.
* NOTE: Need to add pagination
 */
func (m *aiSystemsCoreManager) GetSystemsByUser(ctx context.Context, userId uuid.UUID) ([]aisystems.SystemDTO, error) {
	result := []aisystems.SystemDTO{}

	args := map[string]interface{}{"user_id": userId.String()}
	err := gQuerySystemsByUserStmt.QueryMany(ctx, args, &result)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetSystemsByUser: %v\n", err)
		return nil, err
	}

	return result, nil
}

func (m *aiSystemsCoreManager) GetSystemByID(ctx context.Context, systemId uuid.UUID) (*aisystems.SystemDTO, error) {
	systemInfo := []aisystems.SystemDTO{}
	args := map[string]interface{}{
		"system_id": systemId.String(),
	}

	err := gQuerySystemsByIDStmt.QueryMany(ctx, args, &systemInfo)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetsystemByID: %v\n", err)
		return nil, err
	}

	if len(systemInfo) != 1 {
		return nil, errors.New("GetsystemByID: Not Found")
	}

	return &systemInfo[0], nil
}

func (m *aiSystemsCoreManager) GetSystemByUserAndSlug(ctx context.Context, userId uuid.UUID, slug string) *aisystems.SystemDTO {
	result := []aisystems.SystemDTO{}
	args := map[string]interface{}{
		"user_id": userId.String(),
		"slug":    slug,
	}
	err := gQuerySystemsByUserAndSlugStmt.QueryMany(ctx, args, &result)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetsystemByUserAndSlug: %v\n", err)
		return nil
	}

	if len(result) == 1 {
		return &result[0]
	}
	return nil
}

func (m *aiSystemsCoreManager) GetSystemDetails(ctx context.Context, systemId uuid.UUID) (*aisystems.SystemDTO, *aisystems.SystemVersionDTO, []aisystems.FileDTO, error) {
	systemInfo := []aisystems.SystemDTO{}
	args := map[string]interface{}{
		"system_id": systemId.String(),
	}

	err := gQuerySystemsByIDStmt.QueryMany(ctx, args, &systemInfo)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetSystemDetails: %v\n", err)
		return nil, nil, nil, err
	}

	if len(systemInfo) != 1 {
		return nil, nil, nil, errors.New("GetSystemDetails: Not Found")
	}

	version, files, err := m.loadLatestVersion(ctx, systemId, true)
	if err != nil {
		return nil, nil, nil, err
	}
	return &systemInfo[0], version, files, nil
}

func (m *aiSystemsCoreManager) CreateNewVersion(ctx context.Context, systemId uuid.UUID, tag string, schema string) (int64, error) {
	// Not found, create it
	logrus.WithContext(ctx).Infof("Creating new version for system=%v", systemId.String())

	dto := aisystems.SystemVersionDTO{
		SystemID: systemId,
		Tag:      feather.String(tag),
		Schema:   schema,
		Created:  time.Now().UTC(),
	}

	id, err := gInsertSystemVersionStmt.WriteReturnID(ctx, &dto)
	if err != nil {
		logrus.WithContext(ctx).Errorf("CreateSystemVersion error: %v", err)
		return -1, err
	}

	return id, nil
}

/*
* Load the latest version for a system and return the SystemVersionDTO and all the files for this version.
 */
func (m *aiSystemsCoreManager) loadLatestVersion(ctx context.Context, systemId uuid.UUID, fetchFileInfo bool) (*aisystems.SystemVersionDTO, []aisystems.FileDTO, error) {
	log := logrus.WithContext(ctx)

	// Get the latest system version
	versionInfo := []aisystems.SystemVersionDTO{}

	args := map[string]interface{}{
		"system_id": systemId.String(),
	}
	err := gQueryLatestSystemVersionBySystemIdStmt.QueryMany(ctx, args, &versionInfo)
	if err != nil {
		log.Warnf("RunSystem: Can't get latest version: %v\n", err)
		return nil, nil, err
	}

	if len(versionInfo) != 1 {
		log.Warnf("RunSystem: No versions found for systemId: %v", systemId.String())
		return nil, nil, errors.New("No versions found")
	}

	if fetchFileInfo == true {
		// Get the files for this version
		files := []aisystems.FileDTO{}
		args = map[string]interface{}{
			"version_id": versionInfo[0].VersionID,
		}
		err = gQuerySystemVersionFilesStmt.QueryMany(ctx, args, &files)
		if err != nil {
			log.Warnf("RunSystem: Can't get files for system: %v\n", err)
			return nil, nil, err
		}
		return &versionInfo[0], files, nil
	}
	return &versionInfo[0], nil, nil
}

type ExecutionFileInfo struct {
	Filename string `json:"filename"`
	S3url    string `json:"s3Url"`
}
type ExecutionFiles struct {
	CodeFiles  []ExecutionFileInfo `json:"code_files"`
	ModelFiles []ExecutionFileInfo `json:"model_files"`
}
type ExecuteSystemRequest struct {
	Id         string                   `json:"id"`
	StepToRun  string                   `json:"step_to_run"`
	Definition definition.RawDefinition `json:"definition"`
	Files      ExecutionFiles           `json:"files"`
	InputData  json.RawMessage          `json:"input_data"`
}

// Payload returned by the service-runner
type LambdaResponsePayload struct {
	Result         int               `json:"result"`
	Outputs        []json.RawMessage `json:"outputs"`
	OutputLocation string            `json:"output_location"`
	Error          string            `json:"error"`
}

func (m *aiSystemsCoreManager) lambdaDispatch(ctx context.Context, lambdaName string, request ExecuteSystemRequest) ([]json.RawMessage, string, string, error) {
	log := logrus.WithContext(ctx)

	log.Infof("RunSystem: Invoking Lambda: %s. Payload", lambdaName)
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, "", "", err
	}

	result, err := m.awsLambda.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(lambdaName),
		Payload:      payload})

	// Handle the myriad of errors - first check for AWS specified errors
	if err != nil {
		log.Errorf("RunSystem: Internal error during lambda execution: %v", err)
		return nil, "", "", err
	}

	tty := ""
	if result.LogResult != nil {
		tty = *result.LogResult
	}

	if *result.StatusCode != 200 {
		log.Errorf("RunSystem: Internal error during lambda execution: code=%v  log=%s", *result.StatusCode, *result.LogResult)
		return nil, "", tty, errors.New("Internal Error Running Lambda - check log")
	}

	cleanJson, err := strconv.Unquote(string(result.Payload))
	if err != nil {
		log.Errorf("RunSystem: Could not Clean JSON-response: %v", err)
		return nil, "", tty, err
	}

	var output LambdaResponsePayload
	err = json.Unmarshal([]byte(cleanJson), &output)
	if err != nil {
		log.Errorf("RunSystem: Could not JSON-deserialize payload: %v", err)
		return nil, "", tty, err
	}

	if output.Result != 200 {
		log.Warnf("RunSystem: Error reported from lambda: %v", output.Error)
		return nil, "", tty, errors.New(output.Error)
	}
	return output.Outputs, output.OutputLocation, tty, nil
}

/*
* Run a specific step of a system
 */
func (m *aiSystemsCoreManager) RunSystem(ctx context.Context, systemId uuid.UUID,
	stepIndex int, inputDataJson string) ([]json.RawMessage, string, string, error) {
	log := logrus.WithContext(ctx)

	// Load the configuration for this system, if any
	lambdaName := "generic_runner"
	useGenericRunner := true
	systemConfig := []aisystems.SystemConfigDTO{}
	args := map[string]interface{}{"system_id": systemId.String()}
	err := gQuerySystemConfigByIDStmt.QueryMany(ctx, args, &systemConfig)
	if err != nil {
		log.Warnf("RunSystem: Can't get latest version: %v\n", err)
		return nil, "", "", err
	}

	if len(systemConfig) == 1 {
		lambda_dispatch := systemConfig[0].LambdaDispatch.String()
		if lambda_dispatch != "" && lambda_dispatch != "generic_runner" {
			useGenericRunner = false
			log.Infof("Using customer runner <%s> for system %v", lambda_dispatch, systemId)
			lambdaName = lambda_dispatch
		}
	}

	// Get the latest system version

	versionInfo, files, err := m.loadLatestVersion(ctx, systemId, useGenericRunner)
	if err != nil {
		return nil, "", "", err
	}

	rawDefinition, err := definition.Parse(versionInfo.Schema)
	if err != nil {
		log.Warnf("RunSystem: Could not parse system schema. SystemID=%v Version=%v", systemId.String(), versionInfo.VersionID)
		return nil, "", "", err
	}

	stepName, err := definition.GetStepNameByIndex(rawDefinition, stepIndex)
	if err != nil {
		return nil, "", "", err
	}

	// Build the Execution Request context
	filesReq := ExecutionFiles{
		CodeFiles:  make([]ExecutionFileInfo, 0),
		ModelFiles: make([]ExecutionFileInfo, 0),
	}
	if useGenericRunner == true {
		// Only fetch file info if using the generic runner
		for _, f := range files {
			data := ExecutionFileInfo{
				Filename: f.FileName,
				S3url:    f.URL,
			}
			if f.FileType == "python" {
				filesReq.CodeFiles = append(filesReq.CodeFiles, data)
			} else {
				filesReq.ModelFiles = append(filesReq.ModelFiles, data)
			}
		}
	}

	request := ExecuteSystemRequest{
		Id:         uuid.NewV4().String(),
		StepToRun:  stepName,
		Definition: rawDefinition,
		Files:      filesReq,
		InputData:  json.RawMessage(inputDataJson),
	}

	// Invoke the Lambda

	return m.lambdaDispatch(ctx, lambdaName, request)
}

/*
* Debug function only, to return the schema to run the specified system
 */
func (m *aiSystemsCoreManager) DebugGetRunSystemSchema(ctx context.Context, systemId uuid.UUID, stepIndex int) (interface{}, error) {
	log := logrus.WithContext(ctx)

	versionInfo, files, err := m.loadLatestVersion(ctx, systemId, true)
	if err != nil {
		return "", err
	}

	rawDefinition, err := definition.Parse(versionInfo.Schema)
	if err != nil {
		log.Warnf("RunSystem: Could not parse system schema. SystemID=%v Version=%v", systemId.String(), versionInfo.VersionID)
		return "", err
	}

	stepName, err := definition.GetStepNameByIndex(rawDefinition, stepIndex)
	if err != nil {
		return "", err
	}

	// Build the Execution Request context

	filesReq := ExecutionFiles{
		CodeFiles:  make([]ExecutionFileInfo, 0),
		ModelFiles: make([]ExecutionFileInfo, 0),
	}
	for _, f := range files {
		data := ExecutionFileInfo{
			Filename: f.FileName,
			S3url:    f.URL,
		}
		if f.FileType == "python" {
			filesReq.CodeFiles = append(filesReq.CodeFiles, data)
		} else {
			filesReq.ModelFiles = append(filesReq.ModelFiles, data)
		}
	}

	request := ExecuteSystemRequest{
		Id:         uuid.NewV4().String(),
		StepToRun:  stepName,
		Definition: rawDefinition,
		Files:      filesReq,
		InputData:  json.RawMessage("{}"),
	}

	return request, nil
}
