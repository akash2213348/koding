package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JGroupUpdatePermissionsReader is a Reader for the JGroupUpdatePermissions structure.
type JGroupUpdatePermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupUpdatePermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupUpdatePermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupUpdatePermissionsOK creates a JGroupUpdatePermissionsOK with default headers values
func NewJGroupUpdatePermissionsOK() *JGroupUpdatePermissionsOK {
	return &JGroupUpdatePermissionsOK{}
}

/*JGroupUpdatePermissionsOK handles this case with default header values.

OK
*/
type JGroupUpdatePermissionsOK struct {
	Payload JGroupUpdatePermissionsOKBody
}

func (o *JGroupUpdatePermissionsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.updatePermissions/{id}][%d] jGroupUpdatePermissionsOK  %+v", 200, o.Payload)
}

func (o *JGroupUpdatePermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JGroupUpdatePermissionsOKBody j group update permissions o k body
swagger:model JGroupUpdatePermissionsOKBody
*/
type JGroupUpdatePermissionsOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JGroupUpdatePermissionsOKBody) UnmarshalJSON(raw []byte) error {

	var jGroupUpdatePermissionsOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &jGroupUpdatePermissionsOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = jGroupUpdatePermissionsOKBodyAO0

	var jGroupUpdatePermissionsOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jGroupUpdatePermissionsOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jGroupUpdatePermissionsOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JGroupUpdatePermissionsOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jGroupUpdatePermissionsOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupUpdatePermissionsOKBodyAO0)

	jGroupUpdatePermissionsOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupUpdatePermissionsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j group update permissions o k body
func (o *JGroupUpdatePermissionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
