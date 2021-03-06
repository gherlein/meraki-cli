// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BlinkNetworkDeviceLeds(params *BlinkNetworkDeviceLedsParams, authInfo runtime.ClientAuthInfoWriter) (*BlinkNetworkDeviceLedsOK, error)

	ClaimNetworkDevices(params *ClaimNetworkDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*ClaimNetworkDevicesOK, error)

	CycleDeviceSwitchPorts(params *CycleDeviceSwitchPortsParams, authInfo runtime.ClientAuthInfoWriter) (*CycleDeviceSwitchPortsOK, error)

	GetNetworkDevice(params *GetNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceOK, error)

	GetNetworkDeviceLldpCdp(params *GetNetworkDeviceLldpCdpParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceLldpCdpOK, error)

	GetNetworkDeviceLossAndLatencyHistory(params *GetNetworkDeviceLossAndLatencyHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceLossAndLatencyHistoryOK, error)

	GetNetworkDevicePerformance(params *GetNetworkDevicePerformanceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDevicePerformanceOK, error)

	GetNetworkDeviceUplink(params *GetNetworkDeviceUplinkParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceUplinkOK, error)

	GetNetworkDevices(params *GetNetworkDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDevicesOK, error)

	GetOrganizationDevices(params *GetOrganizationDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationDevicesOK, error)

	RebootNetworkDevice(params *RebootNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*RebootNetworkDeviceOK, error)

	RemoveNetworkDevice(params *RemoveNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveNetworkDeviceNoContent, error)

	UpdateNetworkDevice(params *UpdateNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkDeviceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BlinkNetworkDeviceLeds blinks network device leds

  Blink the LEDs on a device
*/
func (a *Client) BlinkNetworkDeviceLeds(params *BlinkNetworkDeviceLedsParams, authInfo runtime.ClientAuthInfoWriter) (*BlinkNetworkDeviceLedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBlinkNetworkDeviceLedsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "blinkNetworkDeviceLeds",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/devices/{serial}/blinkLeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BlinkNetworkDeviceLedsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BlinkNetworkDeviceLedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for blinkNetworkDeviceLeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ClaimNetworkDevices claims network devices

  Claim devices into a network
*/
func (a *Client) ClaimNetworkDevices(params *ClaimNetworkDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*ClaimNetworkDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClaimNetworkDevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "claimNetworkDevices",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/devices/claim",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ClaimNetworkDevicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClaimNetworkDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for claimNetworkDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CycleDeviceSwitchPorts cycles device switch ports

  Cycle a set of switch ports
*/
func (a *Client) CycleDeviceSwitchPorts(params *CycleDeviceSwitchPortsParams, authInfo runtime.ClientAuthInfoWriter) (*CycleDeviceSwitchPortsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCycleDeviceSwitchPortsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cycleDeviceSwitchPorts",
		Method:             "POST",
		PathPattern:        "/devices/{serial}/switch/ports/cycle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CycleDeviceSwitchPortsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CycleDeviceSwitchPortsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cycleDeviceSwitchPorts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkDevice gets network device

  Return a single device
*/
func (a *Client) GetNetworkDevice(params *GetNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDevice",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices/{serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkDeviceLldpCdp gets network device lldp cdp

  List LLDP and CDP information for a device
*/
func (a *Client) GetNetworkDeviceLldpCdp(params *GetNetworkDeviceLldpCdpParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceLldpCdpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDeviceLldpCdpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDeviceLldp_cdp",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices/{serial}/lldp_cdp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDeviceLldpCdpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDeviceLldpCdpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDeviceLldp_cdp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkDeviceLossAndLatencyHistory gets network device loss and latency history

  Get the uplink loss percentage and latency in milliseconds for a wired network device.
*/
func (a *Client) GetNetworkDeviceLossAndLatencyHistory(params *GetNetworkDeviceLossAndLatencyHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceLossAndLatencyHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDeviceLossAndLatencyHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDeviceLossAndLatencyHistory",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices/{serial}/lossAndLatencyHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDeviceLossAndLatencyHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDeviceLossAndLatencyHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDeviceLossAndLatencyHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkDevicePerformance gets network device performance

  Return the performance score for a single device. Only primary MX devices supported. If no data is available, a 204 error code is returned.
*/
func (a *Client) GetNetworkDevicePerformance(params *GetNetworkDevicePerformanceParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDevicePerformanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDevicePerformanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDevicePerformance",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices/{serial}/performance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDevicePerformanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDevicePerformanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDevicePerformance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkDeviceUplink gets network device uplink

  Return the uplink information for a device.
*/
func (a *Client) GetNetworkDeviceUplink(params *GetNetworkDeviceUplinkParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDeviceUplinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDeviceUplinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDeviceUplink",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices/{serial}/uplink",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDeviceUplinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDeviceUplinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDeviceUplink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkDevices gets network devices

  List the devices in a network
*/
func (a *Client) GetNetworkDevices(params *GetNetworkDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkDevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkDevices",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkDevicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationDevices gets organization devices

  List the devices in an organization
*/
func (a *Client) GetOrganizationDevices(params *GetOrganizationDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationDevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationDevices",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationDevicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RebootNetworkDevice reboots network device

  Reboot a device
*/
func (a *Client) RebootNetworkDevice(params *RebootNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*RebootNetworkDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootNetworkDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rebootNetworkDevice",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/devices/{serial}/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RebootNetworkDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RebootNetworkDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rebootNetworkDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveNetworkDevice removes network device

  Remove a single device
*/
func (a *Client) RemoveNetworkDevice(params *RemoveNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveNetworkDeviceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveNetworkDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeNetworkDevice",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/devices/{serial}/remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveNetworkDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveNetworkDeviceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeNetworkDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkDevice updates network device

  Update the attributes of a device
*/
func (a *Client) UpdateNetworkDevice(params *UpdateNetworkDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkDevice",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/devices/{serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
