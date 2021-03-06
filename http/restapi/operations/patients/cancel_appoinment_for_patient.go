// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CancelAppoinmentForPatientHandlerFunc turns a function with the right signature into a cancel appoinment for patient handler
type CancelAppoinmentForPatientHandlerFunc func(CancelAppoinmentForPatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelAppoinmentForPatientHandlerFunc) Handle(params CancelAppoinmentForPatientParams) middleware.Responder {
	return fn(params)
}

// CancelAppoinmentForPatientHandler interface for that can handle valid cancel appoinment for patient params
type CancelAppoinmentForPatientHandler interface {
	Handle(CancelAppoinmentForPatientParams) middleware.Responder
}

// NewCancelAppoinmentForPatient creates a new http.Handler for the cancel appoinment for patient operation
func NewCancelAppoinmentForPatient(ctx *middleware.Context, handler CancelAppoinmentForPatientHandler) *CancelAppoinmentForPatient {
	return &CancelAppoinmentForPatient{Context: ctx, Handler: handler}
}

/*CancelAppoinmentForPatient swagger:route PUT /v1/patients/{id}/appointments/{idAppointment}/cancel patients cancelAppoinmentForPatient

CancelAppoinmentForPatient cancel appoinment for patient API

*/
type CancelAppoinmentForPatient struct {
	Context *middleware.Context
	Handler CancelAppoinmentForPatientHandler
}

func (o *CancelAppoinmentForPatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelAppoinmentForPatientParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
