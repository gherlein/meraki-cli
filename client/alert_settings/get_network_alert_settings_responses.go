// Code generated by go-swagger; DO NOT EDIT.

package alert_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkAlertSettingsReader is a Reader for the GetNetworkAlertSettings structure.
type GetNetworkAlertSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAlertSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkAlertSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkAlertSettingsOK creates a GetNetworkAlertSettingsOK with default headers values
func NewGetNetworkAlertSettingsOK() *GetNetworkAlertSettingsOK {
	return &GetNetworkAlertSettingsOK{}
}

/*GetNetworkAlertSettingsOK handles this case with default header values.

Successful operation
*/
type GetNetworkAlertSettingsOK struct {
	Payload interface{}
}

func (o *GetNetworkAlertSettingsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/alertSettings][%d] getNetworkAlertSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAlertSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkAlertSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}