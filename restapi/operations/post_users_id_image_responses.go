// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// PostUsersIDImageNoContentCode is the HTTP code returned for type PostUsersIDImageNoContent
const PostUsersIDImageNoContentCode int = 204

/*PostUsersIDImageNoContent Accepted; No response.

swagger:response postUsersIdImageNoContent
*/
type PostUsersIDImageNoContent struct {
}

// NewPostUsersIDImageNoContent creates PostUsersIDImageNoContent with default headers values
func NewPostUsersIDImageNoContent() *PostUsersIDImageNoContent {

	return &PostUsersIDImageNoContent{}
}

// WriteResponse to the client
func (o *PostUsersIDImageNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostUsersIDImageBadRequestCode is the HTTP code returned for type PostUsersIDImageBadRequest
const PostUsersIDImageBadRequestCode int = 400

/*PostUsersIDImageBadRequest Invalid parameter input was passed

swagger:response postUsersIdImageBadRequest
*/
type PostUsersIDImageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewPostUsersIDImageBadRequest creates PostUsersIDImageBadRequest with default headers values
func NewPostUsersIDImageBadRequest() *PostUsersIDImageBadRequest {

	return &PostUsersIDImageBadRequest{}
}

// WithPayload adds the payload to the post users Id image bad request response
func (o *PostUsersIDImageBadRequest) WithPayload(payload *models.InvalidParameterInput) *PostUsersIDImageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users Id image bad request response
func (o *PostUsersIDImageBadRequest) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersIDImageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUsersIDImageForbiddenCode is the HTTP code returned for type PostUsersIDImageForbidden
const PostUsersIDImageForbiddenCode int = 403

/*PostUsersIDImageForbidden Unauthorized

swagger:response postUsersIdImageForbidden
*/
type PostUsersIDImageForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewPostUsersIDImageForbidden creates PostUsersIDImageForbidden with default headers values
func NewPostUsersIDImageForbidden() *PostUsersIDImageForbidden {

	return &PostUsersIDImageForbidden{}
}

// WithPayload adds the payload to the post users Id image forbidden response
func (o *PostUsersIDImageForbidden) WithPayload(payload *models.ErrorModel) *PostUsersIDImageForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users Id image forbidden response
func (o *PostUsersIDImageForbidden) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersIDImageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUsersIDImageInternalServerErrorCode is the HTTP code returned for type PostUsersIDImageInternalServerError
const PostUsersIDImageInternalServerErrorCode int = 500

/*PostUsersIDImageInternalServerError Generic Error

swagger:response postUsersIdImageInternalServerError
*/
type PostUsersIDImageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewPostUsersIDImageInternalServerError creates PostUsersIDImageInternalServerError with default headers values
func NewPostUsersIDImageInternalServerError() *PostUsersIDImageInternalServerError {

	return &PostUsersIDImageInternalServerError{}
}

// WithPayload adds the payload to the post users Id image internal server error response
func (o *PostUsersIDImageInternalServerError) WithPayload(payload *models.ErrorModel) *PostUsersIDImageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users Id image internal server error response
func (o *PostUsersIDImageInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersIDImageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
