// Code generated by go-swagger; DO NOT EDIT.

package appointment_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/appointment/http/appointment_client/appointments"
	"github.com/appointment/http/appointment_client/login"
	"github.com/appointment/http/appointment_client/patients"
	"github.com/appointment/http/appointment_client/professionals"
	"github.com/appointment/http/appointment_client/queue"
	"github.com/appointment/http/appointment_client/specialties"
)

// Default healthy calendar HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new healthy calendar HTTP client.
func NewHTTPClient(formats strfmt.Registry) *HealthyCalendar {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new healthy calendar HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *HealthyCalendar {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new healthy calendar client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *HealthyCalendar {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(HealthyCalendar)
	cli.Transport = transport

	cli.Appointments = appointments.New(transport, formats)

	cli.Login = login.New(transport, formats)

	cli.Patients = patients.New(transport, formats)

	cli.Professionals = professionals.New(transport, formats)

	cli.Queue = queue.New(transport, formats)

	cli.Specialties = specialties.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// HealthyCalendar is a client for healthy calendar
type HealthyCalendar struct {
	Appointments *appointments.Client

	Login *login.Client

	Patients *patients.Client

	Professionals *professionals.Client

	Queue *queue.Client

	Specialties *specialties.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *HealthyCalendar) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Appointments.SetTransport(transport)

	c.Login.SetTransport(transport)

	c.Patients.SetTransport(transport)

	c.Professionals.SetTransport(transport)

	c.Queue.SetTransport(transport)

	c.Specialties.SetTransport(transport)

}
