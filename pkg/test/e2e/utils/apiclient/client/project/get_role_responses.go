// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/*
GetRoleOK describes a response with status code 200, with default header values.

Role
*/
type GetRoleOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this get role o k response has a 2xx status code
func (o *GetRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get role o k response has a 3xx status code
func (o *GetRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role o k response has a 4xx status code
func (o *GetRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get role o k response has a 5xx status code
func (o *GetRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get role o k response a status code equal to that given
func (o *GetRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoleOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRoleOK  %+v", 200, o.Payload)
}

func (o *GetRoleOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRoleOK  %+v", 200, o.Payload)
}

func (o *GetRoleOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleUnauthorized creates a GetRoleUnauthorized with default headers values
func NewGetRoleUnauthorized() *GetRoleUnauthorized {
	return &GetRoleUnauthorized{}
}

/*
GetRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetRoleUnauthorized struct {
}

// IsSuccess returns true when this get role unauthorized response has a 2xx status code
func (o *GetRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role unauthorized response has a 3xx status code
func (o *GetRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role unauthorized response has a 4xx status code
func (o *GetRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role unauthorized response has a 5xx status code
func (o *GetRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get role unauthorized response a status code equal to that given
func (o *GetRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRoleUnauthorized ", 401)
}

func (o *GetRoleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRoleUnauthorized ", 401)
}

func (o *GetRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleForbidden creates a GetRoleForbidden with default headers values
func NewGetRoleForbidden() *GetRoleForbidden {
	return &GetRoleForbidden{}
}

/*
GetRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetRoleForbidden struct {
}

// IsSuccess returns true when this get role forbidden response has a 2xx status code
func (o *GetRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get role forbidden response has a 3xx status code
func (o *GetRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get role forbidden response has a 4xx status code
func (o *GetRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get role forbidden response has a 5xx status code
func (o *GetRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get role forbidden response a status code equal to that given
func (o *GetRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRoleForbidden ", 403)
}

func (o *GetRoleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRoleForbidden ", 403)
}

func (o *GetRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleDefault creates a GetRoleDefault with default headers values
func NewGetRoleDefault(code int) *GetRoleDefault {
	return &GetRoleDefault{
		_statusCode: code,
	}
}

/*
GetRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get role default response
func (o *GetRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get role default response has a 2xx status code
func (o *GetRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get role default response has a 3xx status code
func (o *GetRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get role default response has a 4xx status code
func (o *GetRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get role default response has a 5xx status code
func (o *GetRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get role default response a status code equal to that given
func (o *GetRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetRoleDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}][%d] getRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
