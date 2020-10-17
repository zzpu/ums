// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zzpu/openuser/internal/httpclient/models"
)

// CompleteSelfServiceBrowserSettingsOIDCSettingsFlowReader is a Reader for the CompleteSelfServiceBrowserSettingsOIDCSettingsFlow structure.
type CompleteSelfServiceBrowserSettingsOIDCSettingsFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound creates a CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound with default headers values
func NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound() *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound {
	return &CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound{}
}

/*CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound struct {
}

func (o *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/registration/strategies/oidc/settings/connections][%d] completeSelfServiceBrowserSettingsOIdCSettingsFlowFound ", 302)
}

func (o *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError creates a CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError with default headers values
func NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError() *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError {
	return &CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError{}
}

/*CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/registration/strategies/oidc/settings/connections][%d] completeSelfServiceBrowserSettingsOIdCSettingsFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
