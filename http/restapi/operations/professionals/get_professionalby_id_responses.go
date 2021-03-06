// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/appointment/http/models"
)

// GetProfessionalbyIDOKCode is the HTTP code returned for type GetProfessionalbyIDOK
const GetProfessionalbyIDOKCode int = 200

/*GetProfessionalbyIDOK result professional

swagger:response getProfessionalbyIdOK
*/
type GetProfessionalbyIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Professional `json:"body,omitempty"`
}

// NewGetProfessionalbyIDOK creates GetProfessionalbyIDOK with default headers values
func NewGetProfessionalbyIDOK() *GetProfessionalbyIDOK {

	return &GetProfessionalbyIDOK{}
}

// WithPayload adds the payload to the get professionalby Id o k response
func (o *GetProfessionalbyIDOK) WithPayload(payload *models.Professional) *GetProfessionalbyIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get professionalby Id o k response
func (o *GetProfessionalbyIDOK) SetPayload(payload *models.Professional) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfessionalbyIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProfessionalbyIDInternalServerErrorCode is the HTTP code returned for type GetProfessionalbyIDInternalServerError
const GetProfessionalbyIDInternalServerErrorCode int = 500

/*GetProfessionalbyIDInternalServerError internal server error

swagger:response getProfessionalbyIdInternalServerError
*/
type GetProfessionalbyIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProfessionalbyIDInternalServerError creates GetProfessionalbyIDInternalServerError with default headers values
func NewGetProfessionalbyIDInternalServerError() *GetProfessionalbyIDInternalServerError {

	return &GetProfessionalbyIDInternalServerError{}
}

// WithPayload adds the payload to the get professionalby Id internal server error response
func (o *GetProfessionalbyIDInternalServerError) WithPayload(payload *models.Error) *GetProfessionalbyIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get professionalby Id internal server error response
func (o *GetProfessionalbyIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfessionalbyIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
