// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSmAppPolarisReader is a Reader for the DeleteNetworkSmAppPolaris structure.
type DeleteNetworkSmAppPolarisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSmAppPolarisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkSmAppPolarisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNetworkSmAppPolarisOK creates a DeleteNetworkSmAppPolarisOK with default headers values
func NewDeleteNetworkSmAppPolarisOK() *DeleteNetworkSmAppPolarisOK {
	return &DeleteNetworkSmAppPolarisOK{}
}

/*DeleteNetworkSmAppPolarisOK handles this case with default header values.

Successful operation
*/
type DeleteNetworkSmAppPolarisOK struct {
	Payload interface{}
}

func (o *DeleteNetworkSmAppPolarisOK) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/sm/app/polaris/{appId}][%d] deleteNetworkSmAppPolarisOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkSmAppPolarisOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteNetworkSmAppPolarisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}