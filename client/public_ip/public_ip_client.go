// Code generated by go-swagger; DO NOT EDIT.

package public_ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new public ip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public ip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreatePublicIP create public Ip API
*/
func (a *Client) CreatePublicIP(params *CreatePublicIPParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePublicIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePublicIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePublicIp",
		Method:             "POST",
		PathPattern:        "/CreatePublicIp",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePublicIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePublicIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreatePublicIp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePublicIP delete public Ip API
*/
func (a *Client) DeletePublicIP(params *DeletePublicIPParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePublicIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePublicIp",
		Method:             "POST",
		PathPattern:        "/DeletePublicIp",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePublicIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePublicIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePublicIp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LinkPublicIP link public Ip API
*/
func (a *Client) LinkPublicIP(params *LinkPublicIPParams, authInfo runtime.ClientAuthInfoWriter) (*LinkPublicIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLinkPublicIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LinkPublicIp",
		Method:             "POST",
		PathPattern:        "/LinkPublicIp",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LinkPublicIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LinkPublicIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LinkPublicIp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadPublicIps read public ips API
*/
func (a *Client) ReadPublicIps(params *ReadPublicIpsParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPublicIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadPublicIpsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadPublicIps",
		Method:             "POST",
		PathPattern:        "/ReadPublicIps",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadPublicIpsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadPublicIpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadPublicIps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UnlinkPublicIP unlink public Ip API
*/
func (a *Client) UnlinkPublicIP(params *UnlinkPublicIPParams, authInfo runtime.ClientAuthInfoWriter) (*UnlinkPublicIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnlinkPublicIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UnlinkPublicIp",
		Method:             "POST",
		PathPattern:        "/UnlinkPublicIp",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UnlinkPublicIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnlinkPublicIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UnlinkPublicIp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
