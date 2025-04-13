// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPipelineServiceDeletePipelineV1Params creates a new PipelineServiceDeletePipelineV1Params object
// with the default values initialized.
func NewPipelineServiceDeletePipelineV1Params() *PipelineServiceDeletePipelineV1Params {
	var ()
	return &PipelineServiceDeletePipelineV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPipelineServiceDeletePipelineV1ParamsWithTimeout creates a new PipelineServiceDeletePipelineV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPipelineServiceDeletePipelineV1ParamsWithTimeout(timeout time.Duration) *PipelineServiceDeletePipelineV1Params {
	var ()
	return &PipelineServiceDeletePipelineV1Params{

		timeout: timeout,
	}
}

// NewPipelineServiceDeletePipelineV1ParamsWithContext creates a new PipelineServiceDeletePipelineV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPipelineServiceDeletePipelineV1ParamsWithContext(ctx context.Context) *PipelineServiceDeletePipelineV1Params {
	var ()
	return &PipelineServiceDeletePipelineV1Params{

		Context: ctx,
	}
}

// NewPipelineServiceDeletePipelineV1ParamsWithHTTPClient creates a new PipelineServiceDeletePipelineV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPipelineServiceDeletePipelineV1ParamsWithHTTPClient(client *http.Client) *PipelineServiceDeletePipelineV1Params {
	var ()
	return &PipelineServiceDeletePipelineV1Params{
		HTTPClient: client,
	}
}

/*PipelineServiceDeletePipelineV1Params contains all the parameters to send to the API endpoint
for the pipeline service delete pipeline v1 operation typically these are written to a http.Request
*/
type PipelineServiceDeletePipelineV1Params struct {

	/*Cascade
	  Optional. If true, the pipeline and all its versions will be deleted.
	If false (default), only the pipeline will be deleted if it has no versions.

	*/
	Cascade *bool
	/*ID
	  The ID of the pipeline to be deleted.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) WithTimeout(timeout time.Duration) *PipelineServiceDeletePipelineV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) WithContext(ctx context.Context) *PipelineServiceDeletePipelineV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) WithHTTPClient(client *http.Client) *PipelineServiceDeletePipelineV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCascade adds the cascade to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) WithCascade(cascade *bool) *PipelineServiceDeletePipelineV1Params {
	o.SetCascade(cascade)
	return o
}

// SetCascade adds the cascade to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) SetCascade(cascade *bool) {
	o.Cascade = cascade
}

// WithID adds the id to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) WithID(id string) *PipelineServiceDeletePipelineV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the pipeline service delete pipeline v1 params
func (o *PipelineServiceDeletePipelineV1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PipelineServiceDeletePipelineV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cascade != nil {

		// query param cascade
		var qrCascade bool
		if o.Cascade != nil {
			qrCascade = *o.Cascade
		}
		qCascade := swag.FormatBool(qrCascade)
		if qCascade != "" {
			if err := r.SetQueryParam("cascade", qCascade); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
