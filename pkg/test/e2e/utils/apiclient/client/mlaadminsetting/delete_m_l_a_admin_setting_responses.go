// Code generated by go-swagger; DO NOT EDIT.

package mlaadminsetting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteMLAAdminSettingReader is a Reader for the DeleteMLAAdminSetting structure.
type DeleteMLAAdminSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMLAAdminSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMLAAdminSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMLAAdminSettingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMLAAdminSettingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMLAAdminSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMLAAdminSettingOK creates a DeleteMLAAdminSettingOK with default headers values
func NewDeleteMLAAdminSettingOK() *DeleteMLAAdminSettingOK {
	return &DeleteMLAAdminSettingOK{}
}

/*
DeleteMLAAdminSettingOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteMLAAdminSettingOK struct {
}

// IsSuccess returns true when this delete m l a admin setting o k response has a 2xx status code
func (o *DeleteMLAAdminSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete m l a admin setting o k response has a 3xx status code
func (o *DeleteMLAAdminSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete m l a admin setting o k response has a 4xx status code
func (o *DeleteMLAAdminSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete m l a admin setting o k response has a 5xx status code
func (o *DeleteMLAAdminSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete m l a admin setting o k response a status code equal to that given
func (o *DeleteMLAAdminSettingOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteMLAAdminSettingOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSettingOK ", 200)
}

func (o *DeleteMLAAdminSettingOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSettingOK ", 200)
}

func (o *DeleteMLAAdminSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMLAAdminSettingUnauthorized creates a DeleteMLAAdminSettingUnauthorized with default headers values
func NewDeleteMLAAdminSettingUnauthorized() *DeleteMLAAdminSettingUnauthorized {
	return &DeleteMLAAdminSettingUnauthorized{}
}

/*
DeleteMLAAdminSettingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteMLAAdminSettingUnauthorized struct {
}

// IsSuccess returns true when this delete m l a admin setting unauthorized response has a 2xx status code
func (o *DeleteMLAAdminSettingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete m l a admin setting unauthorized response has a 3xx status code
func (o *DeleteMLAAdminSettingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete m l a admin setting unauthorized response has a 4xx status code
func (o *DeleteMLAAdminSettingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete m l a admin setting unauthorized response has a 5xx status code
func (o *DeleteMLAAdminSettingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete m l a admin setting unauthorized response a status code equal to that given
func (o *DeleteMLAAdminSettingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteMLAAdminSettingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSettingUnauthorized ", 401)
}

func (o *DeleteMLAAdminSettingUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSettingUnauthorized ", 401)
}

func (o *DeleteMLAAdminSettingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMLAAdminSettingForbidden creates a DeleteMLAAdminSettingForbidden with default headers values
func NewDeleteMLAAdminSettingForbidden() *DeleteMLAAdminSettingForbidden {
	return &DeleteMLAAdminSettingForbidden{}
}

/*
DeleteMLAAdminSettingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteMLAAdminSettingForbidden struct {
}

// IsSuccess returns true when this delete m l a admin setting forbidden response has a 2xx status code
func (o *DeleteMLAAdminSettingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete m l a admin setting forbidden response has a 3xx status code
func (o *DeleteMLAAdminSettingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete m l a admin setting forbidden response has a 4xx status code
func (o *DeleteMLAAdminSettingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete m l a admin setting forbidden response has a 5xx status code
func (o *DeleteMLAAdminSettingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete m l a admin setting forbidden response a status code equal to that given
func (o *DeleteMLAAdminSettingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteMLAAdminSettingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSettingForbidden ", 403)
}

func (o *DeleteMLAAdminSettingForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSettingForbidden ", 403)
}

func (o *DeleteMLAAdminSettingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMLAAdminSettingDefault creates a DeleteMLAAdminSettingDefault with default headers values
func NewDeleteMLAAdminSettingDefault(code int) *DeleteMLAAdminSettingDefault {
	return &DeleteMLAAdminSettingDefault{
		_statusCode: code,
	}
}

/*
DeleteMLAAdminSettingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteMLAAdminSettingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete m l a admin setting default response
func (o *DeleteMLAAdminSettingDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete m l a admin setting default response has a 2xx status code
func (o *DeleteMLAAdminSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete m l a admin setting default response has a 3xx status code
func (o *DeleteMLAAdminSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete m l a admin setting default response has a 4xx status code
func (o *DeleteMLAAdminSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete m l a admin setting default response has a 5xx status code
func (o *DeleteMLAAdminSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete m l a admin setting default response a status code equal to that given
func (o *DeleteMLAAdminSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteMLAAdminSettingDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSetting default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMLAAdminSettingDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] deleteMLAAdminSetting default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMLAAdminSettingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteMLAAdminSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
