// Code generated by go-swagger; DO NOT EDIT.

package nic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteNicReader is a Reader for the DeleteNic structure.
type DeleteNicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNicUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNicOK creates a DeleteNicOK with default headers values
func NewDeleteNicOK() *DeleteNicOK {
	return &DeleteNicOK{}
}

/*DeleteNicOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type DeleteNicOK struct {
}

func (o *DeleteNicOK) Error() string {
	return fmt.Sprintf("[POST /DeleteNic][%d] deleteNicOK ", 200)
}

func (o *DeleteNicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNicBadRequest creates a DeleteNicBadRequest with default headers values
func NewDeleteNicBadRequest() *DeleteNicBadRequest {
	return &DeleteNicBadRequest{}
}

/*DeleteNicBadRequest handles this case with default header values.

The HTTP 400 response (Bad Request).
*/
type DeleteNicBadRequest struct {
}

func (o *DeleteNicBadRequest) Error() string {
	return fmt.Sprintf("[POST /DeleteNic][%d] deleteNicBadRequest ", 400)
}

func (o *DeleteNicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNicUnauthorized creates a DeleteNicUnauthorized with default headers values
func NewDeleteNicUnauthorized() *DeleteNicUnauthorized {
	return &DeleteNicUnauthorized{}
}

/*DeleteNicUnauthorized handles this case with default header values.

The HTTP 401 response (Unauthorized).
*/
type DeleteNicUnauthorized struct {
}

func (o *DeleteNicUnauthorized) Error() string {
	return fmt.Sprintf("[POST /DeleteNic][%d] deleteNicUnauthorized ", 401)
}

func (o *DeleteNicUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNicInternalServerError creates a DeleteNicInternalServerError with default headers values
func NewDeleteNicInternalServerError() *DeleteNicInternalServerError {
	return &DeleteNicInternalServerError{}
}

/*DeleteNicInternalServerError handles this case with default header values.

The HTTP 500 response (Internal Server Error).
*/
type DeleteNicInternalServerError struct {
}

func (o *DeleteNicInternalServerError) Error() string {
	return fmt.Sprintf("[POST /DeleteNic][%d] deleteNicInternalServerError ", 500)
}

func (o *DeleteNicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
