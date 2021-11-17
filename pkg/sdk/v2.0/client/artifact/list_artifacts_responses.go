// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// ListArtifactsReader is a Reader for the ListArtifacts structure.
type ListArtifactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArtifactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArtifactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListArtifactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListArtifactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListArtifactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListArtifactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListArtifactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListArtifactsOK creates a ListArtifactsOK with default headers values
func NewListArtifactsOK() *ListArtifactsOK {
	return &ListArtifactsOK{}
}

/* ListArtifactsOK describes a response with status code 200, with default header values.

Success
*/
type ListArtifactsOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of artifacts
	 */
	XTotalCount int64

	Payload []*models.Artifact
}

func (o *ListArtifactsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts][%d] listArtifactsOK  %+v", 200, o.Payload)
}
func (o *ListArtifactsOK) GetPayload() []*models.Artifact {
	return o.Payload
}

func (o *ListArtifactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactsBadRequest creates a ListArtifactsBadRequest with default headers values
func NewListArtifactsBadRequest() *ListArtifactsBadRequest {
	return &ListArtifactsBadRequest{}
}

/* ListArtifactsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListArtifactsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListArtifactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts][%d] listArtifactsBadRequest  %+v", 400, o.Payload)
}
func (o *ListArtifactsBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListArtifactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListArtifactsUnauthorized creates a ListArtifactsUnauthorized with default headers values
func NewListArtifactsUnauthorized() *ListArtifactsUnauthorized {
	return &ListArtifactsUnauthorized{}
}

/* ListArtifactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListArtifactsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListArtifactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts][%d] listArtifactsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListArtifactsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListArtifactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListArtifactsForbidden creates a ListArtifactsForbidden with default headers values
func NewListArtifactsForbidden() *ListArtifactsForbidden {
	return &ListArtifactsForbidden{}
}

/* ListArtifactsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListArtifactsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListArtifactsForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts][%d] listArtifactsForbidden  %+v", 403, o.Payload)
}
func (o *ListArtifactsForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListArtifactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListArtifactsNotFound creates a ListArtifactsNotFound with default headers values
func NewListArtifactsNotFound() *ListArtifactsNotFound {
	return &ListArtifactsNotFound{}
}

/* ListArtifactsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListArtifactsNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListArtifactsNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts][%d] listArtifactsNotFound  %+v", 404, o.Payload)
}
func (o *ListArtifactsNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListArtifactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListArtifactsInternalServerError creates a ListArtifactsInternalServerError with default headers values
func NewListArtifactsInternalServerError() *ListArtifactsInternalServerError {
	return &ListArtifactsInternalServerError{}
}

/* ListArtifactsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListArtifactsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListArtifactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts][%d] listArtifactsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListArtifactsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListArtifactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
