package kloud

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIKloudCheckCredentialParams creates a new PostRemoteAPIKloudCheckCredentialParams object
// with the default values initialized.
func NewPostRemoteAPIKloudCheckCredentialParams() *PostRemoteAPIKloudCheckCredentialParams {
	var ()
	return &PostRemoteAPIKloudCheckCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIKloudCheckCredentialParamsWithTimeout creates a new PostRemoteAPIKloudCheckCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIKloudCheckCredentialParamsWithTimeout(timeout time.Duration) *PostRemoteAPIKloudCheckCredentialParams {
	var ()
	return &PostRemoteAPIKloudCheckCredentialParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIKloudCheckCredentialParamsWithContext creates a new PostRemoteAPIKloudCheckCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIKloudCheckCredentialParamsWithContext(ctx context.Context) *PostRemoteAPIKloudCheckCredentialParams {
	var ()
	return &PostRemoteAPIKloudCheckCredentialParams{

		Context: ctx,
	}
}

/*PostRemoteAPIKloudCheckCredentialParams contains all the parameters to send to the API endpoint
for the post remote API kloud check credential operation typically these are written to a http.Request
*/
type PostRemoteAPIKloudCheckCredentialParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API kloud check credential params
func (o *PostRemoteAPIKloudCheckCredentialParams) WithTimeout(timeout time.Duration) *PostRemoteAPIKloudCheckCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API kloud check credential params
func (o *PostRemoteAPIKloudCheckCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API kloud check credential params
func (o *PostRemoteAPIKloudCheckCredentialParams) WithContext(ctx context.Context) *PostRemoteAPIKloudCheckCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API kloud check credential params
func (o *PostRemoteAPIKloudCheckCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API kloud check credential params
func (o *PostRemoteAPIKloudCheckCredentialParams) WithBody(body models.DefaultSelector) *PostRemoteAPIKloudCheckCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API kloud check credential params
func (o *PostRemoteAPIKloudCheckCredentialParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIKloudCheckCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
