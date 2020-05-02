// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/appointment/http/models"
)

// GetRoleOfUserOKCode is the HTTP code returned for type GetRoleOfUserOK
const GetRoleOfUserOKCode int = 200

/*GetRoleOfUserOK get user role

swagger:response getRoleOfUserOK
*/
type GetRoleOfUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Role `json:"body,omitempty"`
}

// NewGetRoleOfUserOK creates GetRoleOfUserOK with default headers values
func NewGetRoleOfUserOK() *GetRoleOfUserOK {

	return &GetRoleOfUserOK{}
}

// WithPayload adds the payload to the get role of user o k response
func (o *GetRoleOfUserOK) WithPayload(payload *models.Role) *GetRoleOfUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get role of user o k response
func (o *GetRoleOfUserOK) SetPayload(payload *models.Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoleOfUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRoleOfUserInternalServerErrorCode is the HTTP code returned for type GetRoleOfUserInternalServerError
const GetRoleOfUserInternalServerErrorCode int = 500

/*GetRoleOfUserInternalServerError internal server error

swagger:response getRoleOfUserInternalServerError
*/
type GetRoleOfUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRoleOfUserInternalServerError creates GetRoleOfUserInternalServerError with default headers values
func NewGetRoleOfUserInternalServerError() *GetRoleOfUserInternalServerError {

	return &GetRoleOfUserInternalServerError{}
}

// WithPayload adds the payload to the get role of user internal server error response
func (o *GetRoleOfUserInternalServerError) WithPayload(payload *models.Error) *GetRoleOfUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get role of user internal server error response
func (o *GetRoleOfUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoleOfUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
