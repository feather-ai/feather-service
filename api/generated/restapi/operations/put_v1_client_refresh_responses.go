// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"feather-ai/service-core/api/generated/models"
)

// PutV1ClientRefreshOKCode is the HTTP code returned for type PutV1ClientRefreshOK
const PutV1ClientRefreshOKCode int = 200

/*PutV1ClientRefreshOK OK - includes new token

swagger:response putV1ClientRefreshOK
*/
type PutV1ClientRefreshOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewPutV1ClientRefreshOK creates PutV1ClientRefreshOK with default headers values
func NewPutV1ClientRefreshOK() *PutV1ClientRefreshOK {

	return &PutV1ClientRefreshOK{}
}

// WithPayload adds the payload to the put v1 client refresh o k response
func (o *PutV1ClientRefreshOK) WithPayload(payload *models.LoginResponse) *PutV1ClientRefreshOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put v1 client refresh o k response
func (o *PutV1ClientRefreshOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutV1ClientRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutV1ClientRefreshForbiddenCode is the HTTP code returned for type PutV1ClientRefreshForbidden
const PutV1ClientRefreshForbiddenCode int = 403

/*PutV1ClientRefreshForbidden Unauthorized

swagger:response putV1ClientRefreshForbidden
*/
type PutV1ClientRefreshForbidden struct {
}

// NewPutV1ClientRefreshForbidden creates PutV1ClientRefreshForbidden with default headers values
func NewPutV1ClientRefreshForbidden() *PutV1ClientRefreshForbidden {

	return &PutV1ClientRefreshForbidden{}
}

// WriteResponse to the client
func (o *PutV1ClientRefreshForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
