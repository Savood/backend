// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetAllChatsForOfferingOKCode is the HTTP code returned for type GetAllChatsForOfferingOK
const GetAllChatsForOfferingOKCode int = 200

/*GetAllChatsForOfferingOK Object found and returned

swagger:response getAllChatsForOfferingOK
*/
type GetAllChatsForOfferingOK struct {

	/*
	  In: Body
	*/
	Payload models.Chats `json:"body,omitempty"`
}

// NewGetAllChatsForOfferingOK creates GetAllChatsForOfferingOK with default headers values
func NewGetAllChatsForOfferingOK() *GetAllChatsForOfferingOK {

	return &GetAllChatsForOfferingOK{}
}

// WithPayload adds the payload to the get all chats for offering o k response
func (o *GetAllChatsForOfferingOK) WithPayload(payload models.Chats) *GetAllChatsForOfferingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all chats for offering o k response
func (o *GetAllChatsForOfferingOK) SetPayload(payload models.Chats) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChatsForOfferingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Chats, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetAllChatsForOfferingNotFoundCode is the HTTP code returned for type GetAllChatsForOfferingNotFound
const GetAllChatsForOfferingNotFoundCode int = 404

/*GetAllChatsForOfferingNotFound Invalid parameter input was passed

swagger:response getAllChatsForOfferingNotFound
*/
type GetAllChatsForOfferingNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewGetAllChatsForOfferingNotFound creates GetAllChatsForOfferingNotFound with default headers values
func NewGetAllChatsForOfferingNotFound() *GetAllChatsForOfferingNotFound {

	return &GetAllChatsForOfferingNotFound{}
}

// WithPayload adds the payload to the get all chats for offering not found response
func (o *GetAllChatsForOfferingNotFound) WithPayload(payload *models.InvalidParameterInput) *GetAllChatsForOfferingNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all chats for offering not found response
func (o *GetAllChatsForOfferingNotFound) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChatsForOfferingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}