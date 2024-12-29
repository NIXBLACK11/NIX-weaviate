// Code generated by go-swagger; DO NOT EDIT.

package authz

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

// NewAssignRoleParams creates a new AssignRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignRoleParams() *AssignRoleParams {
	return &AssignRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignRoleParamsWithTimeout creates a new AssignRoleParams object
// with the ability to set a timeout on a request.
func NewAssignRoleParamsWithTimeout(timeout time.Duration) *AssignRoleParams {
	return &AssignRoleParams{
		timeout: timeout,
	}
}

// NewAssignRoleParamsWithContext creates a new AssignRoleParams object
// with the ability to set a context for a request.
func NewAssignRoleParamsWithContext(ctx context.Context) *AssignRoleParams {
	return &AssignRoleParams{
		Context: ctx,
	}
}

// NewAssignRoleParamsWithHTTPClient creates a new AssignRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignRoleParamsWithHTTPClient(client *http.Client) *AssignRoleParams {
	return &AssignRoleParams{
		HTTPClient: client,
	}
}

/*
AssignRoleParams contains all the parameters to send to the API endpoint

	for the assign role operation.

	Typically these are written to a http.Request.
*/
type AssignRoleParams struct {

	// Body.
	Body AssignRoleBody

	/* ID.

	   user name
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignRoleParams) WithDefaults() *AssignRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign role params
func (o *AssignRoleParams) WithTimeout(timeout time.Duration) *AssignRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign role params
func (o *AssignRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign role params
func (o *AssignRoleParams) WithContext(ctx context.Context) *AssignRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign role params
func (o *AssignRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign role params
func (o *AssignRoleParams) WithHTTPClient(client *http.Client) *AssignRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign role params
func (o *AssignRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the assign role params
func (o *AssignRoleParams) WithBody(body AssignRoleBody) *AssignRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the assign role params
func (o *AssignRoleParams) SetBody(body AssignRoleBody) {
	o.Body = body
}

// WithID adds the id to the assign role params
func (o *AssignRoleParams) WithID(id string) *AssignRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign role params
func (o *AssignRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AssignRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
