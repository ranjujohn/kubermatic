// Code generated by go-swagger; DO NOT EDIT.

package gke

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListGKEClusterImagesParams creates a new ListGKEClusterImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGKEClusterImagesParams() *ListGKEClusterImagesParams {
	return &ListGKEClusterImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGKEClusterImagesParamsWithTimeout creates a new ListGKEClusterImagesParams object
// with the ability to set a timeout on a request.
func NewListGKEClusterImagesParamsWithTimeout(timeout time.Duration) *ListGKEClusterImagesParams {
	return &ListGKEClusterImagesParams{
		timeout: timeout,
	}
}

// NewListGKEClusterImagesParamsWithContext creates a new ListGKEClusterImagesParams object
// with the ability to set a context for a request.
func NewListGKEClusterImagesParamsWithContext(ctx context.Context) *ListGKEClusterImagesParams {
	return &ListGKEClusterImagesParams{
		Context: ctx,
	}
}

// NewListGKEClusterImagesParamsWithHTTPClient creates a new ListGKEClusterImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGKEClusterImagesParamsWithHTTPClient(client *http.Client) *ListGKEClusterImagesParams {
	return &ListGKEClusterImagesParams{
		HTTPClient: client,
	}
}

/* ListGKEClusterImagesParams contains all the parameters to send to the API endpoint
   for the list g k e cluster images operation.

   Typically these are written to a http.Request.
*/
type ListGKEClusterImagesParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list g k e cluster images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGKEClusterImagesParams) WithDefaults() *ListGKEClusterImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list g k e cluster images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGKEClusterImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) WithTimeout(timeout time.Duration) *ListGKEClusterImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) WithContext(ctx context.Context) *ListGKEClusterImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) WithHTTPClient(client *http.Client) *ListGKEClusterImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) WithClusterID(clusterID string) *ListGKEClusterImagesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) WithProjectID(projectID string) *ListGKEClusterImagesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list g k e cluster images params
func (o *ListGKEClusterImagesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListGKEClusterImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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