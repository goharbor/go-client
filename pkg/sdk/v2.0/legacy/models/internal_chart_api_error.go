// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InternalChartAPIError Internal server error occurred
//
// swagger:model InternalChartAPIError
type InternalChartAPIError struct {
	ChartAPIError
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InternalChartAPIError) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ChartAPIError
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ChartAPIError = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InternalChartAPIError) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ChartAPIError)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this internal chart API error
func (m *InternalChartAPIError) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChartAPIError
	if err := m.ChartAPIError.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this internal chart API error based on the context it is used
func (m *InternalChartAPIError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChartAPIError
	if err := m.ChartAPIError.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InternalChartAPIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalChartAPIError) UnmarshalBinary(b []byte) error {
	var res InternalChartAPIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
