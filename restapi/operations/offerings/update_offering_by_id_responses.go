// Code generated by go-swagger; DO NOT EDIT.

package offerings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// UpdateOfferingByIDNoContentCode is the HTTP code returned for type UpdateOfferingByIDNoContent
const UpdateOfferingByIDNoContentCode int = 204

/*UpdateOfferingByIDNoContent Accepted; No response.

swagger:response updateOfferingByIdNoContent
*/
type UpdateOfferingByIDNoContent struct {
}

// NewUpdateOfferingByIDNoContent creates UpdateOfferingByIDNoContent with default headers values
func NewUpdateOfferingByIDNoContent() *UpdateOfferingByIDNoContent {

	return &UpdateOfferingByIDNoContent{}
}

// WriteResponse to the client
func (o *UpdateOfferingByIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// UpdateOfferingByIDBadRequestCode is the HTTP code returned for type UpdateOfferingByIDBadRequest
const UpdateOfferingByIDBadRequestCode int = 400

/*UpdateOfferingByIDBadRequest Invalid parameter input was passed

swagger:response updateOfferingByIdBadRequest
*/
type UpdateOfferingByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewUpdateOfferingByIDBadRequest creates UpdateOfferingByIDBadRequest with default headers values
func NewUpdateOfferingByIDBadRequest() *UpdateOfferingByIDBadRequest {

	return &UpdateOfferingByIDBadRequest{}
}

// WithPayload adds the payload to the update offering by Id bad request response
func (o *UpdateOfferingByIDBadRequest) WithPayload(payload *models.InvalidParameterInput) *UpdateOfferingByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update offering by Id bad request response
func (o *UpdateOfferingByIDBadRequest) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOfferingByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOfferingByIDForbiddenCode is the HTTP code returned for type UpdateOfferingByIDForbidden
const UpdateOfferingByIDForbiddenCode int = 403

/*UpdateOfferingByIDForbidden Generic Error

swagger:response updateOfferingByIdForbidden
*/
type UpdateOfferingByIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUpdateOfferingByIDForbidden creates UpdateOfferingByIDForbidden with default headers values
func NewUpdateOfferingByIDForbidden() *UpdateOfferingByIDForbidden {

	return &UpdateOfferingByIDForbidden{}
}

// WithPayload adds the payload to the update offering by Id forbidden response
func (o *UpdateOfferingByIDForbidden) WithPayload(payload *models.ErrorModel) *UpdateOfferingByIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update offering by Id forbidden response
func (o *UpdateOfferingByIDForbidden) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOfferingByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOfferingByIDNotFoundCode is the HTTP code returned for type UpdateOfferingByIDNotFound
const UpdateOfferingByIDNotFoundCode int = 404

/*UpdateOfferingByIDNotFound The specified resource was not found

swagger:response updateOfferingByIdNotFound
*/
type UpdateOfferingByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUpdateOfferingByIDNotFound creates UpdateOfferingByIDNotFound with default headers values
func NewUpdateOfferingByIDNotFound() *UpdateOfferingByIDNotFound {

	return &UpdateOfferingByIDNotFound{}
}

// WithPayload adds the payload to the update offering by Id not found response
func (o *UpdateOfferingByIDNotFound) WithPayload(payload *models.ErrorModel) *UpdateOfferingByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update offering by Id not found response
func (o *UpdateOfferingByIDNotFound) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOfferingByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOfferingByIDInternalServerErrorCode is the HTTP code returned for type UpdateOfferingByIDInternalServerError
const UpdateOfferingByIDInternalServerErrorCode int = 500

/*UpdateOfferingByIDInternalServerError Generic Error

swagger:response updateOfferingByIdInternalServerError
*/
type UpdateOfferingByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUpdateOfferingByIDInternalServerError creates UpdateOfferingByIDInternalServerError with default headers values
func NewUpdateOfferingByIDInternalServerError() *UpdateOfferingByIDInternalServerError {

	return &UpdateOfferingByIDInternalServerError{}
}

// WithPayload adds the payload to the update offering by Id internal server error response
func (o *UpdateOfferingByIDInternalServerError) WithPayload(payload *models.ErrorModel) *UpdateOfferingByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update offering by Id internal server error response
func (o *UpdateOfferingByIDInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOfferingByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
