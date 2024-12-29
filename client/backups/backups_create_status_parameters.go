// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// NewBackupsCreateStatusParams creates a new BackupsCreateStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupsCreateStatusParams() *BackupsCreateStatusParams {
	return &BackupsCreateStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupsCreateStatusParamsWithTimeout creates a new BackupsCreateStatusParams object
// with the ability to set a timeout on a request.
func NewBackupsCreateStatusParamsWithTimeout(timeout time.Duration) *BackupsCreateStatusParams {
	return &BackupsCreateStatusParams{
		timeout: timeout,
	}
}

// NewBackupsCreateStatusParamsWithContext creates a new BackupsCreateStatusParams object
// with the ability to set a context for a request.
func NewBackupsCreateStatusParamsWithContext(ctx context.Context) *BackupsCreateStatusParams {
	return &BackupsCreateStatusParams{
		Context: ctx,
	}
}

// NewBackupsCreateStatusParamsWithHTTPClient creates a new BackupsCreateStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupsCreateStatusParamsWithHTTPClient(client *http.Client) *BackupsCreateStatusParams {
	return &BackupsCreateStatusParams{
		HTTPClient: client,
	}
}

/*
BackupsCreateStatusParams contains all the parameters to send to the API endpoint

	for the backups create status operation.

	Typically these are written to a http.Request.
*/
type BackupsCreateStatusParams struct {

	/* Backend.

	   Backup backend name e.g. filesystem, gcs, s3.
	*/
	Backend string

	/* Bucket.

	   Name of the bucket, container, volume, etc
	*/
	Bucket *string

	/* ID.

	   The ID of a backup. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.
	*/
	ID string

	/* Path.

	   The path within the bucket
	*/
	Path *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backups create status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupsCreateStatusParams) WithDefaults() *BackupsCreateStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backups create status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupsCreateStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backups create status params
func (o *BackupsCreateStatusParams) WithTimeout(timeout time.Duration) *BackupsCreateStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backups create status params
func (o *BackupsCreateStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backups create status params
func (o *BackupsCreateStatusParams) WithContext(ctx context.Context) *BackupsCreateStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backups create status params
func (o *BackupsCreateStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backups create status params
func (o *BackupsCreateStatusParams) WithHTTPClient(client *http.Client) *BackupsCreateStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backups create status params
func (o *BackupsCreateStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the backups create status params
func (o *BackupsCreateStatusParams) WithBackend(backend string) *BackupsCreateStatusParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the backups create status params
func (o *BackupsCreateStatusParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithBucket adds the bucket to the backups create status params
func (o *BackupsCreateStatusParams) WithBucket(bucket *string) *BackupsCreateStatusParams {
	o.SetBucket(bucket)
	return o
}

// SetBucket adds the bucket to the backups create status params
func (o *BackupsCreateStatusParams) SetBucket(bucket *string) {
	o.Bucket = bucket
}

// WithID adds the id to the backups create status params
func (o *BackupsCreateStatusParams) WithID(id string) *BackupsCreateStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the backups create status params
func (o *BackupsCreateStatusParams) SetID(id string) {
	o.ID = id
}

// WithPath adds the path to the backups create status params
func (o *BackupsCreateStatusParams) WithPath(path *string) *BackupsCreateStatusParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the backups create status params
func (o *BackupsCreateStatusParams) SetPath(path *string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *BackupsCreateStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backend
	if err := r.SetPathParam("backend", o.Backend); err != nil {
		return err
	}

	if o.Bucket != nil {

		// query param bucket
		var qrBucket string

		if o.Bucket != nil {
			qrBucket = *o.Bucket
		}
		qBucket := qrBucket
		if qBucket != "" {

			if err := r.SetQueryParam("bucket", qBucket); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
