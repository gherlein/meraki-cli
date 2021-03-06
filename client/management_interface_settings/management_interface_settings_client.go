// Code generated by go-swagger; DO NOT EDIT.

package management_interface_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new management interface settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management interface settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkDeviceManagementInterfaceSettings(params *GetNetworkDeviceManagementInterfaceSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceManagementInterfaceSettingsOK, error)

	UpdateNetworkDeviceManagementInterfaceSettings(params *UpdateNetworkDeviceManagementInterfaceSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkDeviceManagementInterfaceSettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkDeviceManagementInterfaceSettings gets network device management interface settings

  Return the management interface settings for a device
*/
func (a *Client) GetNetworkDeviceManagementInterfaceSettings(params *GetNetworkDeviceManagementInterfaceSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceManagementInterfaceSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDeviceManagementInterfaceSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDeviceManagementInterfaceSettings",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices/{serial}/managementInterfaceSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDeviceManagementInterfaceSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDeviceManagementInterfaceSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDeviceManagementInterfaceSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkDeviceManagementInterfaceSettings updates network device management interface settings

  Update the management interface settings for a device
*/
func (a *Client) UpdateNetworkDeviceManagementInterfaceSettings(params *UpdateNetworkDeviceManagementInterfaceSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkDeviceManagementInterfaceSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkDeviceManagementInterfaceSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkDeviceManagementInterfaceSettings",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/devices/{serial}/managementInterfaceSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkDeviceManagementInterfaceSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkDeviceManagementInterfaceSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkDeviceManagementInterfaceSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
