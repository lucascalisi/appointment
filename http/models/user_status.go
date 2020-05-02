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

// UserStatus user status
// swagger:model userStatus
type UserStatus string

const (

	// UserStatusEnabled captures enum value "enabled"
	UserStatusEnabled UserStatus = "enabled"

	// UserStatusDisabled captures enum value "disabled"
	UserStatusDisabled UserStatus = "disabled"

	// UserStatusLocked captures enum value "locked"
	UserStatusLocked UserStatus = "locked"

	// UserStatusPending captures enum value "pending"
	UserStatusPending UserStatus = "pending"
)

// for schema
var userStatusEnum []interface{}

func init() {
	var res []UserStatus
	if err := json.Unmarshal([]byte(`["enabled","disabled","locked","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userStatusEnum = append(userStatusEnum, v)
	}
}

func (m UserStatus) validateUserStatusEnum(path, location string, value UserStatus) error {
	if err := validate.Enum(path, location, value, userStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this user status
func (m UserStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
