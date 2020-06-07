// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/appointment/http/models"
)

// CancelAppoinmentForPatientOKCode is the HTTP code returned for type CancelAppoinmentForPatientOK
const CancelAppoinmentForPatientOKCode int = 200

/*CancelAppoinmentForPatientOK appoint for a patient

swagger:response cancelAppoinmentForPatientOK
*/
type CancelAppoinmentForPatientOK struct {
}

// NewCancelAppoinmentForPatientOK creates CancelAppoinmentForPatientOK with default headers values
func NewCancelAppoinmentForPatientOK() *CancelAppoinmentForPatientOK {

	return &CancelAppoinmentForPatientOK{}
}

// WriteResponse to the client
func (o *CancelAppoinmentForPatientOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CancelAppoinmentForPatientInternalServerErrorCode is the HTTP code returned for type CancelAppoinmentForPatientInternalServerError
const CancelAppoinmentForPatientInternalServerErrorCode int = 500

/*CancelAppoinmentForPatientInternalServerError internal server error

swagger:response cancelAppoinmentForPatientInternalServerError
*/
type CancelAppoinmentForPatientInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCancelAppoinmentForPatientInternalServerError creates CancelAppoinmentForPatientInternalServerError with default headers values
func NewCancelAppoinmentForPatientInternalServerError() *CancelAppoinmentForPatientInternalServerError {

	return &CancelAppoinmentForPatientInternalServerError{}
}

// WithPayload adds the payload to the cancel appoinment for patient internal server error response
func (o *CancelAppoinmentForPatientInternalServerError) WithPayload(payload *models.Error) *CancelAppoinmentForPatientInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cancel appoinment for patient internal server error response
func (o *CancelAppoinmentForPatientInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CancelAppoinmentForPatientInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
