// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"feather-ai/service-core/api/generated/models"
)

// NewPutV1APISystemCompletePublishParams creates a new PutV1APISystemCompletePublishParams object
//
// There are no default values defined in the spec.
func NewPutV1APISystemCompletePublishParams() PutV1APISystemCompletePublishParams {

	return PutV1APISystemCompletePublishParams{}
}

// PutV1APISystemCompletePublishParams contains all the bound params for the put v1 API system complete publish operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutV1APISystemCompletePublish
type PutV1APISystemCompletePublishParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Definition *models.CompletePublishRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutV1APISystemCompletePublishParams() beforehand.
func (o *PutV1APISystemCompletePublishParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CompletePublishRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("definition", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Definition = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
