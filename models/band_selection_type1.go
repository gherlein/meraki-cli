// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BandSelectionType1 BandSelectionType1
//
// Band selection can be set to either 'ssid' or 'ap'.
//
// swagger:model BandSelectionType1
type BandSelectionType1 string

const (

	// BandSelectionType1Ssid captures enum value "ssid"
	BandSelectionType1Ssid BandSelectionType1 = "ssid"

	// BandSelectionType1Ap captures enum value "ap"
	BandSelectionType1Ap BandSelectionType1 = "ap"
)

// for schema
var bandSelectionType1Enum []interface{}

func init() {
	var res []BandSelectionType1
	if err := json.Unmarshal([]byte(`["ssid","ap"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bandSelectionType1Enum = append(bandSelectionType1Enum, v)
	}
}

func (m BandSelectionType1) validateBandSelectionType1Enum(path, location string, value BandSelectionType1) error {
	if err := validate.Enum(path, location, value, bandSelectionType1Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this band selection type1
func (m BandSelectionType1) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBandSelectionType1Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
