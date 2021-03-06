// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rule7 Rule7
//
// swagger:model Rule7
type Rule7 struct {

	// An array of associated forwarding rules
	// Required: true
	PortRules []*PortRule `json:"portRules"`

	// The IP address that will be used to access the internal resource from the WAN
	// Required: true
	PublicIP *string `json:"publicIp"`

	// uplink
	// Required: true
	Uplink Uplink1 `json:"uplink"`
}

// Validate validates this rule7
func (m *Rule7) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePortRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rule7) validatePortRules(formats strfmt.Registry) error {

	if err := validate.Required("portRules", "body", m.PortRules); err != nil {
		return err
	}

	for i := 0; i < len(m.PortRules); i++ {
		if swag.IsZero(m.PortRules[i]) { // not required
			continue
		}

		if m.PortRules[i] != nil {
			if err := m.PortRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Rule7) validatePublicIP(formats strfmt.Registry) error {

	if err := validate.Required("publicIp", "body", m.PublicIP); err != nil {
		return err
	}

	return nil
}

func (m *Rule7) validateUplink(formats strfmt.Registry) error {

	if err := m.Uplink.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("uplink")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule7) UnmarshalBinary(b []byte) error {
	var res Rule7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
