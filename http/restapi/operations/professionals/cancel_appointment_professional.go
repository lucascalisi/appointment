// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CancelAppointmentProfessionalHandlerFunc turns a function with the right signature into a cancel appointment professional handler
type CancelAppointmentProfessionalHandlerFunc func(CancelAppointmentProfessionalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelAppointmentProfessionalHandlerFunc) Handle(params CancelAppointmentProfessionalParams) middleware.Responder {
	return fn(params)
}

// CancelAppointmentProfessionalHandler interface for that can handle valid cancel appointment professional params
type CancelAppointmentProfessionalHandler interface {
	Handle(CancelAppointmentProfessionalParams) middleware.Responder
}

// NewCancelAppointmentProfessional creates a new http.Handler for the cancel appointment professional operation
func NewCancelAppointmentProfessional(ctx *middleware.Context, handler CancelAppointmentProfessionalHandler) *CancelAppointmentProfessional {
	return &CancelAppointmentProfessional{Context: ctx, Handler: handler}
}

/*CancelAppointmentProfessional swagger:route POST /v1/professionals/{id}/appointments/{idappointment}/cancel professionals cancelAppointmentProfessional

CancelAppointmentProfessional cancel appointment professional API

*/
type CancelAppointmentProfessional struct {
	Context *middleware.Context
	Handler CancelAppointmentProfessionalHandler
}

func (o *CancelAppointmentProfessional) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelAppointmentProfessionalParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}