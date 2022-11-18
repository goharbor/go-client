// Code generated by go-swagger; DO NOT EDIT.

package system_cve_allowlist

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

// NewGetSystemCVEAllowlistParams creates a new GetSystemCVEAllowlistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemCVEAllowlistParams() *GetSystemCVEAllowlistParams {
	return &GetSystemCVEAllowlistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemCVEAllowlistParamsWithTimeout creates a new GetSystemCVEAllowlistParams object
// with the ability to set a timeout on a request.
func NewGetSystemCVEAllowlistParamsWithTimeout(timeout time.Duration) *GetSystemCVEAllowlistParams {
	return &GetSystemCVEAllowlistParams{
		timeout: timeout,
	}
}

// NewGetSystemCVEAllowlistParamsWithContext creates a new GetSystemCVEAllowlistParams object
// with the ability to set a context for a request.
func NewGetSystemCVEAllowlistParamsWithContext(ctx context.Context) *GetSystemCVEAllowlistParams {
	return &GetSystemCVEAllowlistParams{
		Context: ctx,
	}
}

// NewGetSystemCVEAllowlistParamsWithHTTPClient creates a new GetSystemCVEAllowlistParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemCVEAllowlistParamsWithHTTPClient(client *http.Client) *GetSystemCVEAllowlistParams {
	return &GetSystemCVEAllowlistParams{
		HTTPClient: client,
	}
}

/*
GetSystemCVEAllowlistParams contains all the parameters to send to the API endpoint

	for the get system CVE allowlist operation.

	Typically these are written to a http.Request.
*/
type GetSystemCVEAllowlistParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system CVE allowlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemCVEAllowlistParams) WithDefaults() *GetSystemCVEAllowlistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system CVE allowlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemCVEAllowlistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) WithTimeout(timeout time.Duration) *GetSystemCVEAllowlistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) WithContext(ctx context.Context) *GetSystemCVEAllowlistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) WithHTTPClient(client *http.Client) *GetSystemCVEAllowlistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) WithXRequestID(xRequestID *string) *GetSystemCVEAllowlistParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get system CVE allowlist params
func (o *GetSystemCVEAllowlistParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemCVEAllowlistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
