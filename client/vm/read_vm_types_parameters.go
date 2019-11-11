// Code generated by go-swagger; DO NOT EDIT.

package vm

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

// NewReadVMTypesParams creates a new ReadVMTypesParams object
// with the default values initialized.
func NewReadVMTypesParams() *ReadVMTypesParams {

	return &ReadVMTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadVMTypesParamsWithTimeout creates a new ReadVMTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadVMTypesParamsWithTimeout(timeout time.Duration) *ReadVMTypesParams {

	return &ReadVMTypesParams{

		timeout: timeout,
	}
}

// NewReadVMTypesParamsWithContext creates a new ReadVMTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadVMTypesParamsWithContext(ctx context.Context) *ReadVMTypesParams {

	return &ReadVMTypesParams{

		Context: ctx,
	}
}

// NewReadVMTypesParamsWithHTTPClient creates a new ReadVMTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadVMTypesParamsWithHTTPClient(client *http.Client) *ReadVMTypesParams {

	return &ReadVMTypesParams{
		HTTPClient: client,
	}
}

/*ReadVMTypesParams contains all the parameters to send to the API endpoint
for the read Vm types operation typically these are written to a http.Request
*/
type ReadVMTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read Vm types params
func (o *ReadVMTypesParams) WithTimeout(timeout time.Duration) *ReadVMTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read Vm types params
func (o *ReadVMTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read Vm types params
func (o *ReadVMTypesParams) WithContext(ctx context.Context) *ReadVMTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read Vm types params
func (o *ReadVMTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read Vm types params
func (o *ReadVMTypesParams) WithHTTPClient(client *http.Client) *ReadVMTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read Vm types params
func (o *ReadVMTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadVMTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
