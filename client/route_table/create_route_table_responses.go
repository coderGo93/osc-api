// Code generated by go-swagger; DO NOT EDIT.

package route_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateRouteTableReader is a Reader for the CreateRouteTable structure.
type CreateRouteTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRouteTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRouteTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRouteTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRouteTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRouteTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRouteTableOK creates a CreateRouteTableOK with default headers values
func NewCreateRouteTableOK() *CreateRouteTableOK {
	return &CreateRouteTableOK{}
}

/*CreateRouteTableOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type CreateRouteTableOK struct {
}

func (o *CreateRouteTableOK) Error() string {
	return fmt.Sprintf("[POST /CreateRouteTable][%d] createRouteTableOK ", 200)
}

func (o *CreateRouteTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRouteTableBadRequest creates a CreateRouteTableBadRequest with default headers values
func NewCreateRouteTableBadRequest() *CreateRouteTableBadRequest {
	return &CreateRouteTableBadRequest{}
}

/*CreateRouteTableBadRequest handles this case with default header values.

The HTTP 400 response (Bad Request).
*/
type CreateRouteTableBadRequest struct {
}

func (o *CreateRouteTableBadRequest) Error() string {
	return fmt.Sprintf("[POST /CreateRouteTable][%d] createRouteTableBadRequest ", 400)
}

func (o *CreateRouteTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRouteTableUnauthorized creates a CreateRouteTableUnauthorized with default headers values
func NewCreateRouteTableUnauthorized() *CreateRouteTableUnauthorized {
	return &CreateRouteTableUnauthorized{}
}

/*CreateRouteTableUnauthorized handles this case with default header values.

The HTTP 401 response (Unauthorized).
*/
type CreateRouteTableUnauthorized struct {
}

func (o *CreateRouteTableUnauthorized) Error() string {
	return fmt.Sprintf("[POST /CreateRouteTable][%d] createRouteTableUnauthorized ", 401)
}

func (o *CreateRouteTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRouteTableInternalServerError creates a CreateRouteTableInternalServerError with default headers values
func NewCreateRouteTableInternalServerError() *CreateRouteTableInternalServerError {
	return &CreateRouteTableInternalServerError{}
}

/*CreateRouteTableInternalServerError handles this case with default header values.

The HTTP 500 response (Internal Server Error).
*/
type CreateRouteTableInternalServerError struct {
}

func (o *CreateRouteTableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /CreateRouteTable][%d] createRouteTableInternalServerError ", 500)
}

func (o *CreateRouteTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
