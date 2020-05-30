// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// RequestAppointmentForPatientURL generates an URL for the request appointment for patient operation
type RequestAppointmentForPatientURL struct {
	ID            int64
	IDAppointment int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RequestAppointmentForPatientURL) WithBasePath(bp string) *RequestAppointmentForPatientURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RequestAppointmentForPatientURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RequestAppointmentForPatientURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/patients/{id}/appointments/{idAppointment}/request"

	id := swag.FormatInt64(o.ID)
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on RequestAppointmentForPatientURL")
	}

	iDAppointment := swag.FormatInt64(o.IDAppointment)
	if iDAppointment != "" {
		_path = strings.Replace(_path, "{idAppointment}", iDAppointment, -1)
	} else {
		return nil, errors.New("idAppointment is required on RequestAppointmentForPatientURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RequestAppointmentForPatientURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RequestAppointmentForPatientURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RequestAppointmentForPatientURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RequestAppointmentForPatientURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RequestAppointmentForPatientURL")
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
func (o *RequestAppointmentForPatientURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
