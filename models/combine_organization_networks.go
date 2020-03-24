// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CombineOrganizationNetworks combineOrganizationNetworks
//
// swagger:model combineOrganizationNetworks
type CombineOrganizationNetworks struct {

	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	EnrollmentString string `json:"enrollmentString,omitempty"`

	// The name of the combined network
	// Required: true
	Name *string `json:"name"`

	// A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
	// Required: true
	NetworkIds []string `json:"networkIds"`
}

// Validate validates this combine organization networks
func (m *CombineOrganizationNetworks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CombineOrganizationNetworks) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CombineOrganizationNetworks) validateNetworkIds(formats strfmt.Registry) error {

	if err := validate.Required("networkIds", "body", m.NetworkIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CombineOrganizationNetworks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CombineOrganizationNetworks) UnmarshalBinary(b []byte) error {
	var res CombineOrganizationNetworks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}