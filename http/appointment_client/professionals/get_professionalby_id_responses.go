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

// GetProfessionalbyIDReader is a Reader for the GetProfessionalbyID structure.
type GetProfessionalbyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfessionalbyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfessionalbyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetProfessionalbyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProfessionalbyIDOK creates a GetProfessionalbyIDOK with default headers values
func NewGetProfessionalbyIDOK() *GetProfessionalbyIDOK {
	return &GetProfessionalbyIDOK{}
}

/*GetProfessionalbyIDOK handles this case with default header values.

result professional
*/
type GetProfessionalbyIDOK struct {
	Payload *models.Professional
}

func (o *GetProfessionalbyIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/professionals/{id}][%d] getProfessionalbyIdOK  %+v", 200, o.Payload)
}

func (o *GetProfessionalbyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Professional)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfessionalbyIDInternalServerError creates a GetProfessionalbyIDInternalServerError with default headers values
func NewGetProfessionalbyIDInternalServerError() *GetProfessionalbyIDInternalServerError {
	return &GetProfessionalbyIDInternalServerError{}
}

/*GetProfessionalbyIDInternalServerError handles this case with default header values.

internal server error
*/
type GetProfessionalbyIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetProfessionalbyIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/professionals/{id}][%d] getProfessionalbyIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProfessionalbyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
