// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// ActionPendingJobsReader is a Reader for the ActionPendingJobs structure.
type ActionPendingJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionPendingJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionPendingJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActionPendingJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActionPendingJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActionPendingJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActionPendingJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActionPendingJobsOK creates a ActionPendingJobsOK with default headers values
func NewActionPendingJobsOK() *ActionPendingJobsOK {
	return &ActionPendingJobsOK{}
}

/*
ActionPendingJobsOK describes a response with status code 200, with default header values.

take action to the jobs in the queue successfully.
*/
type ActionPendingJobsOK struct {
}

// IsSuccess returns true when this action pending jobs o k response has a 2xx status code
func (o *ActionPendingJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action pending jobs o k response has a 3xx status code
func (o *ActionPendingJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action pending jobs o k response has a 4xx status code
func (o *ActionPendingJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action pending jobs o k response has a 5xx status code
func (o *ActionPendingJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action pending jobs o k response a status code equal to that given
func (o *ActionPendingJobsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ActionPendingJobsOK) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsOK ", 200)
}

func (o *ActionPendingJobsOK) String() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsOK ", 200)
}

func (o *ActionPendingJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionPendingJobsUnauthorized creates a ActionPendingJobsUnauthorized with default headers values
func NewActionPendingJobsUnauthorized() *ActionPendingJobsUnauthorized {
	return &ActionPendingJobsUnauthorized{}
}

/*
ActionPendingJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ActionPendingJobsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action pending jobs unauthorized response has a 2xx status code
func (o *ActionPendingJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action pending jobs unauthorized response has a 3xx status code
func (o *ActionPendingJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action pending jobs unauthorized response has a 4xx status code
func (o *ActionPendingJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this action pending jobs unauthorized response has a 5xx status code
func (o *ActionPendingJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this action pending jobs unauthorized response a status code equal to that given
func (o *ActionPendingJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ActionPendingJobsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *ActionPendingJobsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *ActionPendingJobsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionPendingJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionPendingJobsForbidden creates a ActionPendingJobsForbidden with default headers values
func NewActionPendingJobsForbidden() *ActionPendingJobsForbidden {
	return &ActionPendingJobsForbidden{}
}

/*
ActionPendingJobsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ActionPendingJobsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action pending jobs forbidden response has a 2xx status code
func (o *ActionPendingJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action pending jobs forbidden response has a 3xx status code
func (o *ActionPendingJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action pending jobs forbidden response has a 4xx status code
func (o *ActionPendingJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this action pending jobs forbidden response has a 5xx status code
func (o *ActionPendingJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this action pending jobs forbidden response a status code equal to that given
func (o *ActionPendingJobsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ActionPendingJobsForbidden) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsForbidden  %+v", 403, o.Payload)
}

func (o *ActionPendingJobsForbidden) String() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsForbidden  %+v", 403, o.Payload)
}

func (o *ActionPendingJobsForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionPendingJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionPendingJobsNotFound creates a ActionPendingJobsNotFound with default headers values
func NewActionPendingJobsNotFound() *ActionPendingJobsNotFound {
	return &ActionPendingJobsNotFound{}
}

/*
ActionPendingJobsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ActionPendingJobsNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action pending jobs not found response has a 2xx status code
func (o *ActionPendingJobsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action pending jobs not found response has a 3xx status code
func (o *ActionPendingJobsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action pending jobs not found response has a 4xx status code
func (o *ActionPendingJobsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this action pending jobs not found response has a 5xx status code
func (o *ActionPendingJobsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this action pending jobs not found response a status code equal to that given
func (o *ActionPendingJobsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ActionPendingJobsNotFound) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsNotFound  %+v", 404, o.Payload)
}

func (o *ActionPendingJobsNotFound) String() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsNotFound  %+v", 404, o.Payload)
}

func (o *ActionPendingJobsNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionPendingJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionPendingJobsInternalServerError creates a ActionPendingJobsInternalServerError with default headers values
func NewActionPendingJobsInternalServerError() *ActionPendingJobsInternalServerError {
	return &ActionPendingJobsInternalServerError{}
}

/*
ActionPendingJobsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ActionPendingJobsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action pending jobs internal server error response has a 2xx status code
func (o *ActionPendingJobsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action pending jobs internal server error response has a 3xx status code
func (o *ActionPendingJobsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action pending jobs internal server error response has a 4xx status code
func (o *ActionPendingJobsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this action pending jobs internal server error response has a 5xx status code
func (o *ActionPendingJobsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this action pending jobs internal server error response a status code equal to that given
func (o *ActionPendingJobsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ActionPendingJobsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionPendingJobsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionPendingJobsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionPendingJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}