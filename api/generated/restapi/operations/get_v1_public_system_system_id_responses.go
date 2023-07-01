// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"feather-ai/service-core/api/generated/models"
)

// GetV1PublicSystemSystemIDOKCode is the HTTP code returned for type GetV1PublicSystemSystemIDOK
const GetV1PublicSystemSystemIDOKCode int = 200

/*GetV1PublicSystemSystemIDOK Detailed information about a system

swagger:response getV1PublicSystemSystemIdOK
*/
type GetV1PublicSystemSystemIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.SystemDetails `json:"body,omitempty"`
}

// NewGetV1PublicSystemSystemIDOK creates GetV1PublicSystemSystemIDOK with default headers values
func NewGetV1PublicSystemSystemIDOK() *GetV1PublicSystemSystemIDOK {

	return &GetV1PublicSystemSystemIDOK{}
}

// WithPayload adds the payload to the get v1 public system system Id o k response
func (o *GetV1PublicSystemSystemIDOK) WithPayload(payload *models.SystemDetails) *GetV1PublicSystemSystemIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 public system system Id o k response
func (o *GetV1PublicSystemSystemIDOK) SetPayload(payload *models.SystemDetails) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1PublicSystemSystemIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1PublicSystemSystemIDBadRequestCode is the HTTP code returned for type GetV1PublicSystemSystemIDBadRequest
const GetV1PublicSystemSystemIDBadRequestCode int = 400

/*GetV1PublicSystemSystemIDBadRequest Bad request

swagger:response getV1PublicSystemSystemIdBadRequest
*/
type GetV1PublicSystemSystemIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.GenericError `json:"body,omitempty"`
}

// NewGetV1PublicSystemSystemIDBadRequest creates GetV1PublicSystemSystemIDBadRequest with default headers values
func NewGetV1PublicSystemSystemIDBadRequest() *GetV1PublicSystemSystemIDBadRequest {

	return &GetV1PublicSystemSystemIDBadRequest{}
}

// WithPayload adds the payload to the get v1 public system system Id bad request response
func (o *GetV1PublicSystemSystemIDBadRequest) WithPayload(payload models.GenericError) *GetV1PublicSystemSystemIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 public system system Id bad request response
func (o *GetV1PublicSystemSystemIDBadRequest) SetPayload(payload models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1PublicSystemSystemIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
