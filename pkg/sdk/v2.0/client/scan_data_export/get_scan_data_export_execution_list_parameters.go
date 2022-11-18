// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

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

// NewGetScanDataExportExecutionListParams creates a new GetScanDataExportExecutionListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScanDataExportExecutionListParams() *GetScanDataExportExecutionListParams {
	return &GetScanDataExportExecutionListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanDataExportExecutionListParamsWithTimeout creates a new GetScanDataExportExecutionListParams object
// with the ability to set a timeout on a request.
func NewGetScanDataExportExecutionListParamsWithTimeout(timeout time.Duration) *GetScanDataExportExecutionListParams {
	return &GetScanDataExportExecutionListParams{
		timeout: timeout,
	}
}

// NewGetScanDataExportExecutionListParamsWithContext creates a new GetScanDataExportExecutionListParams object
// with the ability to set a context for a request.
func NewGetScanDataExportExecutionListParamsWithContext(ctx context.Context) *GetScanDataExportExecutionListParams {
	return &GetScanDataExportExecutionListParams{
		Context: ctx,
	}
}

// NewGetScanDataExportExecutionListParamsWithHTTPClient creates a new GetScanDataExportExecutionListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScanDataExportExecutionListParamsWithHTTPClient(client *http.Client) *GetScanDataExportExecutionListParams {
	return &GetScanDataExportExecutionListParams{
		HTTPClient: client,
	}
}

/*
GetScanDataExportExecutionListParams contains all the parameters to send to the API endpoint

	for the get scan data export execution list operation.

	Typically these are written to a http.Request.
*/
type GetScanDataExportExecutionListParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scan data export execution list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScanDataExportExecutionListParams) WithDefaults() *GetScanDataExportExecutionListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scan data export execution list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScanDataExportExecutionListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) WithTimeout(timeout time.Duration) *GetScanDataExportExecutionListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) WithContext(ctx context.Context) *GetScanDataExportExecutionListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) WithHTTPClient(client *http.Client) *GetScanDataExportExecutionListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) WithXRequestID(xRequestID *string) *GetScanDataExportExecutionListParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get scan data export execution list params
func (o *GetScanDataExportExecutionListParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanDataExportExecutionListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
