// Code generated by go-swagger; DO NOT EDIT.

package auditlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the auditlog client
type API interface {
	/*
	   ListAuditLogs gets recent logs of the projects which the user is a member of

	   This endpoint let user see the recent operation logs of the projects which he is member of
	*/
	ListAuditLogs(ctx context.Context, params *ListAuditLogsParams) (*ListAuditLogsOK, error)
}

// New creates a new auditlog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for auditlog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ListAuditLogs gets recent logs of the projects which the user is a member of

This endpoint let user see the recent operation logs of the projects which he is member of
*/
func (a *Client) ListAuditLogs(ctx context.Context, params *ListAuditLogsParams) (*ListAuditLogsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAuditLogs",
		Method:             "GET",
		PathPattern:        "/audit-logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAuditLogsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAuditLogsOK), nil

}
