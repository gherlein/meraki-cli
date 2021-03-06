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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkSwitchSettingsQosRuleParams creates a new UpdateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized.
func NewUpdateNetworkSwitchSettingsQosRuleParams() *UpdateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &UpdateNetworkSwitchSettingsQosRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchSettingsQosRuleParamsWithTimeout creates a new UpdateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkSwitchSettingsQosRuleParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &UpdateNetworkSwitchSettingsQosRuleParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchSettingsQosRuleParamsWithContext creates a new UpdateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkSwitchSettingsQosRuleParamsWithContext(ctx context.Context) *UpdateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &UpdateNetworkSwitchSettingsQosRuleParams{

		Context: ctx,
	}
}

// NewUpdateNetworkSwitchSettingsQosRuleParamsWithHTTPClient creates a new UpdateNetworkSwitchSettingsQosRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkSwitchSettingsQosRuleParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchSettingsQosRuleParams {
	var ()
	return &UpdateNetworkSwitchSettingsQosRuleParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkSwitchSettingsQosRuleParams contains all the parameters to send to the API endpoint
for the update network switch settings qos rule operation typically these are written to a http.Request
*/
type UpdateNetworkSwitchSettingsQosRuleParams struct {

	/*NetworkID*/
	NetworkID string
	/*QosRuleID*/
	QosRuleID string
	/*UpdateNetworkSwitchSettingsQosRule*/
	UpdateNetworkSwitchSettingsQosRule *models.UpdateNetworkSwitchSettingsQosRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchSettingsQosRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WithContext(ctx context.Context) *UpdateNetworkSwitchSettingsQosRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchSettingsQosRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WithNetworkID(networkID string) *UpdateNetworkSwitchSettingsQosRuleParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithQosRuleID adds the qosRuleID to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WithQosRuleID(qosRuleID string) *UpdateNetworkSwitchSettingsQosRuleParams {
	o.SetQosRuleID(qosRuleID)
	return o
}

// SetQosRuleID adds the qosRuleId to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) SetQosRuleID(qosRuleID string) {
	o.QosRuleID = qosRuleID
}

// WithUpdateNetworkSwitchSettingsQosRule adds the updateNetworkSwitchSettingsQosRule to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WithUpdateNetworkSwitchSettingsQosRule(updateNetworkSwitchSettingsQosRule *models.UpdateNetworkSwitchSettingsQosRule) *UpdateNetworkSwitchSettingsQosRuleParams {
	o.SetUpdateNetworkSwitchSettingsQosRule(updateNetworkSwitchSettingsQosRule)
	return o
}

// SetUpdateNetworkSwitchSettingsQosRule adds the updateNetworkSwitchSettingsQosRule to the update network switch settings qos rule params
func (o *UpdateNetworkSwitchSettingsQosRuleParams) SetUpdateNetworkSwitchSettingsQosRule(updateNetworkSwitchSettingsQosRule *models.UpdateNetworkSwitchSettingsQosRule) {
	o.UpdateNetworkSwitchSettingsQosRule = updateNetworkSwitchSettingsQosRule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchSettingsQosRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param qosRuleId
	if err := r.SetPathParam("qosRuleId", o.QosRuleID); err != nil {
		return err
	}

	if o.UpdateNetworkSwitchSettingsQosRule != nil {
		if err := r.SetBodyParam(o.UpdateNetworkSwitchSettingsQosRule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
