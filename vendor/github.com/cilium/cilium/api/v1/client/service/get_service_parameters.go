// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

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
)

// NewGetServiceParams creates a new GetServiceParams object
// with the default values initialized.
func NewGetServiceParams() *GetServiceParams {

	return &GetServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceParamsWithTimeout creates a new GetServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceParamsWithTimeout(timeout time.Duration) *GetServiceParams {

	return &GetServiceParams{

		timeout: timeout,
	}
}

// NewGetServiceParamsWithContext creates a new GetServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceParamsWithContext(ctx context.Context) *GetServiceParams {

	return &GetServiceParams{

		Context: ctx,
	}
}

// NewGetServiceParamsWithHTTPClient creates a new GetServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceParamsWithHTTPClient(client *http.Client) *GetServiceParams {

	return &GetServiceParams{
		HTTPClient: client,
	}
}

/*GetServiceParams contains all the parameters to send to the API endpoint
for the get service operation typically these are written to a http.Request
*/
type GetServiceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service params
func (o *GetServiceParams) WithTimeout(timeout time.Duration) *GetServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service params
func (o *GetServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service params
func (o *GetServiceParams) WithContext(ctx context.Context) *GetServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service params
func (o *GetServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service params
func (o *GetServiceParams) WithHTTPClient(client *http.Client) *GetServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service params
func (o *GetServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
