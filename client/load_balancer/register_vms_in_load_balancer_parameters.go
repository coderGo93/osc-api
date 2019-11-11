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

// NewRegisterVmsInLoadBalancerParams creates a new RegisterVmsInLoadBalancerParams object
// with the default values initialized.
func NewRegisterVmsInLoadBalancerParams() *RegisterVmsInLoadBalancerParams {

	return &RegisterVmsInLoadBalancerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterVmsInLoadBalancerParamsWithTimeout creates a new RegisterVmsInLoadBalancerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterVmsInLoadBalancerParamsWithTimeout(timeout time.Duration) *RegisterVmsInLoadBalancerParams {

	return &RegisterVmsInLoadBalancerParams{

		timeout: timeout,
	}
}

// NewRegisterVmsInLoadBalancerParamsWithContext creates a new RegisterVmsInLoadBalancerParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterVmsInLoadBalancerParamsWithContext(ctx context.Context) *RegisterVmsInLoadBalancerParams {

	return &RegisterVmsInLoadBalancerParams{

		Context: ctx,
	}
}

// NewRegisterVmsInLoadBalancerParamsWithHTTPClient creates a new RegisterVmsInLoadBalancerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterVmsInLoadBalancerParamsWithHTTPClient(client *http.Client) *RegisterVmsInLoadBalancerParams {

	return &RegisterVmsInLoadBalancerParams{
		HTTPClient: client,
	}
}

/*RegisterVmsInLoadBalancerParams contains all the parameters to send to the API endpoint
for the register vms in load balancer operation typically these are written to a http.Request
*/
type RegisterVmsInLoadBalancerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register vms in load balancer params
func (o *RegisterVmsInLoadBalancerParams) WithTimeout(timeout time.Duration) *RegisterVmsInLoadBalancerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register vms in load balancer params
func (o *RegisterVmsInLoadBalancerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register vms in load balancer params
func (o *RegisterVmsInLoadBalancerParams) WithContext(ctx context.Context) *RegisterVmsInLoadBalancerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register vms in load balancer params
func (o *RegisterVmsInLoadBalancerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register vms in load balancer params
func (o *RegisterVmsInLoadBalancerParams) WithHTTPClient(client *http.Client) *RegisterVmsInLoadBalancerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register vms in load balancer params
func (o *RegisterVmsInLoadBalancerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterVmsInLoadBalancerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
