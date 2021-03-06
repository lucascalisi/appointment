// Code generated by go-swagger; DO NOT EDIT.

package specialties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new specialties API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for specialties API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SearchSpecialty search specialty API
*/
func (a *Client) SearchSpecialty(params *SearchSpecialtyParams) (*SearchSpecialtyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchSpecialtyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchSpecialty",
		Method:             "GET",
		PathPattern:        "/v1/specialties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchSpecialtyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchSpecialtyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
