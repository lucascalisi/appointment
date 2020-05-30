// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RequestAppointmentForPatientHandlerFunc turns a function with the right signature into a request appointment for patient handler
type RequestAppointmentForPatientHandlerFunc func(RequestAppointmentForPatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RequestAppointmentForPatientHandlerFunc) Handle(params RequestAppointmentForPatientParams) middleware.Responder {
	return fn(params)
}

// RequestAppointmentForPatientHandler interface for that can handle valid request appointment for patient params
type RequestAppointmentForPatientHandler interface {
	Handle(RequestAppointmentForPatientParams) middleware.Responder
}

// NewRequestAppointmentForPatient creates a new http.Handler for the request appointment for patient operation
func NewRequestAppointmentForPatient(ctx *middleware.Context, handler RequestAppointmentForPatientHandler) *RequestAppointmentForPatient {
	return &RequestAppointmentForPatient{Context: ctx, Handler: handler}
}

/*RequestAppointmentForPatient swagger:route PUT /v1/patients/{id}/appointments/{idAppointment}/request patients requestAppointmentForPatient

RequestAppointmentForPatient request appointment for patient API

*/
type RequestAppointmentForPatient struct {
	Context *middleware.Context
	Handler RequestAppointmentForPatientHandler
}

func (o *RequestAppointmentForPatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRequestAppointmentForPatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}