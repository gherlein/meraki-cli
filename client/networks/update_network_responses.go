// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkReader is a Reader for the UpdateNetwork structure.
type UpdateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkOK creates a UpdateNetworkOK with default headers values
func NewUpdateNetworkOK() *UpdateNetworkOK {
	return &UpdateNetworkOK{}
}

/*UpdateNetworkOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkOK struct {
	Payload interface{}
}

func (o *UpdateNetworkOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}][%d] updateNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}