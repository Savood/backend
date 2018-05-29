// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetMessageByIDOKCode is the HTTP code returned for type GetMessageByIDOK
const GetMessageByIDOKCode int = 200

/*GetMessageByIDOK Object found and returned

swagger:response getMessageByIdOK
*/
type GetMessageByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewGetMessageByIDOK creates GetMessageByIDOK with default headers values
func NewGetMessageByIDOK() *GetMessageByIDOK {

	return &GetMessageByIDOK{}
}

// WithPayload adds the payload to the get message by Id o k response
func (o *GetMessageByIDOK) WithPayload(payload *models.Message) *GetMessageByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get message by Id o k response
func (o *GetMessageByIDOK) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMessageByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMessageByIDNotFoundCode is the HTTP code returned for type GetMessageByIDNotFound
const GetMessageByIDNotFoundCode int = 404

/*GetMessageByIDNotFound Invalid parameter input was passed

swagger:response getMessageByIdNotFound
*/
type GetMessageByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewGetMessageByIDNotFound creates GetMessageByIDNotFound with default headers values
func NewGetMessageByIDNotFound() *GetMessageByIDNotFound {

	return &GetMessageByIDNotFound{}
}

// WithPayload adds the payload to the get message by Id not found response
func (o *GetMessageByIDNotFound) WithPayload(payload *models.InvalidParameterInput) *GetMessageByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get message by Id not found response
func (o *GetMessageByIDNotFound) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMessageByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
