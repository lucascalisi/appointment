// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AttendAppointmentProfessionalHandlerFunc turns a function with the right signature into a attend appointment professional handler
type AttendAppointmentProfessionalHandlerFunc func(AttendAppointmentProfessionalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AttendAppointmentProfessionalHandlerFunc) Handle(params AttendAppointmentProfessionalParams) middleware.Responder {
	return fn(params)
}

// AttendAppointmentProfessionalHandler interface for that can handle valid attend appointment professional params
type AttendAppointmentProfessionalHandler interface {
	Handle(AttendAppointmentProfessionalParams) middleware.Responder
}

// NewAttendAppointmentProfessional creates a new http.Handler for the attend appointment professional operation
func NewAttendAppointmentProfessional(ctx *middleware.Context, handler AttendAppointmentProfessionalHandler) *AttendAppointmentProfessional {
	return &AttendAppointmentProfessional{Context: ctx, Handler: handler}
}

/*AttendAppointmentProfessional swagger:route PUT /v1/professionals/{id}/appointments/{idappointment}/attend professionals attendAppointmentProfessional

AttendAppointmentProfessional attend appointment professional API

*/
type AttendAppointmentProfessional struct {
	Context *middleware.Context
	Handler AttendAppointmentProfessionalHandler
}

func (o *AttendAppointmentProfessional) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAttendAppointmentProfessionalParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
