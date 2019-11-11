// Code generated by go-swagger; DO NOT EDIT.

package virtual_gateway

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

// NewReadVirtualGatewaysParams creates a new ReadVirtualGatewaysParams object
// with the default values initialized.
func NewReadVirtualGatewaysParams() *ReadVirtualGatewaysParams {

	return &ReadVirtualGatewaysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadVirtualGatewaysParamsWithTimeout creates a new ReadVirtualGatewaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadVirtualGatewaysParamsWithTimeout(timeout time.Duration) *ReadVirtualGatewaysParams {

	return &ReadVirtualGatewaysParams{

		timeout: timeout,
	}
}

// NewReadVirtualGatewaysParamsWithContext creates a new ReadVirtualGatewaysParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadVirtualGatewaysParamsWithContext(ctx context.Context) *ReadVirtualGatewaysParams {

	return &ReadVirtualGatewaysParams{

		Context: ctx,
	}
}

// NewReadVirtualGatewaysParamsWithHTTPClient creates a new ReadVirtualGatewaysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadVirtualGatewaysParamsWithHTTPClient(client *http.Client) *ReadVirtualGatewaysParams {

	return &ReadVirtualGatewaysParams{
		HTTPClient: client,
	}
}

/*ReadVirtualGatewaysParams contains all the parameters to send to the API endpoint
for the read virtual gateways operation typically these are written to a http.Request
*/
type ReadVirtualGatewaysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read virtual gateways params
func (o *ReadVirtualGatewaysParams) WithTimeout(timeout time.Duration) *ReadVirtualGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read virtual gateways params
func (o *ReadVirtualGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read virtual gateways params
func (o *ReadVirtualGatewaysParams) WithContext(ctx context.Context) *ReadVirtualGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read virtual gateways params
func (o *ReadVirtualGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read virtual gateways params
func (o *ReadVirtualGatewaysParams) WithHTTPClient(client *http.Client) *ReadVirtualGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read virtual gateways params
func (o *ReadVirtualGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadVirtualGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
