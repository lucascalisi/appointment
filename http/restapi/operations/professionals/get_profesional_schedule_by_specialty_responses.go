// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/appointment/http/models"
)

// GetProfesionalScheduleBySpecialtyOKCode is the HTTP code returned for type GetProfesionalScheduleBySpecialtyOK
const GetProfesionalScheduleBySpecialtyOKCode int = 200

/*GetProfesionalScheduleBySpecialtyOK resultado schedule professional by specialty

swagger:response getProfesionalScheduleBySpecialtyOK
*/
type GetProfesionalScheduleBySpecialtyOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Schedule `json:"body,omitempty"`
}

// NewGetProfesionalScheduleBySpecialtyOK creates GetProfesionalScheduleBySpecialtyOK with default headers values
func NewGetProfesionalScheduleBySpecialtyOK() *GetProfesionalScheduleBySpecialtyOK {

	return &GetProfesionalScheduleBySpecialtyOK{}
}

// WithPayload adds the payload to the get profesional schedule by specialty o k response
func (o *GetProfesionalScheduleBySpecialtyOK) WithPayload(payload []*models.Schedule) *GetProfesionalScheduleBySpecialtyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profesional schedule by specialty o k response
func (o *GetProfesionalScheduleBySpecialtyOK) SetPayload(payload []*models.Schedule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfesionalScheduleBySpecialtyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Schedule, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetProfesionalScheduleBySpecialtyInternalServerErrorCode is the HTTP code returned for type GetProfesionalScheduleBySpecialtyInternalServerError
const GetProfesionalScheduleBySpecialtyInternalServerErrorCode int = 500

/*GetProfesionalScheduleBySpecialtyInternalServerError internal server error

swagger:response getProfesionalScheduleBySpecialtyInternalServerError
*/
type GetProfesionalScheduleBySpecialtyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProfesionalScheduleBySpecialtyInternalServerError creates GetProfesionalScheduleBySpecialtyInternalServerError with default headers values
func NewGetProfesionalScheduleBySpecialtyInternalServerError() *GetProfesionalScheduleBySpecialtyInternalServerError {

	return &GetProfesionalScheduleBySpecialtyInternalServerError{}
}

// WithPayload adds the payload to the get profesional schedule by specialty internal server error response
func (o *GetProfesionalScheduleBySpecialtyInternalServerError) WithPayload(payload *models.Error) *GetProfesionalScheduleBySpecialtyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profesional schedule by specialty internal server error response
func (o *GetProfesionalScheduleBySpecialtyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfesionalScheduleBySpecialtyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
