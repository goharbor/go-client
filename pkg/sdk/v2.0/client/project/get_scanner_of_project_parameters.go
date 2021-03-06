// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetScannerOfProjectParams creates a new GetScannerOfProjectParams object
// with the default values initialized.
func NewGetScannerOfProjectParams() *GetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &GetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScannerOfProjectParamsWithTimeout creates a new GetScannerOfProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScannerOfProjectParamsWithTimeout(timeout time.Duration) *GetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &GetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: timeout,
	}
}

// NewGetScannerOfProjectParamsWithContext creates a new GetScannerOfProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScannerOfProjectParamsWithContext(ctx context.Context) *GetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &GetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,

		Context: ctx,
	}
}

// NewGetScannerOfProjectParamsWithHTTPClient creates a new GetScannerOfProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScannerOfProjectParamsWithHTTPClient(client *http.Client) *GetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &GetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
		HTTPClient:      client,
	}
}

/*GetScannerOfProjectParams contains all the parameters to send to the API endpoint
for the get scanner of project operation typically these are written to a http.Request
*/
type GetScannerOfProjectParams struct {

	/*XIsResourceName
	  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.

	*/
	XIsResourceName *bool
	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ProjectNameOrID
	  The name or id of the project

	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scanner of project params
func (o *GetScannerOfProjectParams) WithTimeout(timeout time.Duration) *GetScannerOfProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scanner of project params
func (o *GetScannerOfProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scanner of project params
func (o *GetScannerOfProjectParams) WithContext(ctx context.Context) *GetScannerOfProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scanner of project params
func (o *GetScannerOfProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scanner of project params
func (o *GetScannerOfProjectParams) WithHTTPClient(client *http.Client) *GetScannerOfProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scanner of project params
func (o *GetScannerOfProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the get scanner of project params
func (o *GetScannerOfProjectParams) WithXIsResourceName(xIsResourceName *bool) *GetScannerOfProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the get scanner of project params
func (o *GetScannerOfProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the get scanner of project params
func (o *GetScannerOfProjectParams) WithXRequestID(xRequestID *string) *GetScannerOfProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get scanner of project params
func (o *GetScannerOfProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the get scanner of project params
func (o *GetScannerOfProjectParams) WithProjectNameOrID(projectNameOrID string) *GetScannerOfProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the get scanner of project params
func (o *GetScannerOfProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScannerOfProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
