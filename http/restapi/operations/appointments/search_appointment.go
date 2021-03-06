// Code generated by go-swagger; DO NOT EDIT.

package appointments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SearchAppointmentHandlerFunc turns a function with the right signature into a search appointment handler
type SearchAppointmentHandlerFunc func(SearchAppointmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchAppointmentHandlerFunc) Handle(params SearchAppointmentParams) middleware.Responder {
	return fn(params)
}

// SearchAppointmentHandler interface for that can handle valid search appointment params
type SearchAppointmentHandler interface {
	Handle(SearchAppointmentParams) middleware.Responder
}

// NewSearchAppointment creates a new http.Handler for the search appointment operation
func NewSearchAppointment(ctx *middleware.Context, handler SearchAppointmentHandler) *SearchAppointment {
	return &SearchAppointment{Context: ctx, Handler: handler}
}

/*SearchAppointment swagger:route GET /v1/appointments appointments searchAppointment

SearchAppointment search appointment API

*/
type SearchAppointment struct {
	Context *middleware.Context
	Handler SearchAppointmentHandler
}

func (o *SearchAppointment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchAppointmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
