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

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodParams creates a new CompleteSelfServiceRecoveryFlowWithLinkMethodParams object
// with the default values initialized.
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodParams() *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodParamsWithTimeout creates a new CompleteSelfServiceRecoveryFlowWithLinkMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodParamsWithTimeout(timeout time.Duration) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodParams{

		timeout: timeout,
	}
}

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodParamsWithContext creates a new CompleteSelfServiceRecoveryFlowWithLinkMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodParamsWithContext(ctx context.Context) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodParams{

		Context: ctx,
	}
}

// NewCompleteSelfServiceRecoveryFlowWithLinkMethodParamsWithHTTPClient creates a new CompleteSelfServiceRecoveryFlowWithLinkMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteSelfServiceRecoveryFlowWithLinkMethodParamsWithHTTPClient(client *http.Client) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	var ()
	return &CompleteSelfServiceRecoveryFlowWithLinkMethodParams{
		HTTPClient: client,
	}
}

/*CompleteSelfServiceRecoveryFlowWithLinkMethodParams contains all the parameters to send to the API endpoint
for the complete self service recovery flow with link method operation typically these are written to a http.Request
*/
type CompleteSelfServiceRecoveryFlowWithLinkMethodParams struct {

	/*Body*/
	Body *models.CompleteSelfServiceRecoveryFlowWithLinkMethod
	/*Flow
	  The Flow ID

	format: uuid

	*/
	Flow *string
	/*Token
	  Recovery Token

	The recovery token which completes the recovery request. If the token
	is invalid (e.g. expired) an error will be shown to the end-user.

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WithTimeout(timeout time.Duration) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WithContext(ctx context.Context) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WithHTTPClient(client *http.Client) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WithBody(body *models.CompleteSelfServiceRecoveryFlowWithLinkMethod) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) SetBody(body *models.CompleteSelfServiceRecoveryFlowWithLinkMethod) {
	o.Body = body
}

// WithFlow adds the flow to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WithFlow(flow *string) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	o.SetFlow(flow)
	return o
}

// SetFlow adds the flow to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) SetFlow(flow *string) {
	o.Flow = flow
}

// WithToken adds the token to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WithToken(token *string) *CompleteSelfServiceRecoveryFlowWithLinkMethodParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the complete self service recovery flow with link method params
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteSelfServiceRecoveryFlowWithLinkMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
