// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// RejectLoginRequestReader is a Reader for the RejectLoginRequest structure.
type RejectLoginRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectLoginRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRejectLoginRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRejectLoginRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRejectLoginRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRejectLoginRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRejectLoginRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRejectLoginRequestOK creates a RejectLoginRequestOK with default headers values
func NewRejectLoginRequestOK() *RejectLoginRequestOK {
	return &RejectLoginRequestOK{}
}

/*RejectLoginRequestOK handles this case with default header values.

completedRequest
*/
type RejectLoginRequestOK struct {
	Payload *models.CompletedRequest
}

func (o *RejectLoginRequestOK) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/login/reject][%d] rejectLoginRequestOK  %+v", 200, o.Payload)
}

func (o *RejectLoginRequestOK) GetPayload() *models.CompletedRequest {
	return o.Payload
}

func (o *RejectLoginRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompletedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectLoginRequestBadRequest creates a RejectLoginRequestBadRequest with default headers values
func NewRejectLoginRequestBadRequest() *RejectLoginRequestBadRequest {
	return &RejectLoginRequestBadRequest{}
}

/*RejectLoginRequestBadRequest handles this case with default header values.

genericError
*/
type RejectLoginRequestBadRequest struct {
	Payload *models.GenericError
}

func (o *RejectLoginRequestBadRequest) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/login/reject][%d] rejectLoginRequestBadRequest  %+v", 400, o.Payload)
}

func (o *RejectLoginRequestBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RejectLoginRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectLoginRequestUnauthorized creates a RejectLoginRequestUnauthorized with default headers values
func NewRejectLoginRequestUnauthorized() *RejectLoginRequestUnauthorized {
	return &RejectLoginRequestUnauthorized{}
}

/*RejectLoginRequestUnauthorized handles this case with default header values.

genericError
*/
type RejectLoginRequestUnauthorized struct {
	Payload *models.GenericError
}

func (o *RejectLoginRequestUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/login/reject][%d] rejectLoginRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *RejectLoginRequestUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RejectLoginRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectLoginRequestNotFound creates a RejectLoginRequestNotFound with default headers values
func NewRejectLoginRequestNotFound() *RejectLoginRequestNotFound {
	return &RejectLoginRequestNotFound{}
}

/*RejectLoginRequestNotFound handles this case with default header values.

genericError
*/
type RejectLoginRequestNotFound struct {
	Payload *models.GenericError
}

func (o *RejectLoginRequestNotFound) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/login/reject][%d] rejectLoginRequestNotFound  %+v", 404, o.Payload)
}

func (o *RejectLoginRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RejectLoginRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectLoginRequestInternalServerError creates a RejectLoginRequestInternalServerError with default headers values
func NewRejectLoginRequestInternalServerError() *RejectLoginRequestInternalServerError {
	return &RejectLoginRequestInternalServerError{}
}

/*RejectLoginRequestInternalServerError handles this case with default header values.

genericError
*/
type RejectLoginRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *RejectLoginRequestInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/login/reject][%d] rejectLoginRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *RejectLoginRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *RejectLoginRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
