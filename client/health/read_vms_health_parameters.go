// Code generated by go-swagger; DO NOT EDIT.

package health

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

// NewReadVmsHealthParams creates a new ReadVmsHealthParams object
// with the default values initialized.
func NewReadVmsHealthParams() *ReadVmsHealthParams {

	return &ReadVmsHealthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadVmsHealthParamsWithTimeout creates a new ReadVmsHealthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadVmsHealthParamsWithTimeout(timeout time.Duration) *ReadVmsHealthParams {

	return &ReadVmsHealthParams{

		timeout: timeout,
	}
}

// NewReadVmsHealthParamsWithContext creates a new ReadVmsHealthParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadVmsHealthParamsWithContext(ctx context.Context) *ReadVmsHealthParams {

	return &ReadVmsHealthParams{

		Context: ctx,
	}
}

// NewReadVmsHealthParamsWithHTTPClient creates a new ReadVmsHealthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadVmsHealthParamsWithHTTPClient(client *http.Client) *ReadVmsHealthParams {

	return &ReadVmsHealthParams{
		HTTPClient: client,
	}
}

/*ReadVmsHealthParams contains all the parameters to send to the API endpoint
for the read vms health operation typically these are written to a http.Request
*/
type ReadVmsHealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read vms health params
func (o *ReadVmsHealthParams) WithTimeout(timeout time.Duration) *ReadVmsHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read vms health params
func (o *ReadVmsHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read vms health params
func (o *ReadVmsHealthParams) WithContext(ctx context.Context) *ReadVmsHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read vms health params
func (o *ReadVmsHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read vms health params
func (o *ReadVmsHealthParams) WithHTTPClient(client *http.Client) *ReadVmsHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read vms health params
func (o *ReadVmsHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadVmsHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
