// Code generated by go-swagger; DO NOT EDIT.

package security_group_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSecurityGroupRuleReader is a Reader for the DeleteSecurityGroupRule structure.
type DeleteSecurityGroupRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityGroupRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSecurityGroupRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSecurityGroupRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSecurityGroupRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSecurityGroupRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSecurityGroupRuleOK creates a DeleteSecurityGroupRuleOK with default headers values
func NewDeleteSecurityGroupRuleOK() *DeleteSecurityGroupRuleOK {
	return &DeleteSecurityGroupRuleOK{}
}

/*DeleteSecurityGroupRuleOK handles this case with default header values.

The HTTP 200 response (OK).
*/
type DeleteSecurityGroupRuleOK struct {
}

func (o *DeleteSecurityGroupRuleOK) Error() string {
	return fmt.Sprintf("[POST /DeleteSecurityGroupRule][%d] deleteSecurityGroupRuleOK ", 200)
}

func (o *DeleteSecurityGroupRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityGroupRuleBadRequest creates a DeleteSecurityGroupRuleBadRequest with default headers values
func NewDeleteSecurityGroupRuleBadRequest() *DeleteSecurityGroupRuleBadRequest {
	return &DeleteSecurityGroupRuleBadRequest{}
}

/*DeleteSecurityGroupRuleBadRequest handles this case with default header values.

The HTTP 400 response (Bad Request).
*/
type DeleteSecurityGroupRuleBadRequest struct {
}

func (o *DeleteSecurityGroupRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /DeleteSecurityGroupRule][%d] deleteSecurityGroupRuleBadRequest ", 400)
}

func (o *DeleteSecurityGroupRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityGroupRuleUnauthorized creates a DeleteSecurityGroupRuleUnauthorized with default headers values
func NewDeleteSecurityGroupRuleUnauthorized() *DeleteSecurityGroupRuleUnauthorized {
	return &DeleteSecurityGroupRuleUnauthorized{}
}

/*DeleteSecurityGroupRuleUnauthorized handles this case with default header values.

The HTTP 401 response (Unauthorized).
*/
type DeleteSecurityGroupRuleUnauthorized struct {
}

func (o *DeleteSecurityGroupRuleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /DeleteSecurityGroupRule][%d] deleteSecurityGroupRuleUnauthorized ", 401)
}

func (o *DeleteSecurityGroupRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityGroupRuleInternalServerError creates a DeleteSecurityGroupRuleInternalServerError with default headers values
func NewDeleteSecurityGroupRuleInternalServerError() *DeleteSecurityGroupRuleInternalServerError {
	return &DeleteSecurityGroupRuleInternalServerError{}
}

/*DeleteSecurityGroupRuleInternalServerError handles this case with default header values.

The HTTP 500 response (Internal Server Error).
*/
type DeleteSecurityGroupRuleInternalServerError struct {
}

func (o *DeleteSecurityGroupRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /DeleteSecurityGroupRule][%d] deleteSecurityGroupRuleInternalServerError ", 500)
}

func (o *DeleteSecurityGroupRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
