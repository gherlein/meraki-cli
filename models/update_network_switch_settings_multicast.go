// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkSwitchSettingsMulticast updateNetworkSwitchSettingsMulticast
//
// swagger:model updateNetworkSwitchSettingsMulticast
type UpdateNetworkSwitchSettingsMulticast struct {

	// default settings
	DefaultSettings *DefaultSettings `json:"defaultSettings,omitempty"`

	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []*Override1 `json:"overrides"`
}

// Validate validates this update network switch settings multicast
func (m *UpdateNetworkSwitchSettingsMulticast) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSwitchSettingsMulticast) validateDefaultSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultSettings) { // not required
		return nil
	}

	if m.DefaultSettings != nil {
		if err := m.DefaultSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultSettings")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateNetworkSwitchSettingsMulticast) validateOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(m.Overrides); i++ {
		if swag.IsZero(m.Overrides[i]) { // not required
			continue
		}

		if m.Overrides[i] != nil {
			if err := m.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsMulticast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsMulticast) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchSettingsMulticast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}