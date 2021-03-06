// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the replication client
type API interface {
	/*
	   CreateReplicationPolicy creates a replication policy

	   Create a replication policy*/
	CreateReplicationPolicy(ctx context.Context, params *CreateReplicationPolicyParams) (*CreateReplicationPolicyCreated, error)
	/*
	   DeleteReplicationPolicy deletes the specific replication policy

	   Delete the specific replication policy*/
	DeleteReplicationPolicy(ctx context.Context, params *DeleteReplicationPolicyParams) (*DeleteReplicationPolicyOK, error)
	/*
	   GetReplicationExecution gets the specific replication execution

	   Get the replication execution specified by ID*/
	GetReplicationExecution(ctx context.Context, params *GetReplicationExecutionParams) (*GetReplicationExecutionOK, error)
	/*
	   GetReplicationLog gets the log of the specific replication task

	   Get the log of the specific replication task*/
	GetReplicationLog(ctx context.Context, params *GetReplicationLogParams) (*GetReplicationLogOK, error)
	/*
	   GetReplicationPolicy gets the specific replication policy

	   Get the specific replication policy*/
	GetReplicationPolicy(ctx context.Context, params *GetReplicationPolicyParams) (*GetReplicationPolicyOK, error)
	/*
	   ListReplicationExecutions lists replication executions

	   List replication executions*/
	ListReplicationExecutions(ctx context.Context, params *ListReplicationExecutionsParams) (*ListReplicationExecutionsOK, error)
	/*
	   ListReplicationPolicies lists replication policies

	   List replication policies*/
	ListReplicationPolicies(ctx context.Context, params *ListReplicationPoliciesParams) (*ListReplicationPoliciesOK, error)
	/*
	   ListReplicationTasks lists replication tasks for a specific execution

	   List replication tasks for a specific execution*/
	ListReplicationTasks(ctx context.Context, params *ListReplicationTasksParams) (*ListReplicationTasksOK, error)
	/*
	   StartReplication starts one replication execution

	   Start one replication execution according to the policy*/
	StartReplication(ctx context.Context, params *StartReplicationParams) (*StartReplicationCreated, error)
	/*
	   StopReplication stops the specific replication execution

	   Stop the replication execution specified by ID*/
	StopReplication(ctx context.Context, params *StopReplicationParams) (*StopReplicationOK, error)
	/*
	   UpdateReplicationPolicy updates the replication policy

	   Update the replication policy*/
	UpdateReplicationPolicy(ctx context.Context, params *UpdateReplicationPolicyParams) (*UpdateReplicationPolicyOK, error)
}

// New creates a new replication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for replication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateReplicationPolicy creates a replication policy

Create a replication policy
*/
func (a *Client) CreateReplicationPolicy(ctx context.Context, params *CreateReplicationPolicyParams) (*CreateReplicationPolicyCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createReplicationPolicy",
		Method:             "POST",
		PathPattern:        "/replication/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateReplicationPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateReplicationPolicyCreated), nil

}

/*
DeleteReplicationPolicy deletes the specific replication policy

Delete the specific replication policy
*/
func (a *Client) DeleteReplicationPolicy(ctx context.Context, params *DeleteReplicationPolicyParams) (*DeleteReplicationPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteReplicationPolicy",
		Method:             "DELETE",
		PathPattern:        "/replication/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReplicationPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteReplicationPolicyOK), nil

}

/*
GetReplicationExecution gets the specific replication execution

Get the replication execution specified by ID
*/
func (a *Client) GetReplicationExecution(ctx context.Context, params *GetReplicationExecutionParams) (*GetReplicationExecutionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReplicationExecution",
		Method:             "GET",
		PathPattern:        "/replication/executions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReplicationExecutionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReplicationExecutionOK), nil

}

/*
GetReplicationLog gets the log of the specific replication task

Get the log of the specific replication task
*/
func (a *Client) GetReplicationLog(ctx context.Context, params *GetReplicationLogParams) (*GetReplicationLogOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReplicationLog",
		Method:             "GET",
		PathPattern:        "/replication/executions/{id}/tasks/{task_id}/log",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReplicationLogReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReplicationLogOK), nil

}

/*
GetReplicationPolicy gets the specific replication policy

Get the specific replication policy
*/
func (a *Client) GetReplicationPolicy(ctx context.Context, params *GetReplicationPolicyParams) (*GetReplicationPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReplicationPolicy",
		Method:             "GET",
		PathPattern:        "/replication/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReplicationPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReplicationPolicyOK), nil

}

/*
ListReplicationExecutions lists replication executions

List replication executions
*/
func (a *Client) ListReplicationExecutions(ctx context.Context, params *ListReplicationExecutionsParams) (*ListReplicationExecutionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReplicationExecutions",
		Method:             "GET",
		PathPattern:        "/replication/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReplicationExecutionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReplicationExecutionsOK), nil

}

/*
ListReplicationPolicies lists replication policies

List replication policies
*/
func (a *Client) ListReplicationPolicies(ctx context.Context, params *ListReplicationPoliciesParams) (*ListReplicationPoliciesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReplicationPolicies",
		Method:             "GET",
		PathPattern:        "/replication/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReplicationPoliciesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReplicationPoliciesOK), nil

}

/*
ListReplicationTasks lists replication tasks for a specific execution

List replication tasks for a specific execution
*/
func (a *Client) ListReplicationTasks(ctx context.Context, params *ListReplicationTasksParams) (*ListReplicationTasksOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReplicationTasks",
		Method:             "GET",
		PathPattern:        "/replication/executions/{id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReplicationTasksReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReplicationTasksOK), nil

}

/*
StartReplication starts one replication execution

Start one replication execution according to the policy
*/
func (a *Client) StartReplication(ctx context.Context, params *StartReplicationParams) (*StartReplicationCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startReplication",
		Method:             "POST",
		PathPattern:        "/replication/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartReplicationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartReplicationCreated), nil

}

/*
StopReplication stops the specific replication execution

Stop the replication execution specified by ID
*/
func (a *Client) StopReplication(ctx context.Context, params *StopReplicationParams) (*StopReplicationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopReplication",
		Method:             "PUT",
		PathPattern:        "/replication/executions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopReplicationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopReplicationOK), nil

}

/*
UpdateReplicationPolicy updates the replication policy

Update the replication policy
*/
func (a *Client) UpdateReplicationPolicy(ctx context.Context, params *UpdateReplicationPolicyParams) (*UpdateReplicationPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateReplicationPolicy",
		Method:             "PUT",
		PathPattern:        "/replication/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateReplicationPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateReplicationPolicyOK), nil

}
