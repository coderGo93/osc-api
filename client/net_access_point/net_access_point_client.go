// Code generated by go-swagger; DO NOT EDIT.

package net_access_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new net access point API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for net access point API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateNetAccessPoint create net access point API
*/
func (a *Client) CreateNetAccessPoint(params *CreateNetAccessPointParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetAccessPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetAccessPointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateNetAccessPoint",
		Method:             "POST",
		PathPattern:        "/CreateNetAccessPoint",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNetAccessPointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateNetAccessPointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateNetAccessPoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteNetAccessPoint delete net access point API
*/
func (a *Client) DeleteNetAccessPoint(params *DeleteNetAccessPointParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetAccessPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetAccessPointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetAccessPoint",
		Method:             "POST",
		PathPattern:        "/DeleteNetAccessPoint",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNetAccessPointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetAccessPointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNetAccessPoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadNetAccessPointServices read net access point services API
*/
func (a *Client) ReadNetAccessPointServices(params *ReadNetAccessPointServicesParams) (*ReadNetAccessPointServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadNetAccessPointServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadNetAccessPointServices",
		Method:             "POST",
		PathPattern:        "/ReadNetAccessPointServices",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadNetAccessPointServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadNetAccessPointServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadNetAccessPointServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadNetAccessPoints read net access points API
*/
func (a *Client) ReadNetAccessPoints(params *ReadNetAccessPointsParams, authInfo runtime.ClientAuthInfoWriter) (*ReadNetAccessPointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadNetAccessPointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadNetAccessPoints",
		Method:             "POST",
		PathPattern:        "/ReadNetAccessPoints",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadNetAccessPointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadNetAccessPointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadNetAccessPoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateNetAccessPoint update net access point API
*/
func (a *Client) UpdateNetAccessPoint(params *UpdateNetAccessPointParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetAccessPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetAccessPointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateNetAccessPoint",
		Method:             "POST",
		PathPattern:        "/UpdateNetAccessPoint",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateNetAccessPointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetAccessPointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateNetAccessPoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
