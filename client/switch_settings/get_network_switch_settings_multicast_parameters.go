// Code generated by go-swagger; DO NOT EDIT.

package switch_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetworkSwitchSettingsMulticastParams creates a new GetNetworkSwitchSettingsMulticastParams object
// with the default values initialized.
func NewGetNetworkSwitchSettingsMulticastParams() *GetNetworkSwitchSettingsMulticastParams {
	var ()
	return &GetNetworkSwitchSettingsMulticastParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchSettingsMulticastParamsWithTimeout creates a new GetNetworkSwitchSettingsMulticastParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSwitchSettingsMulticastParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchSettingsMulticastParams {
	var ()
	return &GetNetworkSwitchSettingsMulticastParams{

		timeout: timeout,
	}
}

// NewGetNetworkSwitchSettingsMulticastParamsWithContext creates a new GetNetworkSwitchSettingsMulticastParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSwitchSettingsMulticastParamsWithContext(ctx context.Context) *GetNetworkSwitchSettingsMulticastParams {
	var ()
	return &GetNetworkSwitchSettingsMulticastParams{

		Context: ctx,
	}
}

// NewGetNetworkSwitchSettingsMulticastParamsWithHTTPClient creates a new GetNetworkSwitchSettingsMulticastParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSwitchSettingsMulticastParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchSettingsMulticastParams {
	var ()
	return &GetNetworkSwitchSettingsMulticastParams{
		HTTPClient: client,
	}
}

/*GetNetworkSwitchSettingsMulticastParams contains all the parameters to send to the API endpoint
for the get network switch settings multicast operation typically these are written to a http.Request
*/
type GetNetworkSwitchSettingsMulticastParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchSettingsMulticastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) WithContext(ctx context.Context) *GetNetworkSwitchSettingsMulticastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchSettingsMulticastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) WithNetworkID(networkID string) *GetNetworkSwitchSettingsMulticastParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch settings multicast params
func (o *GetNetworkSwitchSettingsMulticastParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchSettingsMulticastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
