// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/appointment/http/models"
)

// GetPatientbyIDReader is a Reader for the GetPatientbyID structure.
type GetPatientbyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatientbyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPatientbyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetPatientbyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatientbyIDOK creates a GetPatientbyIDOK with default headers values
func NewGetPatientbyIDOK() *GetPatientbyIDOK {
	return &GetPatientbyIDOK{}
}

/*GetPatientbyIDOK handles this case with default header values.

result get patient
*/
type GetPatientbyIDOK struct {
	Payload *models.Patient
}

func (o *GetPatientbyIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/patients/{id}][%d] getPatientbyIdOK  %+v", 200, o.Payload)
}

func (o *GetPatientbyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Patient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatientbyIDInternalServerError creates a GetPatientbyIDInternalServerError with default headers values
func NewGetPatientbyIDInternalServerError() *GetPatientbyIDInternalServerError {
	return &GetPatientbyIDInternalServerError{}
}

/*GetPatientbyIDInternalServerError handles this case with default header values.

internal server error
*/
type GetPatientbyIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetPatientbyIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/patients/{id}][%d] getPatientbyIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPatientbyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
