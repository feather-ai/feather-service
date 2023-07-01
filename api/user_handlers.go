package api

import (
	"feather-ai/service-core/api/generated/models"
	"feather-ai/service-core/api/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

/*
* Install all the handlers needed to manage uploads of models
 */
func InstallUserHandlers(api *operations.CoreAPI) {
	api.GetV1PublicUserUserNameHandler = operations.GetV1PublicUserUserNameHandlerFunc(GetUserInfo)
}

/*
* Get information about a user
 */
func GetUserInfo(params operations.GetV1PublicUserUserNameParams) middleware.Responder {

	ctx := params.HTTPRequest.Context()
	userName := *&params.UserName

	dto, err := gUserManager.GetUserByName(ctx, userName)
	if err != nil {
		return operations.NewGetV1PublicUserUserNameBadRequest().WithPayload(models.GenericError(err.Error()))
	}

	response := models.UserInfo{
		UserID:   dto.UserID.String(),
		UserName: dto.Name,
	}
	return operations.NewGetV1PublicUserUserNameOK().WithPayload(&response)
}
