// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// UpdateUserByIDNoContentCode is the HTTP code returned for type UpdateUserByIDNoContent
const UpdateUserByIDNoContentCode int = 204

/*UpdateUserByIDNoContent Accepted; No response.

swagger:response updateUserByIdNoContent
*/
type UpdateUserByIDNoContent struct {
}

// NewUpdateUserByIDNoContent creates UpdateUserByIDNoContent with default headers values
func NewUpdateUserByIDNoContent() *UpdateUserByIDNoContent {

	return &UpdateUserByIDNoContent{}
}

// WriteResponse to the client
func (o *UpdateUserByIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// UpdateUserByIDBadRequestCode is the HTTP code returned for type UpdateUserByIDBadRequest
const UpdateUserByIDBadRequestCode int = 400

/*UpdateUserByIDBadRequest Invalid parameter input was passed

swagger:response updateUserByIdBadRequest
*/
type UpdateUserByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewUpdateUserByIDBadRequest creates UpdateUserByIDBadRequest with default headers values
func NewUpdateUserByIDBadRequest() *UpdateUserByIDBadRequest {

	return &UpdateUserByIDBadRequest{}
}

// WithPayload adds the payload to the update user by Id bad request response
func (o *UpdateUserByIDBadRequest) WithPayload(payload *models.InvalidParameterInput) *UpdateUserByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user by Id bad request response
func (o *UpdateUserByIDBadRequest) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
