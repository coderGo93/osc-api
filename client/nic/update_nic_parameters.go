// Code generated by go-swagger; DO NOT EDIT.

package nic

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

// NewUpdateNicParams creates a new UpdateNicParams object
// with the default values initialized.
func NewUpdateNicParams() *UpdateNicParams {

	return &UpdateNicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNicParamsWithTimeout creates a new UpdateNicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNicParamsWithTimeout(timeout time.Duration) *UpdateNicParams {

	return &UpdateNicParams{

		timeout: timeout,
	}
}

// NewUpdateNicParamsWithContext creates a new UpdateNicParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNicParamsWithContext(ctx context.Context) *UpdateNicParams {

	return &UpdateNicParams{

		Context: ctx,
	}
}

// NewUpdateNicParamsWithHTTPClient creates a new UpdateNicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNicParamsWithHTTPClient(client *http.Client) *UpdateNicParams {

	return &UpdateNicParams{
		HTTPClient: client,
	}
}

/*UpdateNicParams contains all the parameters to send to the API endpoint
for the update nic operation typically these are written to a http.Request
*/
type UpdateNicParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update nic params
func (o *UpdateNicParams) WithTimeout(timeout time.Duration) *UpdateNicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nic params
func (o *UpdateNicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nic params
func (o *UpdateNicParams) WithContext(ctx context.Context) *UpdateNicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nic params
func (o *UpdateNicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nic params
func (o *UpdateNicParams) WithHTTPClient(client *http.Client) *UpdateNicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nic params
func (o *UpdateNicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
