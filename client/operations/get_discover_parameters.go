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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/choria-io/pdbproxy/models"
)

// NewGetDiscoverParams creates a new GetDiscoverParams object
// with the default values initialized.
func NewGetDiscoverParams() *GetDiscoverParams {
	var ()
	return &GetDiscoverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoverParamsWithTimeout creates a new GetDiscoverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDiscoverParamsWithTimeout(timeout time.Duration) *GetDiscoverParams {
	var ()
	return &GetDiscoverParams{

		timeout: timeout,
	}
}

// NewGetDiscoverParamsWithContext creates a new GetDiscoverParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDiscoverParamsWithContext(ctx context.Context) *GetDiscoverParams {
	var ()
	return &GetDiscoverParams{

		Context: ctx,
	}
}

// NewGetDiscoverParamsWithHTTPClient creates a new GetDiscoverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDiscoverParamsWithHTTPClient(client *http.Client) *GetDiscoverParams {
	var ()
	return &GetDiscoverParams{
		HTTPClient: client,
	}
}

/*GetDiscoverParams contains all the parameters to send to the API endpoint
for the get discover operation typically these are written to a http.Request
*/
type GetDiscoverParams struct {

	/*Request
	  Filter description

	*/
	Request *models.DiscoveryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get discover params
func (o *GetDiscoverParams) WithTimeout(timeout time.Duration) *GetDiscoverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discover params
func (o *GetDiscoverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discover params
func (o *GetDiscoverParams) WithContext(ctx context.Context) *GetDiscoverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discover params
func (o *GetDiscoverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discover params
func (o *GetDiscoverParams) WithHTTPClient(client *http.Client) *GetDiscoverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discover params
func (o *GetDiscoverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the get discover params
func (o *GetDiscoverParams) WithRequest(request *models.DiscoveryRequest) *GetDiscoverParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the get discover params
func (o *GetDiscoverParams) SetRequest(request *models.DiscoveryRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request == nil {
		o.Request = new(models.DiscoveryRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
