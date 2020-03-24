// Code generated by go-swagger; DO NOT EDIT.

package bluetooth_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkBluetoothSettingsReader is a Reader for the GetNetworkBluetoothSettings structure.
type GetNetworkBluetoothSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkBluetoothSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkBluetoothSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkBluetoothSettingsOK creates a GetNetworkBluetoothSettingsOK with default headers values
func NewGetNetworkBluetoothSettingsOK() *GetNetworkBluetoothSettingsOK {
	return &GetNetworkBluetoothSettingsOK{}
}

/*GetNetworkBluetoothSettingsOK handles this case with default header values.

Successful operation
*/
type GetNetworkBluetoothSettingsOK struct {
	Payload interface{}
}

func (o *GetNetworkBluetoothSettingsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/bluetoothSettings][%d] getNetworkBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkBluetoothSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkBluetoothSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}