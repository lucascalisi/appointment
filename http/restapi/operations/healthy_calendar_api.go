// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/appointment/http/restapi/operations/appointments"
	"github.com/appointment/http/restapi/operations/login"
	"github.com/appointment/http/restapi/operations/patients"
	"github.com/appointment/http/restapi/operations/professionals"
	"github.com/appointment/http/restapi/operations/specialties"
)

// NewHealthyCalendarAPI creates a new HealthyCalendar instance
func NewHealthyCalendarAPI(spec *loads.Document) *HealthyCalendarAPI {
	return &HealthyCalendarAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ProfessionalsAttendAppointmentProfessionalHandler: professionals.AttendAppointmentProfessionalHandlerFunc(func(params professionals.AttendAppointmentProfessionalParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsAttendAppointmentProfessional has not yet been implemented")
		}),
		PatientsCancelAppoinmentForPatientHandler: patients.CancelAppoinmentForPatientHandlerFunc(func(params patients.CancelAppoinmentForPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation PatientsCancelAppoinmentForPatient has not yet been implemented")
		}),
		ProfessionalsCancelAppointmentProfessionalHandler: professionals.CancelAppointmentProfessionalHandlerFunc(func(params professionals.CancelAppointmentProfessionalParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsCancelAppointmentProfessional has not yet been implemented")
		}),
		PatientsConfirmAppointmentForPatientHandler: patients.ConfirmAppointmentForPatientHandlerFunc(func(params patients.ConfirmAppointmentForPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation PatientsConfirmAppointmentForPatient has not yet been implemented")
		}),
		ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler: professionals.GetAppointmentByProfessionalAppointmentIDHandlerFunc(func(params professionals.GetAppointmentByProfessionalAppointmentIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsGetAppointmentByProfessionalAppointmentID has not yet been implemented")
		}),
		PatientsGetAppointmentsByPatientHandler: patients.GetAppointmentsByPatientHandlerFunc(func(params patients.GetAppointmentsByPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation PatientsGetAppointmentsByPatient has not yet been implemented")
		}),
		ProfessionalsGetAppointmentsByprofessionalHandler: professionals.GetAppointmentsByprofessionalHandlerFunc(func(params professionals.GetAppointmentsByprofessionalParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsGetAppointmentsByprofessional has not yet been implemented")
		}),
		PatientsGetPatientbyIDHandler: patients.GetPatientbyIDHandlerFunc(func(params patients.GetPatientbyIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PatientsGetPatientbyID has not yet been implemented")
		}),
		ProfessionalsGetProfesionalScheduleBySpecialtyHandler: professionals.GetProfesionalScheduleBySpecialtyHandlerFunc(func(params professionals.GetProfesionalScheduleBySpecialtyParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsGetProfesionalScheduleBySpecialty has not yet been implemented")
		}),
		ProfessionalsGetProfessionalbyIDHandler: professionals.GetProfessionalbyIDHandlerFunc(func(params professionals.GetProfessionalbyIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsGetProfessionalbyID has not yet been implemented")
		}),
		ProfessionalsGetSpecialtiesByProfessionalHandler: professionals.GetSpecialtiesByProfessionalHandlerFunc(func(params professionals.GetSpecialtiesByProfessionalParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsGetSpecialtiesByProfessional has not yet been implemented")
		}),
		LoginLoginHandler: login.LoginHandlerFunc(func(params login.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation LoginLogin has not yet been implemented")
		}),
		PatientsRequestAppointmentForPatientHandler: patients.RequestAppointmentForPatientHandlerFunc(func(params patients.RequestAppointmentForPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation PatientsRequestAppointmentForPatient has not yet been implemented")
		}),
		AppointmentsSearchAppointmentHandler: appointments.SearchAppointmentHandlerFunc(func(params appointments.SearchAppointmentParams) middleware.Responder {
			return middleware.NotImplemented("operation AppointmentsSearchAppointment has not yet been implemented")
		}),
		SpecialtiesSearchSpecialtyHandler: specialties.SearchSpecialtyHandlerFunc(func(params specialties.SearchSpecialtyParams) middleware.Responder {
			return middleware.NotImplemented("operation SpecialtiesSearchSpecialty has not yet been implemented")
		}),
		ProfessionalsSetProfesionalScheduleBySpecialtyHandler: professionals.SetProfesionalScheduleBySpecialtyHandlerFunc(func(params professionals.SetProfesionalScheduleBySpecialtyParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfessionalsSetProfesionalScheduleBySpecialty has not yet been implemented")
		}),
	}
}

/*HealthyCalendarAPI API to TPO UADE (Aplicaciones Distribuidas) */
type HealthyCalendarAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ProfessionalsAttendAppointmentProfessionalHandler sets the operation handler for the attend appointment professional operation
	ProfessionalsAttendAppointmentProfessionalHandler professionals.AttendAppointmentProfessionalHandler
	// PatientsCancelAppoinmentForPatientHandler sets the operation handler for the cancel appoinment for patient operation
	PatientsCancelAppoinmentForPatientHandler patients.CancelAppoinmentForPatientHandler
	// ProfessionalsCancelAppointmentProfessionalHandler sets the operation handler for the cancel appointment professional operation
	ProfessionalsCancelAppointmentProfessionalHandler professionals.CancelAppointmentProfessionalHandler
	// PatientsConfirmAppointmentForPatientHandler sets the operation handler for the confirm appointment for patient operation
	PatientsConfirmAppointmentForPatientHandler patients.ConfirmAppointmentForPatientHandler
	// ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler sets the operation handler for the get appointment by professional appointment Id operation
	ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler professionals.GetAppointmentByProfessionalAppointmentIDHandler
	// PatientsGetAppointmentsByPatientHandler sets the operation handler for the get appointments by patient operation
	PatientsGetAppointmentsByPatientHandler patients.GetAppointmentsByPatientHandler
	// ProfessionalsGetAppointmentsByprofessionalHandler sets the operation handler for the get appointments byprofessional operation
	ProfessionalsGetAppointmentsByprofessionalHandler professionals.GetAppointmentsByprofessionalHandler
	// PatientsGetPatientbyIDHandler sets the operation handler for the get patientby Id operation
	PatientsGetPatientbyIDHandler patients.GetPatientbyIDHandler
	// ProfessionalsGetProfesionalScheduleBySpecialtyHandler sets the operation handler for the get profesional schedule by specialty operation
	ProfessionalsGetProfesionalScheduleBySpecialtyHandler professionals.GetProfesionalScheduleBySpecialtyHandler
	// ProfessionalsGetProfessionalbyIDHandler sets the operation handler for the get professionalby Id operation
	ProfessionalsGetProfessionalbyIDHandler professionals.GetProfessionalbyIDHandler
	// ProfessionalsGetSpecialtiesByProfessionalHandler sets the operation handler for the get specialties by professional operation
	ProfessionalsGetSpecialtiesByProfessionalHandler professionals.GetSpecialtiesByProfessionalHandler
	// LoginLoginHandler sets the operation handler for the login operation
	LoginLoginHandler login.LoginHandler
	// PatientsRequestAppointmentForPatientHandler sets the operation handler for the request appointment for patient operation
	PatientsRequestAppointmentForPatientHandler patients.RequestAppointmentForPatientHandler
	// AppointmentsSearchAppointmentHandler sets the operation handler for the search appointment operation
	AppointmentsSearchAppointmentHandler appointments.SearchAppointmentHandler
	// SpecialtiesSearchSpecialtyHandler sets the operation handler for the search specialty operation
	SpecialtiesSearchSpecialtyHandler specialties.SearchSpecialtyHandler
	// ProfessionalsSetProfesionalScheduleBySpecialtyHandler sets the operation handler for the set profesional schedule by specialty operation
	ProfessionalsSetProfesionalScheduleBySpecialtyHandler professionals.SetProfesionalScheduleBySpecialtyHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *HealthyCalendarAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *HealthyCalendarAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *HealthyCalendarAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *HealthyCalendarAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *HealthyCalendarAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *HealthyCalendarAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *HealthyCalendarAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the HealthyCalendarAPI
func (o *HealthyCalendarAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ProfessionalsAttendAppointmentProfessionalHandler == nil {
		unregistered = append(unregistered, "professionals.AttendAppointmentProfessionalHandler")
	}

	if o.PatientsCancelAppoinmentForPatientHandler == nil {
		unregistered = append(unregistered, "patients.CancelAppoinmentForPatientHandler")
	}

	if o.ProfessionalsCancelAppointmentProfessionalHandler == nil {
		unregistered = append(unregistered, "professionals.CancelAppointmentProfessionalHandler")
	}

	if o.PatientsConfirmAppointmentForPatientHandler == nil {
		unregistered = append(unregistered, "patients.ConfirmAppointmentForPatientHandler")
	}

	if o.ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler == nil {
		unregistered = append(unregistered, "professionals.GetAppointmentByProfessionalAppointmentIDHandler")
	}

	if o.PatientsGetAppointmentsByPatientHandler == nil {
		unregistered = append(unregistered, "patients.GetAppointmentsByPatientHandler")
	}

	if o.ProfessionalsGetAppointmentsByprofessionalHandler == nil {
		unregistered = append(unregistered, "professionals.GetAppointmentsByprofessionalHandler")
	}

	if o.PatientsGetPatientbyIDHandler == nil {
		unregistered = append(unregistered, "patients.GetPatientbyIDHandler")
	}

	if o.ProfessionalsGetProfesionalScheduleBySpecialtyHandler == nil {
		unregistered = append(unregistered, "professionals.GetProfesionalScheduleBySpecialtyHandler")
	}

	if o.ProfessionalsGetProfessionalbyIDHandler == nil {
		unregistered = append(unregistered, "professionals.GetProfessionalbyIDHandler")
	}

	if o.ProfessionalsGetSpecialtiesByProfessionalHandler == nil {
		unregistered = append(unregistered, "professionals.GetSpecialtiesByProfessionalHandler")
	}

	if o.LoginLoginHandler == nil {
		unregistered = append(unregistered, "login.LoginHandler")
	}

	if o.PatientsRequestAppointmentForPatientHandler == nil {
		unregistered = append(unregistered, "patients.RequestAppointmentForPatientHandler")
	}

	if o.AppointmentsSearchAppointmentHandler == nil {
		unregistered = append(unregistered, "appointments.SearchAppointmentHandler")
	}

	if o.SpecialtiesSearchSpecialtyHandler == nil {
		unregistered = append(unregistered, "specialties.SearchSpecialtyHandler")
	}

	if o.ProfessionalsSetProfesionalScheduleBySpecialtyHandler == nil {
		unregistered = append(unregistered, "professionals.SetProfesionalScheduleBySpecialtyHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *HealthyCalendarAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *HealthyCalendarAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *HealthyCalendarAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *HealthyCalendarAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *HealthyCalendarAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *HealthyCalendarAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the healthy calendar API
func (o *HealthyCalendarAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *HealthyCalendarAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/professionals/{id}/appointments/{idappointment}/attend"] = professionals.NewAttendAppointmentProfessional(o.context, o.ProfessionalsAttendAppointmentProfessionalHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/patients/{id}/appointments/{idAppointment}/cancel"] = patients.NewCancelAppoinmentForPatient(o.context, o.PatientsCancelAppoinmentForPatientHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/professionals/{id}/appointments/{idappointment}/cancel"] = professionals.NewCancelAppointmentProfessional(o.context, o.ProfessionalsCancelAppointmentProfessionalHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/patients/{id}/appointments/{idAppointment}/confirm"] = patients.NewConfirmAppointmentForPatient(o.context, o.PatientsConfirmAppointmentForPatientHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/professionals/{id}/appointments/{idappointment}"] = professionals.NewGetAppointmentByProfessionalAppointmentID(o.context, o.ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/patients/{id}/appointments"] = patients.NewGetAppointmentsByPatient(o.context, o.PatientsGetAppointmentsByPatientHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/professionals/{id}/appointments"] = professionals.NewGetAppointmentsByprofessional(o.context, o.ProfessionalsGetAppointmentsByprofessionalHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/patients/{id}"] = patients.NewGetPatientbyID(o.context, o.PatientsGetPatientbyIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/professionals/{id}/specialties/{idSpecialty}/schedule"] = professionals.NewGetProfesionalScheduleBySpecialty(o.context, o.ProfessionalsGetProfesionalScheduleBySpecialtyHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/professionals/{id}"] = professionals.NewGetProfessionalbyID(o.context, o.ProfessionalsGetProfessionalbyIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/professionals/{id}/specialties"] = professionals.NewGetSpecialtiesByProfessional(o.context, o.ProfessionalsGetSpecialtiesByProfessionalHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/login"] = login.NewLogin(o.context, o.LoginLoginHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/patients/{id}/appointments/{idAppointment}/request"] = patients.NewRequestAppointmentForPatient(o.context, o.PatientsRequestAppointmentForPatientHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/appointments"] = appointments.NewSearchAppointment(o.context, o.AppointmentsSearchAppointmentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/specialties"] = specialties.NewSearchSpecialty(o.context, o.SpecialtiesSearchSpecialtyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/professionals/{id}/specialties/{idSpecialty}/schedule"] = professionals.NewSetProfesionalScheduleBySpecialty(o.context, o.ProfessionalsSetProfesionalScheduleBySpecialtyHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *HealthyCalendarAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *HealthyCalendarAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *HealthyCalendarAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *HealthyCalendarAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}