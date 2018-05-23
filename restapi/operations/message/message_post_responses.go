// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// MessagePostOKCode is the HTTP code returned for type MessagePostOK
const MessagePostOKCode int = 200

/*MessagePostOK Object found and returned

swagger:response messagePostOK
*/
type MessagePostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewMessagePostOK creates MessagePostOK with default headers values
func NewMessagePostOK() *MessagePostOK {

	return &MessagePostOK{}
}

// WithPayload adds the payload to the message post o k response
func (o *MessagePostOK) WithPayload(payload *models.Message) *MessagePostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the message post o k response
func (o *MessagePostOK) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MessagePostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// MessagePostNotFoundCode is the HTTP code returned for type MessagePostNotFound
const MessagePostNotFoundCode int = 404

/*MessagePostNotFound Invalid parameter input was passed

swagger:response messagePostNotFound
*/
type MessagePostNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewMessagePostNotFound creates MessagePostNotFound with default headers values
func NewMessagePostNotFound() *MessagePostNotFound {

	return &MessagePostNotFound{}
}

// WithPayload adds the payload to the message post not found response
func (o *MessagePostNotFound) WithPayload(payload *models.InvalidParameterInput) *MessagePostNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the message post not found response
func (o *MessagePostNotFound) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MessagePostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
