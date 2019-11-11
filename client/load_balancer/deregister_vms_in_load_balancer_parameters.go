// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

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

// NewDeregisterVmsInLoadBalancerParams creates a new DeregisterVmsInLoadBalancerParams object
// with the default values initialized.
func NewDeregisterVmsInLoadBalancerParams() *DeregisterVmsInLoadBalancerParams {

	return &DeregisterVmsInLoadBalancerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeregisterVmsInLoadBalancerParamsWithTimeout creates a new DeregisterVmsInLoadBalancerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeregisterVmsInLoadBalancerParamsWithTimeout(timeout time.Duration) *DeregisterVmsInLoadBalancerParams {

	return &DeregisterVmsInLoadBalancerParams{

		timeout: timeout,
	}
}

// NewDeregisterVmsInLoadBalancerParamsWithContext creates a new DeregisterVmsInLoadBalancerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeregisterVmsInLoadBalancerParamsWithContext(ctx context.Context) *DeregisterVmsInLoadBalancerParams {

	return &DeregisterVmsInLoadBalancerParams{

		Context: ctx,
	}
}

// NewDeregisterVmsInLoadBalancerParamsWithHTTPClient creates a new DeregisterVmsInLoadBalancerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeregisterVmsInLoadBalancerParamsWithHTTPClient(client *http.Client) *DeregisterVmsInLoadBalancerParams {

	return &DeregisterVmsInLoadBalancerParams{
		HTTPClient: client,
	}
}

/*DeregisterVmsInLoadBalancerParams contains all the parameters to send to the API endpoint
for the deregister vms in load balancer operation typically these are written to a http.Request
*/
type DeregisterVmsInLoadBalancerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deregister vms in load balancer params
func (o *DeregisterVmsInLoadBalancerParams) WithTimeout(timeout time.Duration) *DeregisterVmsInLoadBalancerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deregister vms in load balancer params
func (o *DeregisterVmsInLoadBalancerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deregister vms in load balancer params
func (o *DeregisterVmsInLoadBalancerParams) WithContext(ctx context.Context) *DeregisterVmsInLoadBalancerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deregister vms in load balancer params
func (o *DeregisterVmsInLoadBalancerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deregister vms in load balancer params
func (o *DeregisterVmsInLoadBalancerParams) WithHTTPClient(client *http.Client) *DeregisterVmsInLoadBalancerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deregister vms in load balancer params
func (o *DeregisterVmsInLoadBalancerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeregisterVmsInLoadBalancerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
