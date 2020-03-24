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

// BandSelectionType BandSelectionType
//
// Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
//
// swagger:model BandSelectionType
type BandSelectionType string

const (

	// BandSelectionTypeSsid captures enum value "ssid"
	BandSelectionTypeSsid BandSelectionType = "ssid"

	// BandSelectionTypeAp captures enum value "ap"
	BandSelectionTypeAp BandSelectionType = "ap"
)

// for schema
var bandSelectionTypeEnum []interface{}

func init() {
	var res []BandSelectionType
	if err := json.Unmarshal([]byte(`["ssid","ap"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bandSelectionTypeEnum = append(bandSelectionTypeEnum, v)
	}
}

func (m BandSelectionType) validateBandSelectionTypeEnum(path, location string, value BandSelectionType) error {
	if err := validate.Enum(path, location, value, bandSelectionTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this band selection type
func (m BandSelectionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBandSelectionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}