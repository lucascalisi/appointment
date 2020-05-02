// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSpecialtiesByProfessionalHandlerFunc turns a function with the right signature into a get specialties by professional handler
type GetSpecialtiesByProfessionalHandlerFunc func(GetSpecialtiesByProfessionalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpecialtiesByProfessionalHandlerFunc) Handle(params GetSpecialtiesByProfessionalParams) middleware.Responder {
	return fn(params)
}

// GetSpecialtiesByProfessionalHandler interface for that can handle valid get specialties by professional params
type GetSpecialtiesByProfessionalHandler interface {
	Handle(GetSpecialtiesByProfessionalParams) middleware.Responder
}

// NewGetSpecialtiesByProfessional creates a new http.Handler for the get specialties by professional operation
func NewGetSpecialtiesByProfessional(ctx *middleware.Context, handler GetSpecialtiesByProfessionalHandler) *GetSpecialtiesByProfessional {
	return &GetSpecialtiesByProfessional{Context: ctx, Handler: handler}
}

/*GetSpecialtiesByProfessional swagger:route GET /v1/professionals/{id}/specialties professionals getSpecialtiesByProfessional

GetSpecialtiesByProfessional get specialties by professional API

*/
type GetSpecialtiesByProfessional struct {
	Context *middleware.Context
	Handler GetSpecialtiesByProfessionalHandler
}

func (o *GetSpecialtiesByProfessional) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSpecialtiesByProfessionalParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
