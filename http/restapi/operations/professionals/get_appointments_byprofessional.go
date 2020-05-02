// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAppointmentsByprofessionalHandlerFunc turns a function with the right signature into a get appointments byprofessional handler
type GetAppointmentsByprofessionalHandlerFunc func(GetAppointmentsByprofessionalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppointmentsByprofessionalHandlerFunc) Handle(params GetAppointmentsByprofessionalParams) middleware.Responder {
	return fn(params)
}

// GetAppointmentsByprofessionalHandler interface for that can handle valid get appointments byprofessional params
type GetAppointmentsByprofessionalHandler interface {
	Handle(GetAppointmentsByprofessionalParams) middleware.Responder
}

// NewGetAppointmentsByprofessional creates a new http.Handler for the get appointments byprofessional operation
func NewGetAppointmentsByprofessional(ctx *middleware.Context, handler GetAppointmentsByprofessionalHandler) *GetAppointmentsByprofessional {
	return &GetAppointmentsByprofessional{Context: ctx, Handler: handler}
}

/*GetAppointmentsByprofessional swagger:route GET /v1/professionals/{id}/appointments professionals getAppointmentsByprofessional

GetAppointmentsByprofessional get appointments byprofessional API

*/
type GetAppointmentsByprofessional struct {
	Context *middleware.Context
	Handler GetAppointmentsByprofessionalHandler
}

func (o *GetAppointmentsByprofessional) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAppointmentsByprofessionalParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
