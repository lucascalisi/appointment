// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AppointmentStatus appointment status
// swagger:model appointmentStatus
type AppointmentStatus string

const (

	// AppointmentStatusConfirmed captures enum value "confirmed"
	AppointmentStatusConfirmed AppointmentStatus = "confirmed"

	// AppointmentStatusCancelled captures enum value "cancelled"
	AppointmentStatusCancelled AppointmentStatus = "cancelled"

	// AppointmentStatusPending captures enum value "pending"
	AppointmentStatusPending AppointmentStatus = "pending"

	// AppointmentStatusAvaiable captures enum value "avaiable"
	AppointmentStatusAvaiable AppointmentStatus = "avaiable"
)

// for schema
var appointmentStatusEnum []interface{}

func init() {
	var res []AppointmentStatus
	if err := json.Unmarshal([]byte(`["confirmed","cancelled","pending","avaiable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appointmentStatusEnum = append(appointmentStatusEnum, v)
	}
}

func (m AppointmentStatus) validateAppointmentStatusEnum(path, location string, value AppointmentStatus) error {
	if err := validate.Enum(path, location, value, appointmentStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this appointment status
func (m AppointmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppointmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
