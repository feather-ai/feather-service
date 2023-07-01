// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"feather-ai/service-core/api/generated/models"
)

// GetV1PublicUserUserNameOKCode is the HTTP code returned for type GetV1PublicUserUserNameOK
const GetV1PublicUserUserNameOKCode int = 200

/*GetV1PublicUserUserNameOK User Information

swagger:response getV1PublicUserUserNameOK
*/
type GetV1PublicUserUserNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserInfo `json:"body,omitempty"`
}

// NewGetV1PublicUserUserNameOK creates GetV1PublicUserUserNameOK with default headers values
func NewGetV1PublicUserUserNameOK() *GetV1PublicUserUserNameOK {

	return &GetV1PublicUserUserNameOK{}
}

// WithPayload adds the payload to the get v1 public user user name o k response
func (o *GetV1PublicUserUserNameOK) WithPayload(payload *models.UserInfo) *GetV1PublicUserUserNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 public user user name o k response
func (o *GetV1PublicUserUserNameOK) SetPayload(payload *models.UserInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1PublicUserUserNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1PublicUserUserNameBadRequestCode is the HTTP code returned for type GetV1PublicUserUserNameBadRequest
const GetV1PublicUserUserNameBadRequestCode int = 400

/*GetV1PublicUserUserNameBadRequest Bad request

swagger:response getV1PublicUserUserNameBadRequest
*/
type GetV1PublicUserUserNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.GenericError `json:"body,omitempty"`
}

// NewGetV1PublicUserUserNameBadRequest creates GetV1PublicUserUserNameBadRequest with default headers values
func NewGetV1PublicUserUserNameBadRequest() *GetV1PublicUserUserNameBadRequest {

	return &GetV1PublicUserUserNameBadRequest{}
}

// WithPayload adds the payload to the get v1 public user user name bad request response
func (o *GetV1PublicUserUserNameBadRequest) WithPayload(payload models.GenericError) *GetV1PublicUserUserNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 public user user name bad request response
func (o *GetV1PublicUserUserNameBadRequest) SetPayload(payload models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1PublicUserUserNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
