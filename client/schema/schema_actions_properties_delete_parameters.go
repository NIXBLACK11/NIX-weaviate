//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

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
)

// NewSchemaActionsPropertiesDeleteParams creates a new SchemaActionsPropertiesDeleteParams object
// with the default values initialized.
func NewSchemaActionsPropertiesDeleteParams() *SchemaActionsPropertiesDeleteParams {
	var ()
	return &SchemaActionsPropertiesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaActionsPropertiesDeleteParamsWithTimeout creates a new SchemaActionsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaActionsPropertiesDeleteParamsWithTimeout(timeout time.Duration) *SchemaActionsPropertiesDeleteParams {
	var ()
	return &SchemaActionsPropertiesDeleteParams{

		timeout: timeout,
	}
}

// NewSchemaActionsPropertiesDeleteParamsWithContext creates a new SchemaActionsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaActionsPropertiesDeleteParamsWithContext(ctx context.Context) *SchemaActionsPropertiesDeleteParams {
	var ()
	return &SchemaActionsPropertiesDeleteParams{

		Context: ctx,
	}
}

// NewSchemaActionsPropertiesDeleteParamsWithHTTPClient creates a new SchemaActionsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaActionsPropertiesDeleteParamsWithHTTPClient(client *http.Client) *SchemaActionsPropertiesDeleteParams {
	var ()
	return &SchemaActionsPropertiesDeleteParams{
		HTTPClient: client,
	}
}

/*SchemaActionsPropertiesDeleteParams contains all the parameters to send to the API endpoint
for the schema actions properties delete operation typically these are written to a http.Request
*/
type SchemaActionsPropertiesDeleteParams struct {

	/*ClassName*/
	ClassName string
	/*PropertyName*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) WithTimeout(timeout time.Duration) *SchemaActionsPropertiesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) WithContext(ctx context.Context) *SchemaActionsPropertiesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) WithHTTPClient(client *http.Client) *SchemaActionsPropertiesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) WithClassName(className string) *SchemaActionsPropertiesDeleteParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) SetClassName(className string) {
	o.ClassName = className
}

// WithPropertyName adds the propertyName to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) WithPropertyName(propertyName string) *SchemaActionsPropertiesDeleteParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the schema actions properties delete params
func (o *SchemaActionsPropertiesDeleteParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaActionsPropertiesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}