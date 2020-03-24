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

// Protocol4 Protocol4
//
// Either of the following: 'tcp', 'udp', 'icmp-ping' or 'any'
//
// swagger:model Protocol4
type Protocol4 string

const (

	// Protocol4TCP captures enum value "tcp"
	Protocol4TCP Protocol4 = "tcp"

	// Protocol4UDP captures enum value "udp"
	Protocol4UDP Protocol4 = "udp"

	// Protocol4IcmpPing captures enum value "icmp-ping"
	Protocol4IcmpPing Protocol4 = "icmp-ping"

	// Protocol4Any captures enum value "any"
	Protocol4Any Protocol4 = "any"
)

// for schema
var protocol4Enum []interface{}

func init() {
	var res []Protocol4
	if err := json.Unmarshal([]byte(`["tcp","udp","icmp-ping","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protocol4Enum = append(protocol4Enum, v)
	}
}

func (m Protocol4) validateProtocol4Enum(path, location string, value Protocol4) error {
	if err := validate.Enum(path, location, value, protocol4Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this protocol4
func (m Protocol4) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProtocol4Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}