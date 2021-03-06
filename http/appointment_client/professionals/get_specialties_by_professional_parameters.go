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

// NewGetSpecialtiesByProfessionalParams creates a new GetSpecialtiesByProfessionalParams object
// with the default values initialized.
func NewGetSpecialtiesByProfessionalParams() *GetSpecialtiesByProfessionalParams {
	var ()
	return &GetSpecialtiesByProfessionalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpecialtiesByProfessionalParamsWithTimeout creates a new GetSpecialtiesByProfessionalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpecialtiesByProfessionalParamsWithTimeout(timeout time.Duration) *GetSpecialtiesByProfessionalParams {
	var ()
	return &GetSpecialtiesByProfessionalParams{

		timeout: timeout,
	}
}

// NewGetSpecialtiesByProfessionalParamsWithContext creates a new GetSpecialtiesByProfessionalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpecialtiesByProfessionalParamsWithContext(ctx context.Context) *GetSpecialtiesByProfessionalParams {
	var ()
	return &GetSpecialtiesByProfessionalParams{

		Context: ctx,
	}
}

// NewGetSpecialtiesByProfessionalParamsWithHTTPClient creates a new GetSpecialtiesByProfessionalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpecialtiesByProfessionalParamsWithHTTPClient(client *http.Client) *GetSpecialtiesByProfessionalParams {
	var ()
	return &GetSpecialtiesByProfessionalParams{
		HTTPClient: client,
	}
}

/*GetSpecialtiesByProfessionalParams contains all the parameters to send to the API endpoint
for the get specialties by professional operation typically these are written to a http.Request
*/
type GetSpecialtiesByProfessionalParams struct {

	/*ID
	  id professional

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) WithTimeout(timeout time.Duration) *GetSpecialtiesByProfessionalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) WithContext(ctx context.Context) *GetSpecialtiesByProfessionalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) WithHTTPClient(client *http.Client) *GetSpecialtiesByProfessionalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) WithID(id int64) *GetSpecialtiesByProfessionalParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get specialties by professional params
func (o *GetSpecialtiesByProfessionalParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpecialtiesByProfessionalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
