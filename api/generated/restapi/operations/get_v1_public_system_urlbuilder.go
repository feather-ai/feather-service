// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetV1PublicSystemURL generates an URL for the get v1 public system operation
type GetV1PublicSystemURL struct {
	Systemname *string
	Username   *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1PublicSystemURL) WithBasePath(bp string) *GetV1PublicSystemURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1PublicSystemURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetV1PublicSystemURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/public/system"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var systemnameQ string
	if o.Systemname != nil {
		systemnameQ = *o.Systemname
	}
	if systemnameQ != "" {
		qs.Set("systemname", systemnameQ)
	}

	var usernameQ string
	if o.Username != nil {
		usernameQ = *o.Username
	}
	if usernameQ != "" {
		qs.Set("username", usernameQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetV1PublicSystemURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetV1PublicSystemURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetV1PublicSystemURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetV1PublicSystemURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetV1PublicSystemURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetV1PublicSystemURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
