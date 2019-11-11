// Code generated by go-swagger; DO NOT EDIT.

package volume

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

// NewDeleteVolumeParams creates a new DeleteVolumeParams object
// with the default values initialized.
func NewDeleteVolumeParams() *DeleteVolumeParams {

	return &DeleteVolumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVolumeParamsWithTimeout creates a new DeleteVolumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVolumeParamsWithTimeout(timeout time.Duration) *DeleteVolumeParams {

	return &DeleteVolumeParams{

		timeout: timeout,
	}
}

// NewDeleteVolumeParamsWithContext creates a new DeleteVolumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVolumeParamsWithContext(ctx context.Context) *DeleteVolumeParams {

	return &DeleteVolumeParams{

		Context: ctx,
	}
}

// NewDeleteVolumeParamsWithHTTPClient creates a new DeleteVolumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVolumeParamsWithHTTPClient(client *http.Client) *DeleteVolumeParams {

	return &DeleteVolumeParams{
		HTTPClient: client,
	}
}

/*DeleteVolumeParams contains all the parameters to send to the API endpoint
for the delete volume operation typically these are written to a http.Request
*/
type DeleteVolumeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete volume params
func (o *DeleteVolumeParams) WithTimeout(timeout time.Duration) *DeleteVolumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete volume params
func (o *DeleteVolumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete volume params
func (o *DeleteVolumeParams) WithContext(ctx context.Context) *DeleteVolumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete volume params
func (o *DeleteVolumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete volume params
func (o *DeleteVolumeParams) WithHTTPClient(client *http.Client) *DeleteVolumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete volume params
func (o *DeleteVolumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVolumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
