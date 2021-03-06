// Code generated by go-swagger; DO NOT EDIT.

package radio_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkDeviceWirelessRadioSettingsReader is a Reader for the GetNetworkDeviceWirelessRadioSettings structure.
type GetNetworkDeviceWirelessRadioSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDeviceWirelessRadioSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkDeviceWirelessRadioSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkDeviceWirelessRadioSettingsOK creates a GetNetworkDeviceWirelessRadioSettingsOK with default headers values
func NewGetNetworkDeviceWirelessRadioSettingsOK() *GetNetworkDeviceWirelessRadioSettingsOK {
	return &GetNetworkDeviceWirelessRadioSettingsOK{}
}

/*GetNetworkDeviceWirelessRadioSettingsOK handles this case with default header values.

Successful operation
*/
type GetNetworkDeviceWirelessRadioSettingsOK struct {
	Payload interface{}
}

func (o *GetNetworkDeviceWirelessRadioSettingsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/devices/{serial}/wireless/radioSettings][%d] getNetworkDeviceWirelessRadioSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkDeviceWirelessRadioSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkDeviceWirelessRadioSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
