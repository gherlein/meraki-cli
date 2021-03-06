// Code generated by go-swagger; DO NOT EDIT.

package radio_settings

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

// NewGetNetworkDeviceWirelessRadioSettingsParams creates a new GetNetworkDeviceWirelessRadioSettingsParams object
// with the default values initialized.
func NewGetNetworkDeviceWirelessRadioSettingsParams() *GetNetworkDeviceWirelessRadioSettingsParams {
	var ()
	return &GetNetworkDeviceWirelessRadioSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkDeviceWirelessRadioSettingsParamsWithTimeout creates a new GetNetworkDeviceWirelessRadioSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkDeviceWirelessRadioSettingsParamsWithTimeout(timeout time.Duration) *GetNetworkDeviceWirelessRadioSettingsParams {
	var ()
	return &GetNetworkDeviceWirelessRadioSettingsParams{

		timeout: timeout,
	}
}

// NewGetNetworkDeviceWirelessRadioSettingsParamsWithContext creates a new GetNetworkDeviceWirelessRadioSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkDeviceWirelessRadioSettingsParamsWithContext(ctx context.Context) *GetNetworkDeviceWirelessRadioSettingsParams {
	var ()
	return &GetNetworkDeviceWirelessRadioSettingsParams{

		Context: ctx,
	}
}

// NewGetNetworkDeviceWirelessRadioSettingsParamsWithHTTPClient creates a new GetNetworkDeviceWirelessRadioSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkDeviceWirelessRadioSettingsParamsWithHTTPClient(client *http.Client) *GetNetworkDeviceWirelessRadioSettingsParams {
	var ()
	return &GetNetworkDeviceWirelessRadioSettingsParams{
		HTTPClient: client,
	}
}

/*GetNetworkDeviceWirelessRadioSettingsParams contains all the parameters to send to the API endpoint
for the get network device wireless radio settings operation typically these are written to a http.Request
*/
type GetNetworkDeviceWirelessRadioSettingsParams struct {

	/*NetworkID*/
	NetworkID string
	/*Serial*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) WithTimeout(timeout time.Duration) *GetNetworkDeviceWirelessRadioSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) WithContext(ctx context.Context) *GetNetworkDeviceWirelessRadioSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) WithHTTPClient(client *http.Client) *GetNetworkDeviceWirelessRadioSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) WithNetworkID(networkID string) *GetNetworkDeviceWirelessRadioSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSerial adds the serial to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) WithSerial(serial string) *GetNetworkDeviceWirelessRadioSettingsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get network device wireless radio settings params
func (o *GetNetworkDeviceWirelessRadioSettingsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkDeviceWirelessRadioSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
