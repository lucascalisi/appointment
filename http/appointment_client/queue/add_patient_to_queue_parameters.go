// Code generated by go-swagger; DO NOT EDIT.

package queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/appointment/http/models"
)

// NewAddPatientToQueueParams creates a new AddPatientToQueueParams object
// with the default values initialized.
func NewAddPatientToQueueParams() *AddPatientToQueueParams {
	var ()
	return &AddPatientToQueueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPatientToQueueParamsWithTimeout creates a new AddPatientToQueueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPatientToQueueParamsWithTimeout(timeout time.Duration) *AddPatientToQueueParams {
	var ()
	return &AddPatientToQueueParams{

		timeout: timeout,
	}
}

// NewAddPatientToQueueParamsWithContext creates a new AddPatientToQueueParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPatientToQueueParamsWithContext(ctx context.Context) *AddPatientToQueueParams {
	var ()
	return &AddPatientToQueueParams{

		Context: ctx,
	}
}

// NewAddPatientToQueueParamsWithHTTPClient creates a new AddPatientToQueueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPatientToQueueParamsWithHTTPClient(client *http.Client) *AddPatientToQueueParams {
	var ()
	return &AddPatientToQueueParams{
		HTTPClient: client,
	}
}

/*AddPatientToQueueParams contains all the parameters to send to the API endpoint
for the add patient to queue operation typically these are written to a http.Request
*/
type AddPatientToQueueParams struct {

	/*Item*/
	Item *models.QueueItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add patient to queue params
func (o *AddPatientToQueueParams) WithTimeout(timeout time.Duration) *AddPatientToQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add patient to queue params
func (o *AddPatientToQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add patient to queue params
func (o *AddPatientToQueueParams) WithContext(ctx context.Context) *AddPatientToQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add patient to queue params
func (o *AddPatientToQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add patient to queue params
func (o *AddPatientToQueueParams) WithHTTPClient(client *http.Client) *AddPatientToQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add patient to queue params
func (o *AddPatientToQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItem adds the item to the add patient to queue params
func (o *AddPatientToQueueParams) WithItem(item *models.QueueItem) *AddPatientToQueueParams {
	o.SetItem(item)
	return o
}

// SetItem adds the item to the add patient to queue params
func (o *AddPatientToQueueParams) SetItem(item *models.QueueItem) {
	o.Item = item
}

// WriteToRequest writes these params to a swagger request
func (o *AddPatientToQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Item != nil {
		if err := r.SetBodyParam(o.Item); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
