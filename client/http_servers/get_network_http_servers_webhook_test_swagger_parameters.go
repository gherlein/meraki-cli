// Code generated by go-swagger; DO NOT EDIT.

package http_servers

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

// NewGetNetworkHTTPServersWebhookTestParams creates a new GetNetworkHTTPServersWebhookTestParams object
// with the default values initialized.
func NewGetNetworkHTTPServersWebhookTestParams() *GetNetworkHTTPServersWebhookTestParams {
	var ()
	return &GetNetworkHTTPServersWebhookTestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkHTTPServersWebhookTestParamsWithTimeout creates a new GetNetworkHTTPServersWebhookTestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkHTTPServersWebhookTestParamsWithTimeout(timeout time.Duration) *GetNetworkHTTPServersWebhookTestParams {
	var ()
	return &GetNetworkHTTPServersWebhookTestParams{

		timeout: timeout,
	}
}

// NewGetNetworkHTTPServersWebhookTestParamsWithContext creates a new GetNetworkHTTPServersWebhookTestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkHTTPServersWebhookTestParamsWithContext(ctx context.Context) *GetNetworkHTTPServersWebhookTestParams {
	var ()
	return &GetNetworkHTTPServersWebhookTestParams{

		Context: ctx,
	}
}

// NewGetNetworkHTTPServersWebhookTestParamsWithHTTPClient creates a new GetNetworkHTTPServersWebhookTestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkHTTPServersWebhookTestParamsWithHTTPClient(client *http.Client) *GetNetworkHTTPServersWebhookTestParams {
	var ()
	return &GetNetworkHTTPServersWebhookTestParams{
		HTTPClient: client,
	}
}

/*GetNetworkHTTPServersWebhookTestParams contains all the parameters to send to the API endpoint
for the get network Http servers webhook test operation typically these are written to a http.Request
*/
type GetNetworkHTTPServersWebhookTestParams struct {

	/*ID*/
	ID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) WithTimeout(timeout time.Duration) *GetNetworkHTTPServersWebhookTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) WithContext(ctx context.Context) *GetNetworkHTTPServersWebhookTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) WithHTTPClient(client *http.Client) *GetNetworkHTTPServersWebhookTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) WithID(id string) *GetNetworkHTTPServersWebhookTestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) WithNetworkID(networkID string) *GetNetworkHTTPServersWebhookTestParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network Http servers webhook test params
func (o *GetNetworkHTTPServersWebhookTestParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkHTTPServersWebhookTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}