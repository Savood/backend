// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/globalsign/mgo/bson"
)

// Chat chat
// swagger:model Chat
type Chat struct {

	// id
	ID bson.ObjectId `json:"_id,omitempty"`

	// offering Id
	OfferingID []bson.ObjectId `json:"offeringId"`

	// partner
	Partner *UserShort `json:"partner,omitempty"`
}

// Validate validates this chat
func (m *Chat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Chat) validatePartner(formats strfmt.Registry) error {

	if swag.IsZero(m.Partner) { // not required
		return nil
	}

	if m.Partner != nil {
		if err := m.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Chat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Chat) UnmarshalBinary(b []byte) error {
	var res Chat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
