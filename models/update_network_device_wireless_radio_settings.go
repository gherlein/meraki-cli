// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkDeviceWirelessRadioSettings updateNetworkDeviceWirelessRadioSettings
//
// swagger:model updateNetworkDeviceWirelessRadioSettings
type UpdateNetworkDeviceWirelessRadioSettings struct {

	// The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile
	//     (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides
	//     on the device (channel width, channel, power).
	RfProfileID int32 `json:"rfProfileId,omitempty"`
}

// Validate validates this update network device wireless radio settings
func (m *UpdateNetworkDeviceWirelessRadioSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkDeviceWirelessRadioSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkDeviceWirelessRadioSettings) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkDeviceWirelessRadioSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
