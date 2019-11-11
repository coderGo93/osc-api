// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vm API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vm API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateVms create vms API
*/
func (a *Client) CreateVms(params *CreateVmsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateVms",
		Method:             "POST",
		PathPattern:        "/CreateVms",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteVms delete vms API
*/
func (a *Client) DeleteVms(params *DeleteVmsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteVmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVms",
		Method:             "POST",
		PathPattern:        "/DeleteVms",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadAdminPassword read admin password API
*/
func (a *Client) ReadAdminPassword(params *ReadAdminPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*ReadAdminPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadAdminPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadAdminPassword",
		Method:             "POST",
		PathPattern:        "/ReadAdminPassword",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadAdminPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadAdminPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadAdminPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadConsoleOutput read console output API
*/
func (a *Client) ReadConsoleOutput(params *ReadConsoleOutputParams, authInfo runtime.ClientAuthInfoWriter) (*ReadConsoleOutputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadConsoleOutputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadConsoleOutput",
		Method:             "POST",
		PathPattern:        "/ReadConsoleOutput",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadConsoleOutputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadConsoleOutputOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadConsoleOutput: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadVMTypes read Vm types API
*/
func (a *Client) ReadVMTypes(params *ReadVMTypesParams) (*ReadVMTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadVMTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadVmTypes",
		Method:             "POST",
		PathPattern:        "/ReadVmTypes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadVMTypesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadVMTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadVmTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadVms read vms API
*/
func (a *Client) ReadVms(params *ReadVmsParams, authInfo runtime.ClientAuthInfoWriter) (*ReadVmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadVmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadVms",
		Method:             "POST",
		PathPattern:        "/ReadVms",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadVmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadVmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadVms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadVmsState read vms state API
*/
func (a *Client) ReadVmsState(params *ReadVmsStateParams, authInfo runtime.ClientAuthInfoWriter) (*ReadVmsStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadVmsStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadVmsState",
		Method:             "POST",
		PathPattern:        "/ReadVmsState",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadVmsStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadVmsStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReadVmsState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RebootVms reboot vms API
*/
func (a *Client) RebootVms(params *RebootVmsParams, authInfo runtime.ClientAuthInfoWriter) (*RebootVmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootVmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RebootVms",
		Method:             "POST",
		PathPattern:        "/RebootVms",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RebootVmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RebootVmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RebootVms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartVms start vms API
*/
func (a *Client) StartVms(params *StartVmsParams, authInfo runtime.ClientAuthInfoWriter) (*StartVmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartVmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartVms",
		Method:             "POST",
		PathPattern:        "/StartVms",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartVmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartVmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartVms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopVms stop vms API
*/
func (a *Client) StopVms(params *StopVmsParams, authInfo runtime.ClientAuthInfoWriter) (*StopVmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopVmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StopVms",
		Method:             "POST",
		PathPattern:        "/StopVms",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopVmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopVmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopVms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateVM update Vm API
*/
func (a *Client) UpdateVM(params *UpdateVMParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVMParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateVm",
		Method:             "POST",
		PathPattern:        "/UpdateVm",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVMReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateVm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
