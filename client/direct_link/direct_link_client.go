// Code generated by go-swagger; DO NOT EDIT.

package direct_link

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new direct link API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for direct link API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDirectLink create direct link API
*/
func (a *Client) CreateDirectLink(params *CreateDirectLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDirectLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDirectLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDirectLink",
		Method:             "POST",
		PathPattern:        "/CreateDirectLink",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDirectLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDirectLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDirectLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDirectLink delete direct link API
*/
func (a *Client) DeleteDirectLink(params *DeleteDirectLinkParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDirectLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDirectLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDirectLink",
		Method:             "POST",
		PathPattern:        "/DeleteDirectLink",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDirectLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDirectLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDirectLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadDirectLinks read direct links API
*/
func (a *Client) ReadDirectLinks(params *ReadDirectLinksParams, authInfo runtime.ClientAuthInfoWriter) (*ReadDirectLinksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadDirectLinksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadDirectLinks",
		Method:             "POST",
		PathPattern:        "/ReadDirectLinks",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadDirectLinksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadDirectLinksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadDirectLinks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
