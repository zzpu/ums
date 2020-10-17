// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/zzpu/openuser/internal/httpclient/models"
)

// NewCompleteSelfServiceVerificationFlowWithLinkMethodParams creates a new CompleteSelfServiceVerificationFlowWithLinkMethodParams object
// with the default values initialized.
func NewCompleteSelfServiceVerificationFlowWithLinkMethodParams() *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceVerificationFlowWithLinkMethodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteSelfServiceVerificationFlowWithLinkMethodParamsWithTimeout creates a new CompleteSelfServiceVerificationFlowWithLinkMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteSelfServiceVerificationFlowWithLinkMethodParamsWithTimeout(timeout time.Duration) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceVerificationFlowWithLinkMethodParams{

		timeout: timeout,
	}
}

// NewCompleteSelfServiceVerificationFlowWithLinkMethodParamsWithContext creates a new CompleteSelfServiceVerificationFlowWithLinkMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteSelfServiceVerificationFlowWithLinkMethodParamsWithContext(ctx context.Context) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceVerificationFlowWithLinkMethodParams{

		Context: ctx,
	}
}

// NewCompleteSelfServiceVerificationFlowWithLinkMethodParamsWithHTTPClient creates a new CompleteSelfServiceVerificationFlowWithLinkMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteSelfServiceVerificationFlowWithLinkMethodParamsWithHTTPClient(client *http.Client) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceVerificationFlowWithLinkMethodParams{
		HTTPClient: client,
	}
}

/*CompleteSelfServiceVerificationFlowWithLinkMethodParams contains all the parameters to send to the API endpoint
for the complete self service verification flow with link method operation typically these are written to a http.Request
*/
type CompleteSelfServiceVerificationFlowWithLinkMethodParams struct {

	/*Body*/
	Body *models.CompleteSelfServiceVerificationFlowWithLinkMethod
	/*Flow
	  The Flow ID

	format: uuid

	*/
	Flow *string
	/*Token
	  Verification Token

	The verification token which completes the verification request. If the token
	is invalid (e.g. expired) an error will be shown to the end-user.

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WithTimeout(timeout time.Duration) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WithContext(ctx context.Context) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WithHTTPClient(client *http.Client) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WithBody(body *models.CompleteSelfServiceVerificationFlowWithLinkMethod) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) SetBody(body *models.CompleteSelfServiceVerificationFlowWithLinkMethod) {
	o.Body = body
}

// WithFlow adds the flow to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WithFlow(flow *string) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	o.SetFlow(flow)
	return o
}

// SetFlow adds the flow to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) SetFlow(flow *string) {
	o.Flow = flow
}

// WithToken adds the token to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WithToken(token *string) *CompleteSelfServiceVerificationFlowWithLinkMethodParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the complete self service verification flow with link method params
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteSelfServiceVerificationFlowWithLinkMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Flow != nil {

		// query param flow
		var qrFlow string
		if o.Flow != nil {
			qrFlow = *o.Flow
		}
		qFlow := qrFlow
		if qFlow != "" {
			if err := r.SetQueryParam("flow", qFlow); err != nil {
				return err
			}
		}

	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
