// Code generated by go-swagger; DO NOT EDIT.

package net_access_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateNetAccessPointReader is a Reader for the UpdateNetAccessPoint structure.
type UpdateNetAccessPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetAccessPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetAccessPointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetAccessPointOK creates a UpdateNetAccessPointOK with default headers values
func NewUpdateNetAccessPointOK() *UpdateNetAccessPointOK {
	return &UpdateNetAccessPointOK{}
}

/*UpdateNetAccessPointOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type UpdateNetAccessPointOK struct {
}

func (o *UpdateNetAccessPointOK) Error() string {
	return fmt.Sprintf("[POST /UpdateNetAccessPoint][%d] updateNetAccessPointOK ", 200)
}

func (o *UpdateNetAccessPointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
