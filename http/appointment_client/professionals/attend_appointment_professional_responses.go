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

// AttendAppointmentProfessionalReader is a Reader for the AttendAppointmentProfessional structure.
type AttendAppointmentProfessionalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttendAppointmentProfessionalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAttendAppointmentProfessionalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewAttendAppointmentProfessionalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAttendAppointmentProfessionalOK creates a AttendAppointmentProfessionalOK with default headers values
func NewAttendAppointmentProfessionalOK() *AttendAppointmentProfessionalOK {
	return &AttendAppointmentProfessionalOK{}
}

/*AttendAppointmentProfessionalOK handles this case with default header values.

set attend appointment for professional
*/
type AttendAppointmentProfessionalOK struct {
}

func (o *AttendAppointmentProfessionalOK) Error() string {
	return fmt.Sprintf("[PUT /v1/professionals/{id}/appointments/{idappointment}/attend][%d] attendAppointmentProfessionalOK ", 200)
}

func (o *AttendAppointmentProfessionalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAttendAppointmentProfessionalInternalServerError creates a AttendAppointmentProfessionalInternalServerError with default headers values
func NewAttendAppointmentProfessionalInternalServerError() *AttendAppointmentProfessionalInternalServerError {
	return &AttendAppointmentProfessionalInternalServerError{}
}

/*AttendAppointmentProfessionalInternalServerError handles this case with default header values.

internal server error
*/
type AttendAppointmentProfessionalInternalServerError struct {
	Payload *models.Error
}

func (o *AttendAppointmentProfessionalInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/professionals/{id}/appointments/{idappointment}/attend][%d] attendAppointmentProfessionalInternalServerError  %+v", 500, o.Payload)
}

func (o *AttendAppointmentProfessionalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
