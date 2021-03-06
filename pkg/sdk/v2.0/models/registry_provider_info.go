// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistryProviderInfo The registry provider info contains the base info and capability declarations of the registry provider
//
// swagger:model RegistryProviderInfo
type RegistryProviderInfo struct {

	// The credential pattern
	CredentialPattern *RegistryProviderCredentialPattern `json:"credential_pattern,omitempty"`

	// The endpoint pattern
	EndpointPattern *RegistryProviderEndpointPattern `json:"endpoint_pattern,omitempty"`
}

// Validate validates this registry provider info
func (m *RegistryProviderInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointPattern(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryProviderInfo) validateCredentialPattern(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialPattern) { // not required
		return nil
	}

	if m.CredentialPattern != nil {
		if err := m.CredentialPattern.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential_pattern")
			}
			return err
		}
	}

	return nil
}

func (m *RegistryProviderInfo) validateEndpointPattern(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointPattern) { // not required
		return nil
	}

	if m.EndpointPattern != nil {
		if err := m.EndpointPattern.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint_pattern")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryProviderInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryProviderInfo) UnmarshalBinary(b []byte) error {
	var res RegistryProviderInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
