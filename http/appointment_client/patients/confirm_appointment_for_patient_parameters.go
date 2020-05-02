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

// NewConfirmAppointmentForPatientParams creates a new ConfirmAppointmentForPatientParams object
// with the default values initialized.
func NewConfirmAppointmentForPatientParams() *ConfirmAppointmentForPatientParams {
	var ()
	return &ConfirmAppointmentForPatientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmAppointmentForPatientParamsWithTimeout creates a new ConfirmAppointmentForPatientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfirmAppointmentForPatientParamsWithTimeout(timeout time.Duration) *ConfirmAppointmentForPatientParams {
	var ()
	return &ConfirmAppointmentForPatientParams{

		timeout: timeout,
	}
}

// NewConfirmAppointmentForPatientParamsWithContext creates a new ConfirmAppointmentForPatientParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfirmAppointmentForPatientParamsWithContext(ctx context.Context) *ConfirmAppointmentForPatientParams {
	var ()
	return &ConfirmAppointmentForPatientParams{

		Context: ctx,
	}
}

// NewConfirmAppointmentForPatientParamsWithHTTPClient creates a new ConfirmAppointmentForPatientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfirmAppointmentForPatientParamsWithHTTPClient(client *http.Client) *ConfirmAppointmentForPatientParams {
	var ()
	return &ConfirmAppointmentForPatientParams{
		HTTPClient: client,
	}
}

/*ConfirmAppointmentForPatientParams contains all the parameters to send to the API endpoint
for the confirm appointment for patient operation typically these are written to a http.Request
*/
type ConfirmAppointmentForPatientParams struct {

	/*ID
	  id patient

	*/
	ID int64
	/*IDAppointment
	  confirm appoinment

	*/
	IDAppointment int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) WithTimeout(timeout time.Duration) *ConfirmAppointmentForPatientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) WithContext(ctx context.Context) *ConfirmAppointmentForPatientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) WithHTTPClient(client *http.Client) *ConfirmAppointmentForPatientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) WithID(id int64) *ConfirmAppointmentForPatientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) SetID(id int64) {
	o.ID = id
}

// WithIDAppointment adds the iDAppointment to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) WithIDAppointment(iDAppointment int64) *ConfirmAppointmentForPatientParams {
	o.SetIDAppointment(iDAppointment)
	return o
}

// SetIDAppointment adds the idAppointment to the confirm appointment for patient params
func (o *ConfirmAppointmentForPatientParams) SetIDAppointment(iDAppointment int64) {
	o.IDAppointment = iDAppointment
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmAppointmentForPatientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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