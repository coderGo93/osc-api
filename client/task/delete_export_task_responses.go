// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteExportTaskReader is a Reader for the DeleteExportTask structure.
type DeleteExportTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExportTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExportTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExportTaskOK creates a DeleteExportTaskOK with default headers values
func NewDeleteExportTaskOK() *DeleteExportTaskOK {
	return &DeleteExportTaskOK{}
}

/*DeleteExportTaskOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type DeleteExportTaskOK struct {
}

func (o *DeleteExportTaskOK) Error() string {
	return fmt.Sprintf("[POST /DeleteExportTask][%d] deleteExportTaskOK ", 200)
}

func (o *DeleteExportTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
