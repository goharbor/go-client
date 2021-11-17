// Code generated by go-swagger; DO NOT EDIT.

package member

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

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewCreateProjectMemberParams creates a new CreateProjectMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProjectMemberParams() *CreateProjectMemberParams {
	return &CreateProjectMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectMemberParamsWithTimeout creates a new CreateProjectMemberParams object
// with the ability to set a timeout on a request.
func NewCreateProjectMemberParamsWithTimeout(timeout time.Duration) *CreateProjectMemberParams {
	return &CreateProjectMemberParams{
		timeout: timeout,
	}
}

// NewCreateProjectMemberParamsWithContext creates a new CreateProjectMemberParams object
// with the ability to set a context for a request.
func NewCreateProjectMemberParamsWithContext(ctx context.Context) *CreateProjectMemberParams {
	return &CreateProjectMemberParams{
		Context: ctx,
	}
}

// NewCreateProjectMemberParamsWithHTTPClient creates a new CreateProjectMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProjectMemberParamsWithHTTPClient(client *http.Client) *CreateProjectMemberParams {
	return &CreateProjectMemberParams{
		HTTPClient: client,
	}
}

/* CreateProjectMemberParams contains all the parameters to send to the API endpoint
   for the create project member operation.

   Typically these are written to a http.Request.
*/
type CreateProjectMemberParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	// ProjectMember.
	ProjectMember *models.ProjectMember

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectMemberParams) WithDefaults() *CreateProjectMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectMemberParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := CreateProjectMemberParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create project member params
func (o *CreateProjectMemberParams) WithTimeout(timeout time.Duration) *CreateProjectMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project member params
func (o *CreateProjectMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project member params
func (o *CreateProjectMemberParams) WithContext(ctx context.Context) *CreateProjectMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project member params
func (o *CreateProjectMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project member params
func (o *CreateProjectMemberParams) WithHTTPClient(client *http.Client) *CreateProjectMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project member params
func (o *CreateProjectMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the create project member params
func (o *CreateProjectMemberParams) WithXIsResourceName(xIsResourceName *bool) *CreateProjectMemberParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the create project member params
func (o *CreateProjectMemberParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the create project member params
func (o *CreateProjectMemberParams) WithXRequestID(xRequestID *string) *CreateProjectMemberParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create project member params
func (o *CreateProjectMemberParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectMember adds the projectMember to the create project member params
func (o *CreateProjectMemberParams) WithProjectMember(projectMember *models.ProjectMember) *CreateProjectMemberParams {
	o.SetProjectMember(projectMember)
	return o
}

// SetProjectMember adds the projectMember to the create project member params
func (o *CreateProjectMemberParams) SetProjectMember(projectMember *models.ProjectMember) {
	o.ProjectMember = projectMember
}

// WithProjectNameOrID adds the projectNameOrID to the create project member params
func (o *CreateProjectMemberParams) WithProjectNameOrID(projectNameOrID string) *CreateProjectMemberParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the create project member params
func (o *CreateProjectMemberParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.ProjectMember != nil {
		if err := r.SetBodyParam(o.ProjectMember); err != nil {
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
