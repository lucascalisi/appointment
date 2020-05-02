// Code generated by go-swagger; DO NOT EDIT.

package professionals

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

// NewGetAppointmentByProfessionalAppointmentIDParams creates a new GetAppointmentByProfessionalAppointmentIDParams object
// with the default values initialized.
func NewGetAppointmentByProfessionalAppointmentIDParams() *GetAppointmentByProfessionalAppointmentIDParams {
	var ()
	return &GetAppointmentByProfessionalAppointmentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppointmentByProfessionalAppointmentIDParamsWithTimeout creates a new GetAppointmentByProfessionalAppointmentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppointmentByProfessionalAppointmentIDParamsWithTimeout(timeout time.Duration) *GetAppointmentByProfessionalAppointmentIDParams {
	var ()
	return &GetAppointmentByProfessionalAppointmentIDParams{

		timeout: timeout,
	}
}

// NewGetAppointmentByProfessionalAppointmentIDParamsWithContext creates a new GetAppointmentByProfessionalAppointmentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppointmentByProfessionalAppointmentIDParamsWithContext(ctx context.Context) *GetAppointmentByProfessionalAppointmentIDParams {
	var ()
	return &GetAppointmentByProfessionalAppointmentIDParams{

		Context: ctx,
	}
}

// NewGetAppointmentByProfessionalAppointmentIDParamsWithHTTPClient creates a new GetAppointmentByProfessionalAppointmentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppointmentByProfessionalAppointmentIDParamsWithHTTPClient(client *http.Client) *GetAppointmentByProfessionalAppointmentIDParams {
	var ()
	return &GetAppointmentByProfessionalAppointmentIDParams{
		HTTPClient: client,
	}
}

/*GetAppointmentByProfessionalAppointmentIDParams contains all the parameters to send to the API endpoint
for the get appointment by professional appointment Id operation typically these are written to a http.Request
*/
type GetAppointmentByProfessionalAppointmentIDParams struct {

	/*ID
	  id professional

	*/
	ID int64
	/*Idappointment
	  id appointment

	*/
	Idappointment int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) WithTimeout(timeout time.Duration) *GetAppointmentByProfessionalAppointmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) WithContext(ctx context.Context) *GetAppointmentByProfessionalAppointmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) WithHTTPClient(client *http.Client) *GetAppointmentByProfessionalAppointmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) WithID(id int64) *GetAppointmentByProfessionalAppointmentIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) SetID(id int64) {
	o.ID = id
}

// WithIdappointment adds the idappointment to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) WithIdappointment(idappointment int64) *GetAppointmentByProfessionalAppointmentIDParams {
	o.SetIdappointment(idappointment)
	return o
}

// SetIdappointment adds the idappointment to the get appointment by professional appointment Id params
func (o *GetAppointmentByProfessionalAppointmentIDParams) SetIdappointment(idappointment int64) {
	o.Idappointment = idappointment
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppointmentByProfessionalAppointmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param idappointment
	if err := r.SetPathParam("idappointment", swag.FormatInt64(o.Idappointment)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
