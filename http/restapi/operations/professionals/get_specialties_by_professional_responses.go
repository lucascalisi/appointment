// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/appointment/http/models"
)

// GetSpecialtiesByProfessionalOKCode is the HTTP code returned for type GetSpecialtiesByProfessionalOK
const GetSpecialtiesByProfessionalOKCode int = 200

/*GetSpecialtiesByProfessionalOK resultado specialties professional

swagger:response getSpecialtiesByProfessionalOK
*/
type GetSpecialtiesByProfessionalOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Specialty `json:"body,omitempty"`
}

// NewGetSpecialtiesByProfessionalOK creates GetSpecialtiesByProfessionalOK with default headers values
func NewGetSpecialtiesByProfessionalOK() *GetSpecialtiesByProfessionalOK {

	return &GetSpecialtiesByProfessionalOK{}
}

// WithPayload adds the payload to the get specialties by professional o k response
func (o *GetSpecialtiesByProfessionalOK) WithPayload(payload []*models.Specialty) *GetSpecialtiesByProfessionalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get specialties by professional o k response
func (o *GetSpecialtiesByProfessionalOK) SetPayload(payload []*models.Specialty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpecialtiesByProfessionalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Specialty, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSpecialtiesByProfessionalInternalServerErrorCode is the HTTP code returned for type GetSpecialtiesByProfessionalInternalServerError
const GetSpecialtiesByProfessionalInternalServerErrorCode int = 500

/*GetSpecialtiesByProfessionalInternalServerError internal server error

swagger:response getSpecialtiesByProfessionalInternalServerError
*/
type GetSpecialtiesByProfessionalInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpecialtiesByProfessionalInternalServerError creates GetSpecialtiesByProfessionalInternalServerError with default headers values
func NewGetSpecialtiesByProfessionalInternalServerError() *GetSpecialtiesByProfessionalInternalServerError {

	return &GetSpecialtiesByProfessionalInternalServerError{}
}

// WithPayload adds the payload to the get specialties by professional internal server error response
func (o *GetSpecialtiesByProfessionalInternalServerError) WithPayload(payload *models.Error) *GetSpecialtiesByProfessionalInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get specialties by professional internal server error response
func (o *GetSpecialtiesByProfessionalInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpecialtiesByProfessionalInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
