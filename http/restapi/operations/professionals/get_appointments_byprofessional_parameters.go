// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppointmentsByprofessionalParams creates a new GetAppointmentsByprofessionalParams object
// with the default values initialized.
func NewGetAppointmentsByprofessionalParams() GetAppointmentsByprofessionalParams {

	var (
		// initialize parameters with default values

		finishDateDefault = strfmt.DateTime{}

		idspecialtyDefault = int64(0)
		startDateDefault   = strfmt.DateTime{}
		statusDefault      = []string{"confirmed", "pending"}
	)

	finishDateDefault.UnmarshalText([]byte("2040-01-01T00:00:00Z"))

	startDateDefault.UnmarshalText([]byte("2000-01-01T00:00:00Z"))

	return GetAppointmentsByprofessionalParams{
		FinishDate: &finishDateDefault,

		Idspecialty: &idspecialtyDefault,

		StartDate: &startDateDefault,

		Status: statusDefault,
	}
}

// GetAppointmentsByprofessionalParams contains all the bound params for the get appointments byprofessional operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAppointmentsByprofessional
type GetAppointmentsByprofessionalParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*finish date for appointment
	  In: query
	  Default: "2040-01-01T00:00:00Z"
	*/
	FinishDate *strfmt.DateTime
	/*id professional
	  Required: true
	  In: path
	*/
	ID int64
	/*id for specialty
	  In: query
	  Default: 0
	*/
	Idspecialty *int64
	/*start date for appointment
	  In: query
	  Default: "2000-01-01T00:00:00Z"
	*/
	StartDate *strfmt.DateTime
	/*appointment status
	  In: query
	  Default: []interface {}{"confirmed", "pending"}
	*/
	Status []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAppointmentsByprofessionalParams() beforehand.
func (o *GetAppointmentsByprofessionalParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFinishDate, qhkFinishDate, _ := qs.GetOK("finishDate")
	if err := o.bindFinishDate(qFinishDate, qhkFinishDate, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qIdspecialty, qhkIdspecialty, _ := qs.GetOK("idspecialty")
	if err := o.bindIdspecialty(qIdspecialty, qhkIdspecialty, route.Formats); err != nil {
		res = append(res, err)
	}

	qStartDate, qhkStartDate, _ := qs.GetOK("startDate")
	if err := o.bindStartDate(qStartDate, qhkStartDate, route.Formats); err != nil {
		res = append(res, err)
	}

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFinishDate binds and validates parameter FinishDate from query.
func (o *GetAppointmentsByprofessionalParams) bindFinishDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAppointmentsByprofessionalParams()
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("finishDate", "query", "strfmt.DateTime", raw)
	}
	o.FinishDate = (value.(*strfmt.DateTime))

	if err := o.validateFinishDate(formats); err != nil {
		return err
	}

	return nil
}

// validateFinishDate carries on validations for parameter FinishDate
func (o *GetAppointmentsByprofessionalParams) validateFinishDate(formats strfmt.Registry) error {

	if err := validate.FormatOf("finishDate", "query", "date-time", o.FinishDate.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetAppointmentsByprofessionalParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindIdspecialty binds and validates parameter Idspecialty from query.
func (o *GetAppointmentsByprofessionalParams) bindIdspecialty(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAppointmentsByprofessionalParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("idspecialty", "query", "int64", raw)
	}
	o.Idspecialty = &value

	return nil
}

// bindStartDate binds and validates parameter StartDate from query.
func (o *GetAppointmentsByprofessionalParams) bindStartDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAppointmentsByprofessionalParams()
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("startDate", "query", "strfmt.DateTime", raw)
	}
	o.StartDate = (value.(*strfmt.DateTime))

	if err := o.validateStartDate(formats); err != nil {
		return err
	}

	return nil
}

// validateStartDate carries on validations for parameter StartDate
func (o *GetAppointmentsByprofessionalParams) validateStartDate(formats strfmt.Registry) error {

	if err := validate.FormatOf("startDate", "query", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindStatus binds and validates array parameter Status from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetAppointmentsByprofessionalParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvStatus string
	if len(rawData) > 0 {
		qvStatus = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	statusIC := swag.SplitByFormat(qvStatus, "")
	if len(statusIC) == 0 {
		// Default values have been previously initialized by NewGetAppointmentsByprofessionalParams()
		return nil
	}

	var statusIR []string
	for i, statusIV := range statusIC {
		statusI := statusIV

		if err := validate.Enum(fmt.Sprintf("%s.%v", "status", i), "query", statusI, []interface{}{"confirmed", "cancelled", "pending", "avaiable"}); err != nil {
			return err
		}

		statusIR = append(statusIR, statusI)
	}

	o.Status = statusIR

	return nil
}
