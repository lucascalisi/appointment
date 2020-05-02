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

// NewGetAppointmentsByprofessionalParams creates a new GetAppointmentsByprofessionalParams object
// with the default values initialized.
func NewGetAppointmentsByprofessionalParams() *GetAppointmentsByprofessionalParams {
	var ()
	return &GetAppointmentsByprofessionalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppointmentsByprofessionalParamsWithTimeout creates a new GetAppointmentsByprofessionalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppointmentsByprofessionalParamsWithTimeout(timeout time.Duration) *GetAppointmentsByprofessionalParams {
	var ()
	return &GetAppointmentsByprofessionalParams{

		timeout: timeout,
	}
}

// NewGetAppointmentsByprofessionalParamsWithContext creates a new GetAppointmentsByprofessionalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppointmentsByprofessionalParamsWithContext(ctx context.Context) *GetAppointmentsByprofessionalParams {
	var ()
	return &GetAppointmentsByprofessionalParams{

		Context: ctx,
	}
}

// NewGetAppointmentsByprofessionalParamsWithHTTPClient creates a new GetAppointmentsByprofessionalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppointmentsByprofessionalParamsWithHTTPClient(client *http.Client) *GetAppointmentsByprofessionalParams {
	var ()
	return &GetAppointmentsByprofessionalParams{
		HTTPClient: client,
	}
}

/*GetAppointmentsByprofessionalParams contains all the parameters to send to the API endpoint
for the get appointments byprofessional operation typically these are written to a http.Request
*/
type GetAppointmentsByprofessionalParams struct {

	/*ID
	  id professional

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) WithTimeout(timeout time.Duration) *GetAppointmentsByprofessionalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) WithContext(ctx context.Context) *GetAppointmentsByprofessionalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) WithHTTPClient(client *http.Client) *GetAppointmentsByprofessionalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) WithID(id int64) *GetAppointmentsByprofessionalParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get appointments byprofessional params
func (o *GetAppointmentsByprofessionalParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppointmentsByprofessionalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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