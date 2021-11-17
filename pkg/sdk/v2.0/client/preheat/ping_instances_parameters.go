// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewPingInstancesParams creates a new PingInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPingInstancesParams() *PingInstancesParams {
	return &PingInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPingInstancesParamsWithTimeout creates a new PingInstancesParams object
// with the ability to set a timeout on a request.
func NewPingInstancesParamsWithTimeout(timeout time.Duration) *PingInstancesParams {
	return &PingInstancesParams{
		timeout: timeout,
	}
}

// NewPingInstancesParamsWithContext creates a new PingInstancesParams object
// with the ability to set a context for a request.
func NewPingInstancesParamsWithContext(ctx context.Context) *PingInstancesParams {
	return &PingInstancesParams{
		Context: ctx,
	}
}

// NewPingInstancesParamsWithHTTPClient creates a new PingInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPingInstancesParamsWithHTTPClient(client *http.Client) *PingInstancesParams {
	return &PingInstancesParams{
		HTTPClient: client,
	}
}

/* PingInstancesParams contains all the parameters to send to the API endpoint
   for the ping instances operation.

   Typically these are written to a http.Request.
*/
type PingInstancesParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Instance.

	   The JSON object of instance.
	*/
	Instance *models.Instance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ping instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingInstancesParams) WithDefaults() *PingInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ping instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ping instances params
func (o *PingInstancesParams) WithTimeout(timeout time.Duration) *PingInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping instances params
func (o *PingInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping instances params
func (o *PingInstancesParams) WithContext(ctx context.Context) *PingInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping instances params
func (o *PingInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping instances params
func (o *PingInstancesParams) WithHTTPClient(client *http.Client) *PingInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping instances params
func (o *PingInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the ping instances params
func (o *PingInstancesParams) WithXRequestID(xRequestID *string) *PingInstancesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the ping instances params
func (o *PingInstancesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithInstance adds the instance to the ping instances params
func (o *PingInstancesParams) WithInstance(instance *models.Instance) *PingInstancesParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the ping instances params
func (o *PingInstancesParams) SetInstance(instance *models.Instance) {
	o.Instance = instance
}

// WriteToRequest writes these params to a swagger request
func (o *PingInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Instance != nil {
		if err := r.SetBodyParam(o.Instance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
