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

// NewGetNetworkHTTPServersParams creates a new GetNetworkHTTPServersParams object
// with the default values initialized.
func NewGetNetworkHTTPServersParams() *GetNetworkHTTPServersParams {
	var ()
	return &GetNetworkHTTPServersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkHTTPServersParamsWithTimeout creates a new GetNetworkHTTPServersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkHTTPServersParamsWithTimeout(timeout time.Duration) *GetNetworkHTTPServersParams {
	var ()
	return &GetNetworkHTTPServersParams{

		timeout: timeout,
	}
}

// NewGetNetworkHTTPServersParamsWithContext creates a new GetNetworkHTTPServersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkHTTPServersParamsWithContext(ctx context.Context) *GetNetworkHTTPServersParams {
	var ()
	return &GetNetworkHTTPServersParams{

		Context: ctx,
	}
}

// NewGetNetworkHTTPServersParamsWithHTTPClient creates a new GetNetworkHTTPServersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkHTTPServersParamsWithHTTPClient(client *http.Client) *GetNetworkHTTPServersParams {
	var ()
	return &GetNetworkHTTPServersParams{
		HTTPClient: client,
	}
}

/*GetNetworkHTTPServersParams contains all the parameters to send to the API endpoint
for the get network Http servers operation typically these are written to a http.Request
*/
type GetNetworkHTTPServersParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network Http servers params
func (o *GetNetworkHTTPServersParams) WithTimeout(timeout time.Duration) *GetNetworkHTTPServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network Http servers params
func (o *GetNetworkHTTPServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network Http servers params
func (o *GetNetworkHTTPServersParams) WithContext(ctx context.Context) *GetNetworkHTTPServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network Http servers params
func (o *GetNetworkHTTPServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network Http servers params
func (o *GetNetworkHTTPServersParams) WithHTTPClient(client *http.Client) *GetNetworkHTTPServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network Http servers params
func (o *GetNetworkHTTPServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network Http servers params
func (o *GetNetworkHTTPServersParams) WithNetworkID(networkID string) *GetNetworkHTTPServersParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network Http servers params
func (o *GetNetworkHTTPServersParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkHTTPServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
