// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ReadQuotasReader is a Reader for the ReadQuotas structure.
type ReadQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadQuotasOK creates a ReadQuotasOK with default headers values
func NewReadQuotasOK() *ReadQuotasOK {
	return &ReadQuotasOK{}
}

/*ReadQuotasOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type ReadQuotasOK struct {
}

func (o *ReadQuotasOK) Error() string {
	return fmt.Sprintf("[POST /ReadQuotas][%d] readQuotasOK ", 200)
}

func (o *ReadQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
