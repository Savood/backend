// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// DeleteMessageByIDNoContentCode is the HTTP code returned for type DeleteMessageByIDNoContent
const DeleteMessageByIDNoContentCode int = 204

/*DeleteMessageByIDNoContent Removed; No response.

swagger:response deleteMessageByIdNoContent
*/
type DeleteMessageByIDNoContent struct {
}

// NewDeleteMessageByIDNoContent creates DeleteMessageByIDNoContent with default headers values
func NewDeleteMessageByIDNoContent() *DeleteMessageByIDNoContent {

	return &DeleteMessageByIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteMessageByIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteMessageByIDNotFoundCode is the HTTP code returned for type DeleteMessageByIDNotFound
const DeleteMessageByIDNotFoundCode int = 404

/*DeleteMessageByIDNotFound Invalid parameter input was passed

swagger:response deleteMessageByIdNotFound
*/
type DeleteMessageByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewDeleteMessageByIDNotFound creates DeleteMessageByIDNotFound with default headers values
func NewDeleteMessageByIDNotFound() *DeleteMessageByIDNotFound {

	return &DeleteMessageByIDNotFound{}
}

// WithPayload adds the payload to the delete message by Id not found response
func (o *DeleteMessageByIDNotFound) WithPayload(payload *models.InvalidParameterInput) *DeleteMessageByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete message by Id not found response
func (o *DeleteMessageByIDNotFound) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMessageByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
