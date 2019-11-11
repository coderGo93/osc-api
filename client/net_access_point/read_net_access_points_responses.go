// Code generated by go-swagger; DO NOT EDIT.

package net_access_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ReadNetAccessPointsReader is a Reader for the ReadNetAccessPoints structure.
type ReadNetAccessPointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadNetAccessPointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadNetAccessPointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadNetAccessPointsOK creates a ReadNetAccessPointsOK with default headers values
func NewReadNetAccessPointsOK() *ReadNetAccessPointsOK {
	return &ReadNetAccessPointsOK{}
}

/*ReadNetAccessPointsOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type ReadNetAccessPointsOK struct {
}

func (o *ReadNetAccessPointsOK) Error() string {
	return fmt.Sprintf("[POST /ReadNetAccessPoints][%d] readNetAccessPointsOK ", 200)
}

func (o *ReadNetAccessPointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
