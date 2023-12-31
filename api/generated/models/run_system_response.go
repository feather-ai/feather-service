// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunSystemResponse run system response
//
// swagger:model runSystemResponse
type RunSystemResponse struct {

	// output location
	OutputLocation string `json:"outputLocation,omitempty"`

	// outputs
	Outputs []interface{} `json:"outputs"`

	// tty
	Tty string `json:"tty,omitempty"`
}

// Validate validates this run system response
func (m *RunSystemResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this run system response based on context it is used
func (m *RunSystemResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunSystemResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunSystemResponse) UnmarshalBinary(b []byte) error {
	var res RunSystemResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
