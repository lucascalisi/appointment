// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ConfirmAppointmentForPatientHandlerFunc turns a function with the right signature into a confirm appointment for patient handler
type ConfirmAppointmentForPatientHandlerFunc func(ConfirmAppointmentForPatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConfirmAppointmentForPatientHandlerFunc) Handle(params ConfirmAppointmentForPatientParams) middleware.Responder {
	return fn(params)
}

// ConfirmAppointmentForPatientHandler interface for that can handle valid confirm appointment for patient params
type ConfirmAppointmentForPatientHandler interface {
	Handle(ConfirmAppointmentForPatientParams) middleware.Responder
}

// NewConfirmAppointmentForPatient creates a new http.Handler for the confirm appointment for patient operation
func NewConfirmAppointmentForPatient(ctx *middleware.Context, handler ConfirmAppointmentForPatientHandler) *ConfirmAppointmentForPatient {
	return &ConfirmAppointmentForPatient{Context: ctx, Handler: handler}
}

/*ConfirmAppointmentForPatient swagger:route PUT /v1/patients/{id}/appointments/{idAppointment}/confirm patients confirmAppointmentForPatient

ConfirmAppointmentForPatient confirm appointment for patient API

*/
type ConfirmAppointmentForPatient struct {
	Context *middleware.Context
	Handler ConfirmAppointmentForPatientHandler
}

func (o *ConfirmAppointmentForPatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConfirmAppointmentForPatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
