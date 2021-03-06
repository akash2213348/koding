package j_user

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

// NewPostRemoteAPIJUserSetSSHKeysParams creates a new PostRemoteAPIJUserSetSSHKeysParams object
// with the default values initialized.
func NewPostRemoteAPIJUserSetSSHKeysParams() *PostRemoteAPIJUserSetSSHKeysParams {
	var ()
	return &PostRemoteAPIJUserSetSSHKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserSetSSHKeysParamsWithTimeout creates a new PostRemoteAPIJUserSetSSHKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserSetSSHKeysParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserSetSSHKeysParams {
	var ()
	return &PostRemoteAPIJUserSetSSHKeysParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserSetSSHKeysParamsWithContext creates a new PostRemoteAPIJUserSetSSHKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserSetSSHKeysParamsWithContext(ctx context.Context) *PostRemoteAPIJUserSetSSHKeysParams {
	var ()
	return &PostRemoteAPIJUserSetSSHKeysParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserSetSSHKeysParams contains all the parameters to send to the API endpoint
for the post remote API j user set SSH keys operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserSetSSHKeysParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j user set SSH keys params
func (o *PostRemoteAPIJUserSetSSHKeysParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserSetSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user set SSH keys params
func (o *PostRemoteAPIJUserSetSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user set SSH keys params
func (o *PostRemoteAPIJUserSetSSHKeysParams) WithContext(ctx context.Context) *PostRemoteAPIJUserSetSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user set SSH keys params
func (o *PostRemoteAPIJUserSetSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j user set SSH keys params
func (o *PostRemoteAPIJUserSetSSHKeysParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJUserSetSSHKeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j user set SSH keys params
func (o *PostRemoteAPIJUserSetSSHKeysParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserSetSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
