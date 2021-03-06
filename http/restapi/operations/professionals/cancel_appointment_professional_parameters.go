// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCancelAppointmentProfessionalParams creates a new CancelAppointmentProfessionalParams object
// no default values defined in spec.
func NewCancelAppointmentProfessionalParams() CancelAppointmentProfessionalParams {

	return CancelAppointmentProfessionalParams{}
}

// CancelAppointmentProfessionalParams contains all the bound params for the cancel appointment professional operation
// typically these are obtained from a http.Request
//
// swagger:parameters cancelAppointmentProfessional
type CancelAppointmentProfessionalParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id professional
	  Required: true
	  In: path
	*/
	ID int64
	/*id appointment
	  Required: true
	  In: path
	*/
	Idappointment int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCancelAppointmentProfessionalParams() beforehand.
func (o *CancelAppointmentProfessionalParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rIdappointment, rhkIdappointment, _ := route.Params.GetOK("idappointment")
	if err := o.bindIdappointment(rIdappointment, rhkIdappointment, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *CancelAppointmentProfessionalParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}

// bindIdappointment binds and validates parameter Idappointment from path.
func (o *CancelAppointmentProfessionalParams) bindIdappointment(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("idappointment", "path", "int64", raw)
	}
	o.Idappointment = value

	return nil
}
