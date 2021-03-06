// Code generated by go-swagger; DO NOT EDIT.

package intrusion_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new intrusion settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for intrusion settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkSecurityIntrusionSettings(params *GetNetworkSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSecurityIntrusionSettingsOK, error)

	GetOrganizationSecurityIntrusionSettings(params *GetOrganizationSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationSecurityIntrusionSettingsOK, error)

	UpdateNetworkSecurityIntrusionSettings(params *UpdateNetworkSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSecurityIntrusionSettingsOK, error)

	UpdateOrganizationSecurityIntrusionSettings(params *UpdateOrganizationSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationSecurityIntrusionSettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkSecurityIntrusionSettings gets network security intrusion settings

  Returns all supported intrusion settings for an MX network
*/
func (a *Client) GetNetworkSecurityIntrusionSettings(params *GetNetworkSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSecurityIntrusionSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkSecurityIntrusionSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkSecurityIntrusionSettings",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/security/intrusionSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkSecurityIntrusionSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkSecurityIntrusionSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkSecurityIntrusionSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationSecurityIntrusionSettings gets organization security intrusion settings

  Returns all supported intrusion settings for an organization
*/
func (a *Client) GetOrganizationSecurityIntrusionSettings(params *GetOrganizationSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationSecurityIntrusionSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationSecurityIntrusionSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationSecurityIntrusionSettings",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/security/intrusionSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationSecurityIntrusionSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationSecurityIntrusionSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationSecurityIntrusionSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkSecurityIntrusionSettings updates network security intrusion settings

  Set the supported intrusion settings for an MX network
*/
func (a *Client) UpdateNetworkSecurityIntrusionSettings(params *UpdateNetworkSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSecurityIntrusionSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkSecurityIntrusionSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkSecurityIntrusionSettings",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/security/intrusionSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkSecurityIntrusionSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkSecurityIntrusionSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkSecurityIntrusionSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOrganizationSecurityIntrusionSettings updates organization security intrusion settings

  Sets supported intrusion settings for an organization
*/
func (a *Client) UpdateOrganizationSecurityIntrusionSettings(params *UpdateOrganizationSecurityIntrusionSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationSecurityIntrusionSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationSecurityIntrusionSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationSecurityIntrusionSettings",
		Method:             "PUT",
		PathPattern:        "/organizations/{organizationId}/security/intrusionSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationSecurityIntrusionSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationSecurityIntrusionSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrganizationSecurityIntrusionSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
