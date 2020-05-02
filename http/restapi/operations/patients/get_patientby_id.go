// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPatientbyIDHandlerFunc turns a function with the right signature into a get patientby Id handler
type GetPatientbyIDHandlerFunc func(GetPatientbyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPatientbyIDHandlerFunc) Handle(params GetPatientbyIDParams) middleware.Responder {
	return fn(params)
}

// GetPatientbyIDHandler interface for that can handle valid get patientby Id params
type GetPatientbyIDHandler interface {
	Handle(GetPatientbyIDParams) middleware.Responder
}

// NewGetPatientbyID creates a new http.Handler for the get patientby Id operation
func NewGetPatientbyID(ctx *middleware.Context, handler GetPatientbyIDHandler) *GetPatientbyID {
	return &GetPatientbyID{Context: ctx, Handler: handler}
}

/*GetPatientbyID swagger:route GET /v1/patients/{id} patients getPatientbyId

GetPatientbyID get patientby Id API

*/
type GetPatientbyID struct {
	Context *middleware.Context
	Handler GetPatientbyIDHandler
}

func (o *GetPatientbyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPatientbyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}