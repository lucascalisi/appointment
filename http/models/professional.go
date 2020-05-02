// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Professional professional
// swagger:model professional
type Professional struct {

	// birth day
	// Required: true
	// Format: date
	BirthDay *strfmt.Date `json:"birthDay"`

	// dni
	// Required: true
	Dni *int64 `json:"dni"`

	// doctor number
	// Required: true
	DoctorNumber *int64 `json:"doctorNumber"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this professional
func (m *Professional) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDni(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoctorNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Professional) validateBirthDay(formats strfmt.Registry) error {

	if err := validate.Required("birthDay", "body", m.BirthDay); err != nil {
		return err
	}

	if err := validate.FormatOf("birthDay", "body", "date", m.BirthDay.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Professional) validateDni(formats strfmt.Registry) error {

	if err := validate.Required("dni", "body", m.Dni); err != nil {
		return err
	}

	return nil
}

func (m *Professional) validateDoctorNumber(formats strfmt.Registry) error {

	if err := validate.Required("doctorNumber", "body", m.DoctorNumber); err != nil {
		return err
	}

	return nil
}

func (m *Professional) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Professional) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Professional) UnmarshalBinary(b []byte) error {
	var res Professional
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
