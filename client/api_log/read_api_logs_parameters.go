// Code generated by go-swagger; DO NOT EDIT.

package api_log

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

// NewReadAPILogsParams creates a new ReadAPILogsParams object
// with the default values initialized.
func NewReadAPILogsParams() *ReadAPILogsParams {

	return &ReadAPILogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadAPILogsParamsWithTimeout creates a new ReadAPILogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadAPILogsParamsWithTimeout(timeout time.Duration) *ReadAPILogsParams {

	return &ReadAPILogsParams{

		timeout: timeout,
	}
}

// NewReadAPILogsParamsWithContext creates a new ReadAPILogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadAPILogsParamsWithContext(ctx context.Context) *ReadAPILogsParams {

	return &ReadAPILogsParams{

		Context: ctx,
	}
}

// NewReadAPILogsParamsWithHTTPClient creates a new ReadAPILogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadAPILogsParamsWithHTTPClient(client *http.Client) *ReadAPILogsParams {

	return &ReadAPILogsParams{
		HTTPClient: client,
	}
}

/*ReadAPILogsParams contains all the parameters to send to the API endpoint
for the read Api logs operation typically these are written to a http.Request
*/
type ReadAPILogsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read Api logs params
func (o *ReadAPILogsParams) WithTimeout(timeout time.Duration) *ReadAPILogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read Api logs params
func (o *ReadAPILogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read Api logs params
func (o *ReadAPILogsParams) WithContext(ctx context.Context) *ReadAPILogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read Api logs params
func (o *ReadAPILogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read Api logs params
func (o *ReadAPILogsParams) WithHTTPClient(client *http.Client) *ReadAPILogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read Api logs params
func (o *ReadAPILogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAPILogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
