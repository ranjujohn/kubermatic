// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// NewCreateClusterRoleParams creates a new CreateClusterRoleParams object
// with the default values initialized.
func NewCreateClusterRoleParams() *CreateClusterRoleParams {
	var ()
	return &CreateClusterRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterRoleParamsWithTimeout creates a new CreateClusterRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterRoleParamsWithTimeout(timeout time.Duration) *CreateClusterRoleParams {
	var ()
	return &CreateClusterRoleParams{

		timeout: timeout,
	}
}

// NewCreateClusterRoleParamsWithContext creates a new CreateClusterRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterRoleParamsWithContext(ctx context.Context) *CreateClusterRoleParams {
	var ()
	return &CreateClusterRoleParams{

		Context: ctx,
	}
}

// NewCreateClusterRoleParamsWithHTTPClient creates a new CreateClusterRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterRoleParamsWithHTTPClient(client *http.Client) *CreateClusterRoleParams {
	var ()
	return &CreateClusterRoleParams{
		HTTPClient: client,
	}
}

/*CreateClusterRoleParams contains all the parameters to send to the API endpoint
for the create cluster role operation typically these are written to a http.Request
*/
type CreateClusterRoleParams struct {

	/*Body*/
	Body *models.UserClusterRole
	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster role params
func (o *CreateClusterRoleParams) WithTimeout(timeout time.Duration) *CreateClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster role params
func (o *CreateClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster role params
func (o *CreateClusterRoleParams) WithContext(ctx context.Context) *CreateClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster role params
func (o *CreateClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster role params
func (o *CreateClusterRoleParams) WithHTTPClient(client *http.Client) *CreateClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster role params
func (o *CreateClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster role params
func (o *CreateClusterRoleParams) WithBody(body *models.UserClusterRole) *CreateClusterRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster role params
func (o *CreateClusterRoleParams) SetBody(body *models.UserClusterRole) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create cluster role params
func (o *CreateClusterRoleParams) WithClusterID(clusterID string) *CreateClusterRoleParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create cluster role params
func (o *CreateClusterRoleParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the create cluster role params
func (o *CreateClusterRoleParams) WithDc(dc string) *CreateClusterRoleParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the create cluster role params
func (o *CreateClusterRoleParams) SetDc(dc string) {
	o.Dc = dc
}

// WithProjectID adds the projectID to the create cluster role params
func (o *CreateClusterRoleParams) WithProjectID(projectID string) *CreateClusterRoleParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create cluster role params
func (o *CreateClusterRoleParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
