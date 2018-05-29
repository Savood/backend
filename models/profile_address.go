// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProfileAddress profile address
// swagger:model profileAddress
type ProfileAddress struct {

	// city
	City string `json:"city,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// street
	Street string `json:"street,omitempty"`

	// zip
	Zip int64 `json:"zip,omitempty"`
}

// Validate validates this profile address
func (m *ProfileAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProfileAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProfileAddress) UnmarshalBinary(b []byte) error {
	var res ProfileAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
