// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateImageReader is a Reader for the CreateImage structure.
type CreateImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateImageOK creates a CreateImageOK with default headers values
func NewCreateImageOK() *CreateImageOK {
	return &CreateImageOK{}
}

/*CreateImageOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type CreateImageOK struct {
}

func (o *CreateImageOK) Error() string {
	return fmt.Sprintf("[POST /CreateImage][%d] createImageOK ", 200)
}

func (o *CreateImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateImageBadRequest creates a CreateImageBadRequest with default headers values
func NewCreateImageBadRequest() *CreateImageBadRequest {
	return &CreateImageBadRequest{}
}

/*CreateImageBadRequest handles this case with default header values.

The HTTP 400 response (Bad Request).
*/
type CreateImageBadRequest struct {
}

func (o *CreateImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /CreateImage][%d] createImageBadRequest ", 400)
}

func (o *CreateImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateImageUnauthorized creates a CreateImageUnauthorized with default headers values
func NewCreateImageUnauthorized() *CreateImageUnauthorized {
	return &CreateImageUnauthorized{}
}

/*CreateImageUnauthorized handles this case with default header values.

The HTTP 401 response (Unauthorized).
*/
type CreateImageUnauthorized struct {
}

func (o *CreateImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /CreateImage][%d] createImageUnauthorized ", 401)
}

func (o *CreateImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateImageInternalServerError creates a CreateImageInternalServerError with default headers values
func NewCreateImageInternalServerError() *CreateImageInternalServerError {
	return &CreateImageInternalServerError{}
}

/*CreateImageInternalServerError handles this case with default header values.

The HTTP 500 response (Internal Server Error).
*/
type CreateImageInternalServerError struct {
}

func (o *CreateImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /CreateImage][%d] createImageInternalServerError ", 500)
}

func (o *CreateImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
