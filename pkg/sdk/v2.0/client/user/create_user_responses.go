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

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/* CreateUserCreated describes a response with status code 201, with default header values.

Created
*/
type CreateUserCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserCreated ", 201)
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/* CreateUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateUserBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserBadRequest  %+v", 400, o.Payload)
}
func (o *CreateUserBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/* CreateUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateUserUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateUserUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/* CreateUserForbidden describes a response with status code 403, with default header values.

When the  self registration is disabled, non-admin does not have permission to create user.  When self registration is enabled, this API can only be called from UI portal, calling it via script will get a 403 error.
*/
type CreateUserForbidden struct {
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserForbidden ", 403)
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/* CreateUserConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateUserConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserConflict  %+v", 409, o.Payload)
}
func (o *CreateUserConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/* CreateUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateUserInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateUserInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
