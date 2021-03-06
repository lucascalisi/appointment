// Code generated by go-swagger; DO NOT EDIT.

package professionals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/appointment/http/models"
)

// GetAppointmentByProfessionalAppointmentIDReader is a Reader for the GetAppointmentByProfessionalAppointmentID structure.
type GetAppointmentByProfessionalAppointmentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppointmentByProfessionalAppointmentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppointmentByProfessionalAppointmentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetAppointmentByProfessionalAppointmentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppointmentByProfessionalAppointmentIDOK creates a GetAppointmentByProfessionalAppointmentIDOK with default headers values
func NewGetAppointmentByProfessionalAppointmentIDOK() *GetAppointmentByProfessionalAppointmentIDOK {
	return &GetAppointmentByProfessionalAppointmentIDOK{}
}

/*GetAppointmentByProfessionalAppointmentIDOK handles this case with default header values.

get appointment for professional
*/
type GetAppointmentByProfessionalAppointmentIDOK struct {
	Payload *models.Appointment
}

func (o *GetAppointmentByProfessionalAppointmentIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/professionals/{id}/appointments/{idappointment}][%d] getAppointmentByProfessionalAppointmentIdOK  %+v", 200, o.Payload)
}

func (o *GetAppointmentByProfessionalAppointmentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Appointment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppointmentByProfessionalAppointmentIDInternalServerError creates a GetAppointmentByProfessionalAppointmentIDInternalServerError with default headers values
func NewGetAppointmentByProfessionalAppointmentIDInternalServerError() *GetAppointmentByProfessionalAppointmentIDInternalServerError {
	return &GetAppointmentByProfessionalAppointmentIDInternalServerError{}
}

/*GetAppointmentByProfessionalAppointmentIDInternalServerError handles this case with default header values.

internal server error
*/
type GetAppointmentByProfessionalAppointmentIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetAppointmentByProfessionalAppointmentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/professionals/{id}/appointments/{idappointment}][%d] getAppointmentByProfessionalAppointmentIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppointmentByProfessionalAppointmentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
