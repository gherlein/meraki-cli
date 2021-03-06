// Code generated by go-swagger; DO NOT EDIT.

package m_x_warm_spare_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkWarmSpareSettingsReader is a Reader for the UpdateNetworkWarmSpareSettings structure.
type UpdateNetworkWarmSpareSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWarmSpareSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWarmSpareSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkWarmSpareSettingsOK creates a UpdateNetworkWarmSpareSettingsOK with default headers values
func NewUpdateNetworkWarmSpareSettingsOK() *UpdateNetworkWarmSpareSettingsOK {
	return &UpdateNetworkWarmSpareSettingsOK{}
}

/*UpdateNetworkWarmSpareSettingsOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkWarmSpareSettingsOK struct {
	Payload interface{}
}

func (o *UpdateNetworkWarmSpareSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/warmSpareSettings][%d] updateNetworkWarmSpareSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWarmSpareSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWarmSpareSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
