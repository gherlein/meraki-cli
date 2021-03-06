// Code generated by go-swagger; DO NOT EDIT.

package content_filtering_rules

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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkContentFilteringParams creates a new UpdateNetworkContentFilteringParams object
// with the default values initialized.
func NewUpdateNetworkContentFilteringParams() *UpdateNetworkContentFilteringParams {
	var ()
	return &UpdateNetworkContentFilteringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkContentFilteringParamsWithTimeout creates a new UpdateNetworkContentFilteringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkContentFilteringParamsWithTimeout(timeout time.Duration) *UpdateNetworkContentFilteringParams {
	var ()
	return &UpdateNetworkContentFilteringParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkContentFilteringParamsWithContext creates a new UpdateNetworkContentFilteringParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkContentFilteringParamsWithContext(ctx context.Context) *UpdateNetworkContentFilteringParams {
	var ()
	return &UpdateNetworkContentFilteringParams{

		Context: ctx,
	}
}

// NewUpdateNetworkContentFilteringParamsWithHTTPClient creates a new UpdateNetworkContentFilteringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkContentFilteringParamsWithHTTPClient(client *http.Client) *UpdateNetworkContentFilteringParams {
	var ()
	return &UpdateNetworkContentFilteringParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkContentFilteringParams contains all the parameters to send to the API endpoint
for the update network content filtering operation typically these are written to a http.Request
*/
type UpdateNetworkContentFilteringParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkContentFiltering*/
	UpdateNetworkContentFiltering *models.UpdateNetworkContentFiltering

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) WithTimeout(timeout time.Duration) *UpdateNetworkContentFilteringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) WithContext(ctx context.Context) *UpdateNetworkContentFilteringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) WithHTTPClient(client *http.Client) *UpdateNetworkContentFilteringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) WithNetworkID(networkID string) *UpdateNetworkContentFilteringParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkContentFiltering adds the updateNetworkContentFiltering to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) WithUpdateNetworkContentFiltering(updateNetworkContentFiltering *models.UpdateNetworkContentFiltering) *UpdateNetworkContentFilteringParams {
	o.SetUpdateNetworkContentFiltering(updateNetworkContentFiltering)
	return o
}

// SetUpdateNetworkContentFiltering adds the updateNetworkContentFiltering to the update network content filtering params
func (o *UpdateNetworkContentFilteringParams) SetUpdateNetworkContentFiltering(updateNetworkContentFiltering *models.UpdateNetworkContentFiltering) {
	o.UpdateNetworkContentFiltering = updateNetworkContentFiltering
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkContentFilteringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkContentFiltering != nil {
		if err := r.SetBodyParam(o.UpdateNetworkContentFiltering); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
