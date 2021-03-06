// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkClientPolicyReader is a Reader for the UpdateNetworkClientPolicy structure.
type UpdateNetworkClientPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkClientPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkClientPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkClientPolicyOK creates a UpdateNetworkClientPolicyOK with default headers values
func NewUpdateNetworkClientPolicyOK() *UpdateNetworkClientPolicyOK {
	return &UpdateNetworkClientPolicyOK{}
}

/*UpdateNetworkClientPolicyOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkClientPolicyOK struct {
	Payload interface{}
}

func (o *UpdateNetworkClientPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/clients/{clientId}/policy][%d] updateNetworkClientPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkClientPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkClientPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
