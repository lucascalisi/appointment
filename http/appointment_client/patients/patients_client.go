// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new patients API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for patients API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CancelAppoinmentForPatient cancel appoinment for patient API
*/
func (a *Client) CancelAppoinmentForPatient(params *CancelAppoinmentForPatientParams) (*CancelAppoinmentForPatientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelAppoinmentForPatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancelAppoinmentForPatient",
		Method:             "POST",
		PathPattern:        "/v1/patients/{id}/appointments/{idAppointment}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelAppoinmentForPatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelAppoinmentForPatientOK), nil

}

/*
ConfirmAppointmentForPatient confirm appointment for patient API
*/
func (a *Client) ConfirmAppointmentForPatient(params *ConfirmAppointmentForPatientParams) (*ConfirmAppointmentForPatientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfirmAppointmentForPatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "confirmAppointmentForPatient",
		Method:             "POST",
		PathPattern:        "/v1/patients/{id}/appointments/{idAppointment}/confirm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConfirmAppointmentForPatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConfirmAppointmentForPatientOK), nil

}

/*
GetAppointmentsByPatient get appointments by patient API
*/
func (a *Client) GetAppointmentsByPatient(params *GetAppointmentsByPatientParams) (*GetAppointmentsByPatientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppointmentsByPatientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAppointmentsByPatient",
		Method:             "GET",
		PathPattern:        "/v1/patients/{id}/appointments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAppointmentsByPatientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppointmentsByPatientOK), nil

}

/*
GetPatientbyID get patientby Id API
*/
func (a *Client) GetPatientbyID(params *GetPatientbyIDParams) (*GetPatientbyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPatientbyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPatientbyId",
		Method:             "GET",
		PathPattern:        "/v1/patients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPatientbyIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPatientbyIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}