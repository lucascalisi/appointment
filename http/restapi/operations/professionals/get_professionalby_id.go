// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetProfessionalbyIDHandlerFunc turns a function with the right signature into a get professionalby Id handler
type GetProfessionalbyIDHandlerFunc func(GetProfessionalbyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProfessionalbyIDHandlerFunc) Handle(params GetProfessionalbyIDParams) middleware.Responder {
	return fn(params)
}

// GetProfessionalbyIDHandler interface for that can handle valid get professionalby Id params
type GetProfessionalbyIDHandler interface {
	Handle(GetProfessionalbyIDParams) middleware.Responder
}

// NewGetProfessionalbyID creates a new http.Handler for the get professionalby Id operation
func NewGetProfessionalbyID(ctx *middleware.Context, handler GetProfessionalbyIDHandler) *GetProfessionalbyID {
	return &GetProfessionalbyID{Context: ctx, Handler: handler}
}

/*GetProfessionalbyID swagger:route GET /v1/professionals/{id} professionals getProfessionalbyId

GetProfessionalbyID get professionalby Id API

*/
type GetProfessionalbyID struct {
	Context *middleware.Context
	Handler GetProfessionalbyIDHandler
}

func (o *GetProfessionalbyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProfessionalbyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
