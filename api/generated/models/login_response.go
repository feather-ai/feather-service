// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginResponse login response
//
// swagger:model loginResponse
type LoginResponse struct {

	// expire at
	// Format: date-time
	ExpireAt strfmt.DateTime `json:"expireAt,omitempty"`

	// feather token
	FeatherToken string `json:"featherToken,omitempty"`
}

// Validate validates this login response
func (m *LoginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpireAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginResponse) validateExpireAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expireAt", "body", "date-time", m.ExpireAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login response based on context it is used
func (m *LoginResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginResponse) UnmarshalBinary(b []byte) error {
	var res LoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
