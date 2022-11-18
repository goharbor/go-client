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

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewPutSystemCVEAllowlistParams creates a new PutSystemCVEAllowlistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutSystemCVEAllowlistParams() *PutSystemCVEAllowlistParams {
	return &PutSystemCVEAllowlistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutSystemCVEAllowlistParamsWithTimeout creates a new PutSystemCVEAllowlistParams object
// with the ability to set a timeout on a request.
func NewPutSystemCVEAllowlistParamsWithTimeout(timeout time.Duration) *PutSystemCVEAllowlistParams {
	return &PutSystemCVEAllowlistParams{
		timeout: timeout,
	}
}

// NewPutSystemCVEAllowlistParamsWithContext creates a new PutSystemCVEAllowlistParams object
// with the ability to set a context for a request.
func NewPutSystemCVEAllowlistParamsWithContext(ctx context.Context) *PutSystemCVEAllowlistParams {
	return &PutSystemCVEAllowlistParams{
		Context: ctx,
	}
}

// NewPutSystemCVEAllowlistParamsWithHTTPClient creates a new PutSystemCVEAllowlistParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutSystemCVEAllowlistParamsWithHTTPClient(client *http.Client) *PutSystemCVEAllowlistParams {
	return &PutSystemCVEAllowlistParams{
		HTTPClient: client,
	}
}

/*
PutSystemCVEAllowlistParams contains all the parameters to send to the API endpoint

	for the put system CVE allowlist operation.

	Typically these are written to a http.Request.
*/
type PutSystemCVEAllowlistParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Allowlist.

	   The allowlist with new content
	*/
	Allowlist *models.CVEAllowlist

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put system CVE allowlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSystemCVEAllowlistParams) WithDefaults() *PutSystemCVEAllowlistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put system CVE allowlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSystemCVEAllowlistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) WithTimeout(timeout time.Duration) *PutSystemCVEAllowlistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) WithContext(ctx context.Context) *PutSystemCVEAllowlistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) WithHTTPClient(client *http.Client) *PutSystemCVEAllowlistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) WithXRequestID(xRequestID *string) *PutSystemCVEAllowlistParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAllowlist adds the allowlist to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) WithAllowlist(allowlist *models.CVEAllowlist) *PutSystemCVEAllowlistParams {
	o.SetAllowlist(allowlist)
	return o
}

// SetAllowlist adds the allowlist to the put system CVE allowlist params
func (o *PutSystemCVEAllowlistParams) SetAllowlist(allowlist *models.CVEAllowlist) {
	o.Allowlist = allowlist
}

// WriteToRequest writes these params to a swagger request
func (o *PutSystemCVEAllowlistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Allowlist != nil {
		if err := r.SetBodyParam(o.Allowlist); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
