// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListCredentialsReader is a Reader for the ListCredentials structure.
type ListCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCredentialsOK creates a ListCredentialsOK with default headers values
func NewListCredentialsOK() *ListCredentialsOK {
	return &ListCredentialsOK{}
}

/*
ListCredentialsOK describes a response with status code 200, with default header values.

CredentialList
*/
type ListCredentialsOK struct {
	Payload *models.CredentialList
}

// IsSuccess returns true when this list credentials o k response has a 2xx status code
func (o *ListCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list credentials o k response has a 3xx status code
func (o *ListCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list credentials o k response has a 4xx status code
func (o *ListCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list credentials o k response has a 5xx status code
func (o *ListCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list credentials o k response a status code equal to that given
func (o *ListCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/presets/credentials][%d] listCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/presets/credentials][%d] listCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListCredentialsOK) GetPayload() *models.CredentialList {
	return o.Payload
}

func (o *ListCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCredentialsDefault creates a ListCredentialsDefault with default headers values
func NewListCredentialsDefault(code int) *ListCredentialsDefault {
	return &ListCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list credentials default response
func (o *ListCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list credentials default response has a 2xx status code
func (o *ListCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list credentials default response has a 3xx status code
func (o *ListCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list credentials default response has a 4xx status code
func (o *ListCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list credentials default response has a 5xx status code
func (o *ListCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list credentials default response a status code equal to that given
func (o *ListCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/presets/credentials][%d] listCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/presets/credentials][%d] listCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
