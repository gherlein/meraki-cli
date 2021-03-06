// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkDeviceReader is a Reader for the GetNetworkDevice structure.
type GetNetworkDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkDeviceOK creates a GetNetworkDeviceOK with default headers values
func NewGetNetworkDeviceOK() *GetNetworkDeviceOK {
	return &GetNetworkDeviceOK{}
}

/*GetNetworkDeviceOK handles this case with default header values.

Successful operation
*/
type GetNetworkDeviceOK struct {
	Payload interface{}
}

func (o *GetNetworkDeviceOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/devices/{serial}][%d] getNetworkDeviceOK  %+v", 200, o.Payload)
}

func (o *GetNetworkDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
