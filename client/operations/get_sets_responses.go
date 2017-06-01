package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/choria-io/discovery_proxy/models"
)

// GetSetsReader is a Reader for the GetSets structure.
type GetSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSetsOK creates a GetSetsOK with default headers values
func NewGetSetsOK() *GetSetsOK {
	return &GetSetsOK{}
}

/*GetSetsOK handles this case with default header values.

Known Sets
*/
type GetSetsOK struct {
	Payload *models.Sets
}

func (o *GetSetsOK) Error() string {
	return fmt.Sprintf("[GET /sets][%d] getSetsOK  %+v", 200, o.Payload)
}

func (o *GetSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
