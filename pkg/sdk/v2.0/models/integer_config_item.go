// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegerConfigItem integer config item
//
// swagger:model IntegerConfigItem
type IntegerConfigItem struct {

	// The configure item can be updated or not
	Editable bool `json:"editable"`

	// The integer value of current config item
	Value int64 `json:"value"`
}

// Validate validates this integer config item
func (m *IntegerConfigItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegerConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegerConfigItem) UnmarshalBinary(b []byte) error {
	var res IntegerConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
