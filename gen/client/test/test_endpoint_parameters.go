// Code generated by go-swagger; DO NOT EDIT.

package test

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

// NewTestEndpointParams creates a new TestEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestEndpointParams() *TestEndpointParams {
	return &TestEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestEndpointParamsWithTimeout creates a new TestEndpointParams object
// with the ability to set a timeout on a request.
func NewTestEndpointParamsWithTimeout(timeout time.Duration) *TestEndpointParams {
	return &TestEndpointParams{
		timeout: timeout,
	}
}

// NewTestEndpointParamsWithContext creates a new TestEndpointParams object
// with the ability to set a context for a request.
func NewTestEndpointParamsWithContext(ctx context.Context) *TestEndpointParams {
	return &TestEndpointParams{
		Context: ctx,
	}
}

// NewTestEndpointParamsWithHTTPClient creates a new TestEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestEndpointParamsWithHTTPClient(client *http.Client) *TestEndpointParams {
	return &TestEndpointParams{
		HTTPClient: client,
	}
}

/* TestEndpointParams contains all the parameters to send to the API endpoint
   for the test endpoint operation.

   Typically these are written to a http.Request.
*/
type TestEndpointParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestEndpointParams) WithDefaults() *TestEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test endpoint params
func (o *TestEndpointParams) WithTimeout(timeout time.Duration) *TestEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test endpoint params
func (o *TestEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test endpoint params
func (o *TestEndpointParams) WithContext(ctx context.Context) *TestEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test endpoint params
func (o *TestEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test endpoint params
func (o *TestEndpointParams) WithHTTPClient(client *http.Client) *TestEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test endpoint params
func (o *TestEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TestEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
