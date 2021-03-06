// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/choria-io/discovery_proxy/models"
)

// PostSetOKCode is the HTTP code returned for type PostSetOK
const PostSetOKCode int = 200

/*PostSetOK Basic successful request

swagger:response postSetOK
*/
type PostSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessModel `json:"body,omitempty"`
}

// NewPostSetOK creates PostSetOK with default headers values
func NewPostSetOK() *PostSetOK {
	return &PostSetOK{}
}

// WithPayload adds the payload to the post set o k response
func (o *PostSetOK) WithPayload(payload *models.SuccessModel) *PostSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post set o k response
func (o *PostSetOK) SetPayload(payload *models.SuccessModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSetBadRequestCode is the HTTP code returned for type PostSetBadRequest
const PostSetBadRequestCode int = 400

/*PostSetBadRequest Standard Error Format

swagger:response postSetBadRequest
*/
type PostSetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewPostSetBadRequest creates PostSetBadRequest with default headers values
func NewPostSetBadRequest() *PostSetBadRequest {
	return &PostSetBadRequest{}
}

// WithPayload adds the payload to the post set bad request response
func (o *PostSetBadRequest) WithPayload(payload *models.ErrorModel) *PostSetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post set bad request response
func (o *PostSetBadRequest) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSetInternalServerErrorCode is the HTTP code returned for type PostSetInternalServerError
const PostSetInternalServerErrorCode int = 500

/*PostSetInternalServerError Standard Error Format

swagger:response postSetInternalServerError
*/
type PostSetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewPostSetInternalServerError creates PostSetInternalServerError with default headers values
func NewPostSetInternalServerError() *PostSetInternalServerError {
	return &PostSetInternalServerError{}
}

// WithPayload adds the payload to the post set internal server error response
func (o *PostSetInternalServerError) WithPayload(payload *models.ErrorModel) *PostSetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post set internal server error response
func (o *PostSetInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
