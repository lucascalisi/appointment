// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCancelAppoinmentForPatientParams creates a new CancelAppoinmentForPatientParams object
// no default values defined in spec.
func NewCancelAppoinmentForPatientParams() CancelAppoinmentForPatientParams {

	return CancelAppoinmentForPatientParams{}
}

// CancelAppoinmentForPatientParams contains all the bound params for the cancel appoinment for patient operation
// typically these are obtained from a http.Request
//
// swagger:parameters cancelAppoinmentForPatient
type CancelAppoinmentForPatientParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id patient
	  Required: true
	  In: path
	*/
	ID int64
	/*confirm appoinment
	  Required: true
	  In: path
	*/
	IDAppointment int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCancelAppoinmentForPatientParams() beforehand.
func (o *CancelAppoinmentForPatientParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rIDAppointment, rhkIDAppointment, _ := route.Params.GetOK("idAppointment")
	if err := o.bindIDAppointment(rIDAppointment, rhkIDAppointment, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *CancelAppoinmentForPatientParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindIDAppointment binds and validates parameter IDAppointment from path.
func (o *CancelAppoinmentForPatientParams) bindIDAppointment(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("idAppointment", "path", "int64", raw)
	}
	o.IDAppointment = value

	return nil
}
