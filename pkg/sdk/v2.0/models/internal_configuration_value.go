// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InternalConfigurationValue internal configuration value
//
// swagger:model InternalConfigurationValue
type InternalConfigurationValue struct {

	// The configure item can be updated or not
	Editable bool `json:"editable"`

	// The value of current config item
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this internal configuration value
func (m *InternalConfigurationValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InternalConfigurationValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalConfigurationValue) UnmarshalBinary(b []byte) error {
	var res InternalConfigurationValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
