// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Email Email represents an email address
// swagger:model Email
type Email struct {

	// Primary: True if the field is the primary field; false if the field
	// is a secondary
	// field.
	Primary bool `json:"primary,omitempty"`

	// Type: The type of the address. The type can be custom or one of these
	// predefined values:
	//
	// `home`
	// `work`
	// `other`
	Type string `json:"type,omitempty"`

	// Value: The email address.
	Value string `json:"value,omitempty"`

	// Verified: True if the field is verified; false if the field is
	// unverified. A
	// verified field is typically a name, email address, phone number,
	// or
	// website that has been confirmed to be owned by the person.
	Verified bool `json:"verified,omitempty"`

	// metadata
	Metadata RawMessage `json:"metadata,omitempty"`
}

// Validate validates this email
func (m *Email) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Email) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := m.Metadata.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metadata")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Email) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Email) UnmarshalBinary(b []byte) error {
	var res Email
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}