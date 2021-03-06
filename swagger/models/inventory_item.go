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

// InventoryItem inventory item
// swagger:model InventoryItem
type InventoryItem struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// manufacturer
	Manufacturer *Manufacturer `json:"manufacturer,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// release date
	// Format: date-time
	ReleaseDate strfmt.DateTime `json:"releaseDate,omitempty"`
}

// Validate validates this inventory item
func (m *InventoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InventoryItem) validateManufacturer(formats strfmt.Registry) error {

	if swag.IsZero(m.Manufacturer) { // not required
		return nil
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryItem) validateReleaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseDate", "body", "date-time", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryItem) UnmarshalBinary(b []byte) error {
	var res InventoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
