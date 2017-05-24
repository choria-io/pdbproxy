package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSetSetParams creates a new GetSetSetParams object
// with the default values initialized.
func NewGetSetSetParams() *GetSetSetParams {
	var (
		discoverDefault = bool(false)
	)
	return &GetSetSetParams{
		Discover: &discoverDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSetSetParamsWithTimeout creates a new GetSetSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSetSetParamsWithTimeout(timeout time.Duration) *GetSetSetParams {
	var (
		discoverDefault = bool(false)
	)
	return &GetSetSetParams{
		Discover: &discoverDefault,

		timeout: timeout,
	}
}

// NewGetSetSetParamsWithContext creates a new GetSetSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSetSetParamsWithContext(ctx context.Context) *GetSetSetParams {
	var (
		discoverDefault = bool(false)
	)
	return &GetSetSetParams{
		Discover: &discoverDefault,

		Context: ctx,
	}
}

// NewGetSetSetParamsWithHTTPClient creates a new GetSetSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSetSetParamsWithHTTPClient(client *http.Client) *GetSetSetParams {
	var (
		discoverDefault = bool(false)
	)
	return &GetSetSetParams{
		Discover:   &discoverDefault,
		HTTPClient: client,
	}
}

/*GetSetSetParams contains all the parameters to send to the API endpoint
for the get set set operation typically these are written to a http.Request
*/
type GetSetSetParams struct {

	/*Discover
	  Also include the matching nodes if the set is a PQL query

	*/
	Discover *bool
	/*Set
	  Node set to retrieve

	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get set set params
func (o *GetSetSetParams) WithTimeout(timeout time.Duration) *GetSetSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get set set params
func (o *GetSetSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get set set params
func (o *GetSetSetParams) WithContext(ctx context.Context) *GetSetSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get set set params
func (o *GetSetSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get set set params
func (o *GetSetSetParams) WithHTTPClient(client *http.Client) *GetSetSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get set set params
func (o *GetSetSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiscover adds the discover to the get set set params
func (o *GetSetSetParams) WithDiscover(discover *bool) *GetSetSetParams {
	o.SetDiscover(discover)
	return o
}

// SetDiscover adds the discover to the get set set params
func (o *GetSetSetParams) SetDiscover(discover *bool) {
	o.Discover = discover
}

// WithSet adds the set to the get set set params
func (o *GetSetSetParams) WithSet(set string) *GetSetSetParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the get set set params
func (o *GetSetSetParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *GetSetSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Discover != nil {

		// query param discover
		var qrDiscover bool
		if o.Discover != nil {
			qrDiscover = *o.Discover
		}
		qDiscover := swag.FormatBool(qrDiscover)
		if qDiscover != "" {
			if err := r.SetQueryParam("discover", qDiscover); err != nil {
				return err
			}
		}

	}

	// path param set
	if err := r.SetPathParam("set", o.Set); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}