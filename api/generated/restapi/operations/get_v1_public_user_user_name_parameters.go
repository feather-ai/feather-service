// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetV1PublicUserUserNameParams creates a new GetV1PublicUserUserNameParams object
//
// There are no default values defined in the spec.
func NewGetV1PublicUserUserNameParams() GetV1PublicUserUserNameParams {

	return GetV1PublicUserUserNameParams{}
}

// GetV1PublicUserUserNameParams contains all the bound params for the get v1 public user user name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetV1PublicUserUserName
type GetV1PublicUserUserNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The user name of the user to lookup
	  Required: true
	  In: path
	*/
	UserName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetV1PublicUserUserNameParams() beforehand.
func (o *GetV1PublicUserUserNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserName, rhkUserName, _ := route.Params.GetOK("userName")
	if err := o.bindUserName(rUserName, rhkUserName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserName binds and validates parameter UserName from path.
func (o *GetV1PublicUserUserNameParams) bindUserName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UserName = raw

	return nil
}
