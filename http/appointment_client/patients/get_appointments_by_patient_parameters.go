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

// NewGetAppointmentsByPatientParams creates a new GetAppointmentsByPatientParams object
// with the default values initialized.
func NewGetAppointmentsByPatientParams() *GetAppointmentsByPatientParams {
	var ()
	return &GetAppointmentsByPatientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppointmentsByPatientParamsWithTimeout creates a new GetAppointmentsByPatientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppointmentsByPatientParamsWithTimeout(timeout time.Duration) *GetAppointmentsByPatientParams {
	var ()
	return &GetAppointmentsByPatientParams{

		timeout: timeout,
	}
}

// NewGetAppointmentsByPatientParamsWithContext creates a new GetAppointmentsByPatientParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppointmentsByPatientParamsWithContext(ctx context.Context) *GetAppointmentsByPatientParams {
	var ()
	return &GetAppointmentsByPatientParams{

		Context: ctx,
	}
}

// NewGetAppointmentsByPatientParamsWithHTTPClient creates a new GetAppointmentsByPatientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppointmentsByPatientParamsWithHTTPClient(client *http.Client) *GetAppointmentsByPatientParams {
	var ()
	return &GetAppointmentsByPatientParams{
		HTTPClient: client,
	}
}

/*GetAppointmentsByPatientParams contains all the parameters to send to the API endpoint
for the get appointments by patient operation typically these are written to a http.Request
*/
type GetAppointmentsByPatientParams struct {

	/*ID
	  id patient

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) WithTimeout(timeout time.Duration) *GetAppointmentsByPatientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) WithContext(ctx context.Context) *GetAppointmentsByPatientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) WithHTTPClient(client *http.Client) *GetAppointmentsByPatientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) WithID(id int64) *GetAppointmentsByPatientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get appointments by patient params
func (o *GetAppointmentsByPatientParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppointmentsByPatientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}