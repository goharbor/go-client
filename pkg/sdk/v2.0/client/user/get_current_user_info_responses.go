// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetCurrentUserInfoReader is a Reader for the GetCurrentUserInfo structure.
type GetCurrentUserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentUserInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentUserInfoOK creates a GetCurrentUserInfoOK with default headers values
func NewGetCurrentUserInfoOK() *GetCurrentUserInfoOK {
	return &GetCurrentUserInfoOK{}
}

/* GetCurrentUserInfoOK describes a response with status code 200, with default header values.

Get current user information successfully.
*/
type GetCurrentUserInfoOK struct {
	Payload *models.UserResp
}

func (o *GetCurrentUserInfoOK) Error() string {
	return fmt.Sprintf("[GET /users/current][%d] getCurrentUserInfoOK  %+v", 200, o.Payload)
}
func (o *GetCurrentUserInfoOK) GetPayload() *models.UserResp {
	return o.Payload
}

func (o *GetCurrentUserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserInfoUnauthorized creates a GetCurrentUserInfoUnauthorized with default headers values
func NewGetCurrentUserInfoUnauthorized() *GetCurrentUserInfoUnauthorized {
	return &GetCurrentUserInfoUnauthorized{}
}

/* GetCurrentUserInfoUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCurrentUserInfoUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetCurrentUserInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current][%d] getCurrentUserInfoUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCurrentUserInfoUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetCurrentUserInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCurrentUserInfoInternalServerError creates a GetCurrentUserInfoInternalServerError with default headers values
func NewGetCurrentUserInfoInternalServerError() *GetCurrentUserInfoInternalServerError {
	return &GetCurrentUserInfoInternalServerError{}
}

/* GetCurrentUserInfoInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCurrentUserInfoInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetCurrentUserInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current][%d] getCurrentUserInfoInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCurrentUserInfoInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetCurrentUserInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
