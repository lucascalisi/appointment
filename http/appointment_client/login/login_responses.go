// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/appointment/http/models"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLoginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*LoginOK handles this case with default header values.

login successful
*/
type LoginOK struct {
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] loginOK ", 200)
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginUnauthorized creates a LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {
	return &LoginUnauthorized{}
}

/*LoginUnauthorized handles this case with default header values.

access denied
*/
type LoginUnauthorized struct {
}

func (o *LoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] loginUnauthorized ", 401)
}

func (o *LoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginInternalServerError creates a LoginInternalServerError with default headers values
func NewLoginInternalServerError() *LoginInternalServerError {
	return &LoginInternalServerError{}
}

/*LoginInternalServerError handles this case with default header values.

internal server error
*/
type LoginInternalServerError struct {
	Payload *models.Error
}

func (o *LoginInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] loginInternalServerError  %+v", 500, o.Payload)
}

func (o *LoginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LoginBody login body
swagger:model LoginBody
*/
type LoginBody struct {

	// password
	// Required: true
	// Format: password
	Password *strfmt.Password `json:"password"`

	// user
	// Required: true
	User *string `json:"user"`
}

// Validate validates this login body
func (o *LoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("credentials"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("credentials"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *LoginBody) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("credentials"+"."+"user", "body", o.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginBody) UnmarshalBinary(b []byte) error {
	var res LoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}