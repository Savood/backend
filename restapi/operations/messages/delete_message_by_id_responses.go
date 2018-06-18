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

// DeleteMessageByIDBadRequestCode is the HTTP code returned for type DeleteMessageByIDBadRequest
const DeleteMessageByIDBadRequestCode int = 400

/*DeleteMessageByIDBadRequest Invalid parameter input was passed

swagger:response deleteMessageByIdBadRequest
*/
type DeleteMessageByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewDeleteMessageByIDBadRequest creates DeleteMessageByIDBadRequest with default headers values
func NewDeleteMessageByIDBadRequest() *DeleteMessageByIDBadRequest {

	return &DeleteMessageByIDBadRequest{}
}

// WithPayload adds the payload to the delete message by Id bad request response
func (o *DeleteMessageByIDBadRequest) WithPayload(payload *models.InvalidParameterInput) *DeleteMessageByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete message by Id bad request response
func (o *DeleteMessageByIDBadRequest) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMessageByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
