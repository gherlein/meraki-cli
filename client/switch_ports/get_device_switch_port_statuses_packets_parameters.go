// Code generated by go-swagger; DO NOT EDIT.

package switch_ports

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
	"github.com/go-openapi/swag"
)

// NewGetDeviceSwitchPortStatusesPacketsParams creates a new GetDeviceSwitchPortStatusesPacketsParams object
// with the default values initialized.
func NewGetDeviceSwitchPortStatusesPacketsParams() *GetDeviceSwitchPortStatusesPacketsParams {
	var ()
	return &GetDeviceSwitchPortStatusesPacketsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceSwitchPortStatusesPacketsParamsWithTimeout creates a new GetDeviceSwitchPortStatusesPacketsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceSwitchPortStatusesPacketsParamsWithTimeout(timeout time.Duration) *GetDeviceSwitchPortStatusesPacketsParams {
	var ()
	return &GetDeviceSwitchPortStatusesPacketsParams{

		timeout: timeout,
	}
}

// NewGetDeviceSwitchPortStatusesPacketsParamsWithContext creates a new GetDeviceSwitchPortStatusesPacketsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceSwitchPortStatusesPacketsParamsWithContext(ctx context.Context) *GetDeviceSwitchPortStatusesPacketsParams {
	var ()
	return &GetDeviceSwitchPortStatusesPacketsParams{

		Context: ctx,
	}
}

// NewGetDeviceSwitchPortStatusesPacketsParamsWithHTTPClient creates a new GetDeviceSwitchPortStatusesPacketsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceSwitchPortStatusesPacketsParamsWithHTTPClient(client *http.Client) *GetDeviceSwitchPortStatusesPacketsParams {
	var ()
	return &GetDeviceSwitchPortStatusesPacketsParams{
		HTTPClient: client,
	}
}

/*GetDeviceSwitchPortStatusesPacketsParams contains all the parameters to send to the API endpoint
for the get device switch port statuses packets operation typically these are written to a http.Request
*/
type GetDeviceSwitchPortStatusesPacketsParams struct {

	/*Serial*/
	Serial string
	/*T0
	  The beginning of the timespan for the data. The maximum lookback period is 1 day from today.

	*/
	T0 *string
	/*Timespan
	  The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day.

	*/
	Timespan *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) WithTimeout(timeout time.Duration) *GetDeviceSwitchPortStatusesPacketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) WithContext(ctx context.Context) *GetDeviceSwitchPortStatusesPacketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) WithHTTPClient(client *http.Client) *GetDeviceSwitchPortStatusesPacketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) WithSerial(serial string) *GetDeviceSwitchPortStatusesPacketsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithT0 adds the t0 to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) WithT0(t0 *string) *GetDeviceSwitchPortStatusesPacketsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithTimespan adds the timespan to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) WithTimespan(timespan *float64) *GetDeviceSwitchPortStatusesPacketsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get device switch port statuses packets params
func (o *GetDeviceSwitchPortStatusesPacketsParams) SetTimespan(timespan *float64) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceSwitchPortStatusesPacketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string
		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {
			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}

	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float64
		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat64(qrTimespan)
		if qTimespan != "" {
			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}