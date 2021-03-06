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

// Quality1 Quality1
//
// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
//
// swagger:model Quality1
type Quality1 string

const (

	// Quality1Standard captures enum value "Standard"
	Quality1Standard Quality1 = "Standard"

	// Quality1Enhanced captures enum value "Enhanced"
	Quality1Enhanced Quality1 = "Enhanced"

	// Quality1High captures enum value "High"
	Quality1High Quality1 = "High"
)

// for schema
var quality1Enum []interface{}

func init() {
	var res []Quality1
	if err := json.Unmarshal([]byte(`["Standard","Enhanced","High"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		quality1Enum = append(quality1Enum, v)
	}
}

func (m Quality1) validateQuality1Enum(path, location string, value Quality1) error {
	if err := validate.Enum(path, location, value, quality1Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this quality1
func (m Quality1) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQuality1Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
