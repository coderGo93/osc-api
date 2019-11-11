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

// NewLinkVirtualGatewayParams creates a new LinkVirtualGatewayParams object
// with the default values initialized.
func NewLinkVirtualGatewayParams() *LinkVirtualGatewayParams {

	return &LinkVirtualGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLinkVirtualGatewayParamsWithTimeout creates a new LinkVirtualGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLinkVirtualGatewayParamsWithTimeout(timeout time.Duration) *LinkVirtualGatewayParams {

	return &LinkVirtualGatewayParams{

		timeout: timeout,
	}
}

// NewLinkVirtualGatewayParamsWithContext creates a new LinkVirtualGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewLinkVirtualGatewayParamsWithContext(ctx context.Context) *LinkVirtualGatewayParams {

	return &LinkVirtualGatewayParams{

		Context: ctx,
	}
}

// NewLinkVirtualGatewayParamsWithHTTPClient creates a new LinkVirtualGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLinkVirtualGatewayParamsWithHTTPClient(client *http.Client) *LinkVirtualGatewayParams {

	return &LinkVirtualGatewayParams{
		HTTPClient: client,
	}
}

/*LinkVirtualGatewayParams contains all the parameters to send to the API endpoint
for the link virtual gateway operation typically these are written to a http.Request
*/
type LinkVirtualGatewayParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the link virtual gateway params
func (o *LinkVirtualGatewayParams) WithTimeout(timeout time.Duration) *LinkVirtualGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the link virtual gateway params
func (o *LinkVirtualGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the link virtual gateway params
func (o *LinkVirtualGatewayParams) WithContext(ctx context.Context) *LinkVirtualGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the link virtual gateway params
func (o *LinkVirtualGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the link virtual gateway params
func (o *LinkVirtualGatewayParams) WithHTTPClient(client *http.Client) *LinkVirtualGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the link virtual gateway params
func (o *LinkVirtualGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LinkVirtualGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
