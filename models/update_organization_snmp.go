// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrganizationSnmp updateOrganizationSnmp
//
// swagger:model updateOrganizationSnmp
type UpdateOrganizationSnmp struct {

	// The IPs that are allowed to access the SNMP server. This list should be IPv4 addresses separated by semi-colons (ie. "1.2.3.4;2.3.4.5").
	PeerIps string `json:"peerIps,omitempty"`

	// Boolean indicating whether SNMP version 2c is enabled for the organization.
	V2cEnabled bool `json:"v2cEnabled,omitempty"`

	// v3 auth mode
	V3AuthMode V3AuthMode `json:"v3AuthMode,omitempty"`

	// The SNMP version 3 authentication password. Must be at least 8 characters if specified.
	V3AuthPass string `json:"v3AuthPass,omitempty"`

	// Boolean indicating whether SNMP version 3 is enabled for the organization.
	V3Enabled bool `json:"v3Enabled,omitempty"`

	// v3 priv mode
	V3PrivMode V3PrivMode `json:"v3PrivMode,omitempty"`

	// The SNMP version 3 privacy password. Must be at least 8 characters if specified.
	V3PrivPass string `json:"v3PrivPass,omitempty"`
}

// Validate validates this update organization snmp
func (m *UpdateOrganizationSnmp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateV3AuthMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateV3PrivMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOrganizationSnmp) validateV3AuthMode(formats strfmt.Registry) error {

	if swag.IsZero(m.V3AuthMode) { // not required
		return nil
	}

	if err := m.V3AuthMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("v3AuthMode")
		}
		return err
	}

	return nil
}

func (m *UpdateOrganizationSnmp) validateV3PrivMode(formats strfmt.Registry) error {

	if swag.IsZero(m.V3PrivMode) { // not required
		return nil
	}

	if err := m.V3PrivMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("v3PrivMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrganizationSnmp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrganizationSnmp) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationSnmp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
