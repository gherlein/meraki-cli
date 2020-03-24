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

// NewGetNetworkHTTPServerParams creates a new GetNetworkHTTPServerParams object
// with the default values initialized.
func NewGetNetworkHTTPServerParams() *GetNetworkHTTPServerParams {
	var ()
	return &GetNetworkHTTPServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkHTTPServerParamsWithTimeout creates a new GetNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkHTTPServerParamsWithTimeout(timeout time.Duration) *GetNetworkHTTPServerParams {
	var ()
	return &GetNetworkHTTPServerParams{

		timeout: timeout,
	}
}

// NewGetNetworkHTTPServerParamsWithContext creates a new GetNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkHTTPServerParamsWithContext(ctx context.Context) *GetNetworkHTTPServerParams {
	var ()
	return &GetNetworkHTTPServerParams{

		Context: ctx,
	}
}

// NewGetNetworkHTTPServerParamsWithHTTPClient creates a new GetNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkHTTPServerParamsWithHTTPClient(client *http.Client) *GetNetworkHTTPServerParams {
	var ()
	return &GetNetworkHTTPServerParams{
		HTTPClient: client,
	}
}

/*GetNetworkHTTPServerParams contains all the parameters to send to the API endpoint
for the get network Http server operation typically these are written to a http.Request
*/
type GetNetworkHTTPServerParams struct {

	/*ID*/
	ID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network Http server params
func (o *GetNetworkHTTPServerParams) WithTimeout(timeout time.Duration) *GetNetworkHTTPServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network Http server params
func (o *GetNetworkHTTPServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network Http server params
func (o *GetNetworkHTTPServerParams) WithContext(ctx context.Context) *GetNetworkHTTPServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network Http server params
func (o *GetNetworkHTTPServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network Http server params
func (o *GetNetworkHTTPServerParams) WithHTTPClient(client *http.Client) *GetNetworkHTTPServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network Http server params
func (o *GetNetworkHTTPServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network Http server params
func (o *GetNetworkHTTPServerParams) WithID(id string) *GetNetworkHTTPServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network Http server params
func (o *GetNetworkHTTPServerParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the get network Http server params
func (o *GetNetworkHTTPServerParams) WithNetworkID(networkID string) *GetNetworkHTTPServerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network Http server params
func (o *GetNetworkHTTPServerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkHTTPServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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