// Code generated by go-swagger; DO NOT EDIT.

package patients

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

// NewRequestAppointmentForPatientParams creates a new RequestAppointmentForPatientParams object
// with the default values initialized.
func NewRequestAppointmentForPatientParams() *RequestAppointmentForPatientParams {
	var ()
	return &RequestAppointmentForPatientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestAppointmentForPatientParamsWithTimeout creates a new RequestAppointmentForPatientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestAppointmentForPatientParamsWithTimeout(timeout time.Duration) *RequestAppointmentForPatientParams {
	var ()
	return &RequestAppointmentForPatientParams{

		timeout: timeout,
	}
}

// NewRequestAppointmentForPatientParamsWithContext creates a new RequestAppointmentForPatientParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestAppointmentForPatientParamsWithContext(ctx context.Context) *RequestAppointmentForPatientParams {
	var ()
	return &RequestAppointmentForPatientParams{

		Context: ctx,
	}
}

// NewRequestAppointmentForPatientParamsWithHTTPClient creates a new RequestAppointmentForPatientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestAppointmentForPatientParamsWithHTTPClient(client *http.Client) *RequestAppointmentForPatientParams {
	var ()
	return &RequestAppointmentForPatientParams{
		HTTPClient: client,
	}
}

/*RequestAppointmentForPatientParams contains all the parameters to send to the API endpoint
for the request appointment for patient operation typically these are written to a http.Request
*/
type RequestAppointmentForPatientParams struct {

	/*ID
	  id patient

	*/
	ID int64
	/*IDAppointment
	  request appoinment

	*/
	IDAppointment int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) WithTimeout(timeout time.Duration) *RequestAppointmentForPatientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) WithContext(ctx context.Context) *RequestAppointmentForPatientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) WithHTTPClient(client *http.Client) *RequestAppointmentForPatientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) WithID(id int64) *RequestAppointmentForPatientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) SetID(id int64) {
	o.ID = id
}

// WithIDAppointment adds the iDAppointment to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) WithIDAppointment(iDAppointment int64) *RequestAppointmentForPatientParams {
	o.SetIDAppointment(iDAppointment)
	return o
}

// SetIDAppointment adds the idAppointment to the request appointment for patient params
func (o *RequestAppointmentForPatientParams) SetIDAppointment(iDAppointment int64) {
	o.IDAppointment = iDAppointment
}

// WriteToRequest writes these params to a swagger request
func (o *RequestAppointmentForPatientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param idAppointment
	if err := r.SetPathParam("idAppointment", swag.FormatInt64(o.IDAppointment)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}