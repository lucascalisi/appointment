// Code generated by go-swagger; DO NOT EDIT.

package appointments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchAppointmentParams creates a new SearchAppointmentParams object
// with the default values initialized.
func NewSearchAppointmentParams() *SearchAppointmentParams {
	var (
		finishDateDefault     = strfmt.DateTime("2040-01-01T00:00:00Z")
		iDPatientDefault      = int64(0)
		iDProfessionalDefault = int64(0)
		idspecialtyDefault    = int64(0)
		startDateDefault      = strfmt.DateTime("2000-01-01T00:00:00Z")
		statusDefault         = []interface{}{"avaiable"}
	)
	return &SearchAppointmentParams{
		FinishDate:     &finishDateDefault,
		IDPatient:      &iDPatientDefault,
		IDProfessional: &iDProfessionalDefault,
		Idspecialty:    &idspecialtyDefault,
		StartDate:      &startDateDefault,
		Status:         statusDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAppointmentParamsWithTimeout creates a new SearchAppointmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchAppointmentParamsWithTimeout(timeout time.Duration) *SearchAppointmentParams {
	var (
		finishDateDefault     = strfmt.DateTime("2040-01-01T00:00:00Z")
		iDPatientDefault      = int64(0)
		iDProfessionalDefault = int64(0)
		idspecialtyDefault    = int64(0)
		startDateDefault      = strfmt.DateTime("2000-01-01T00:00:00Z")
		statusDefault         = []interface{}{"avaiable"}
	)
	return &SearchAppointmentParams{
		FinishDate:     &finishDateDefault,
		IDPatient:      &iDPatientDefault,
		IDProfessional: &iDProfessionalDefault,
		Idspecialty:    &idspecialtyDefault,
		StartDate:      &startDateDefault,
		Status:         statusDefault,

		timeout: timeout,
	}
}

// NewSearchAppointmentParamsWithContext creates a new SearchAppointmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchAppointmentParamsWithContext(ctx context.Context) *SearchAppointmentParams {
	var (
		finishDateDefault     = strfmt.DateTime("2040-01-01T00:00:00Z")
		idPatientDefault      = int64(0)
		idProfessionalDefault = int64(0)
		idspecialtyDefault    = int64(0)
		startDateDefault      = strfmt.DateTime("2000-01-01T00:00:00Z")
		statusDefault         = []interface{}{"avaiable"}
	)
	return &SearchAppointmentParams{
		FinishDate:     &finishDateDefault,
		IDPatient:      &idPatientDefault,
		IDProfessional: &idProfessionalDefault,
		Idspecialty:    &idspecialtyDefault,
		StartDate:      &startDateDefault,
		Status:         statusDefault,

		Context: ctx,
	}
}

// NewSearchAppointmentParamsWithHTTPClient creates a new SearchAppointmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchAppointmentParamsWithHTTPClient(client *http.Client) *SearchAppointmentParams {
	var (
		finishDateDefault     = strfmt.DateTime("2040-01-01T00:00:00Z")
		idPatientDefault      = int64(0)
		idProfessionalDefault = int64(0)
		idspecialtyDefault    = int64(0)
		startDateDefault      = strfmt.DateTime("2000-01-01T00:00:00Z")
		statusDefault         = []interface{}{"avaiable"}
	)
	return &SearchAppointmentParams{
		FinishDate:     &finishDateDefault,
		IDPatient:      &idPatientDefault,
		IDProfessional: &idProfessionalDefault,
		Idspecialty:    &idspecialtyDefault,
		StartDate:      &startDateDefault,
		Status:         statusDefault,
		HTTPClient:     client,
	}
}

/*SearchAppointmentParams contains all the parameters to send to the API endpoint
for the search appointment operation typically these are written to a http.Request
*/
type SearchAppointmentParams struct {

	/*FinishDate
	  finish date for appointment

	*/
	FinishDate *strfmt.DateTime
	/*IDPatient
	  id of the patient

	*/
	IDPatient *int64
	/*IDProfessional
	  id of the professional

	*/
	IDProfessional *int64
	/*Idspecialty
	  id for specialty

	*/
	Idspecialty *int64
	/*StartDate
	  start date for appointment

	*/
	StartDate *strfmt.DateTime
	/*Status
	  appointment status

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search appointment params
func (o *SearchAppointmentParams) WithTimeout(timeout time.Duration) *SearchAppointmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search appointment params
func (o *SearchAppointmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search appointment params
func (o *SearchAppointmentParams) WithContext(ctx context.Context) *SearchAppointmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search appointment params
func (o *SearchAppointmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search appointment params
func (o *SearchAppointmentParams) WithHTTPClient(client *http.Client) *SearchAppointmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search appointment params
func (o *SearchAppointmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFinishDate adds the finishDate to the search appointment params
func (o *SearchAppointmentParams) WithFinishDate(finishDate *strfmt.DateTime) *SearchAppointmentParams {
	o.SetFinishDate(finishDate)
	return o
}

// SetFinishDate adds the finishDate to the search appointment params
func (o *SearchAppointmentParams) SetFinishDate(finishDate *strfmt.DateTime) {
	o.FinishDate = finishDate
}

// WithIDPatient adds the iDPatient to the search appointment params
func (o *SearchAppointmentParams) WithIDPatient(iDPatient *int64) *SearchAppointmentParams {
	o.SetIDPatient(iDPatient)
	return o
}

// SetIDPatient adds the idPatient to the search appointment params
func (o *SearchAppointmentParams) SetIDPatient(iDPatient *int64) {
	o.IDPatient = iDPatient
}

// WithIDProfessional adds the iDProfessional to the search appointment params
func (o *SearchAppointmentParams) WithIDProfessional(iDProfessional *int64) *SearchAppointmentParams {
	o.SetIDProfessional(iDProfessional)
	return o
}

// SetIDProfessional adds the idProfessional to the search appointment params
func (o *SearchAppointmentParams) SetIDProfessional(iDProfessional *int64) {
	o.IDProfessional = iDProfessional
}

// WithIdspecialty adds the idspecialty to the search appointment params
func (o *SearchAppointmentParams) WithIdspecialty(idspecialty *int64) *SearchAppointmentParams {
	o.SetIdspecialty(idspecialty)
	return o
}

// SetIdspecialty adds the idspecialty to the search appointment params
func (o *SearchAppointmentParams) SetIdspecialty(idspecialty *int64) {
	o.Idspecialty = idspecialty
}

// WithStartDate adds the startDate to the search appointment params
func (o *SearchAppointmentParams) WithStartDate(startDate *strfmt.DateTime) *SearchAppointmentParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the search appointment params
func (o *SearchAppointmentParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithStatus adds the status to the search appointment params
func (o *SearchAppointmentParams) WithStatus(status []string) *SearchAppointmentParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the search appointment params
func (o *SearchAppointmentParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAppointmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FinishDate != nil {

		// query param finishDate
		var qrFinishDate strfmt.DateTime
		if o.FinishDate != nil {
			qrFinishDate = *o.FinishDate
		}
		qFinishDate := qrFinishDate.String()
		if qFinishDate != "" {
			if err := r.SetQueryParam("finishDate", qFinishDate); err != nil {
				return err
			}
		}

	}

	if o.IDPatient != nil {

		// query param idPatient
		var qrIDPatient int64
		if o.IDPatient != nil {
			qrIDPatient = *o.IDPatient
		}
		qIDPatient := swag.FormatInt64(qrIDPatient)
		if qIDPatient != "" {
			if err := r.SetQueryParam("idPatient", qIDPatient); err != nil {
				return err
			}
		}

	}

	if o.IDProfessional != nil {

		// query param idProfessional
		var qrIDProfessional int64
		if o.IDProfessional != nil {
			qrIDProfessional = *o.IDProfessional
		}
		qIDProfessional := swag.FormatInt64(qrIDProfessional)
		if qIDProfessional != "" {
			if err := r.SetQueryParam("idProfessional", qIDProfessional); err != nil {
				return err
			}
		}

	}

	if o.Idspecialty != nil {

		// query param idspecialty
		var qrIdspecialty int64
		if o.Idspecialty != nil {
			qrIdspecialty = *o.Idspecialty
		}
		qIdspecialty := swag.FormatInt64(qrIdspecialty)
		if qIdspecialty != "" {
			if err := r.SetQueryParam("idspecialty", qIdspecialty); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.DateTime
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
