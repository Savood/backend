// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// UnSavoodOKCode is the HTTP code returned for type UnSavoodOK
const UnSavoodOKCode int = 200

/*UnSavoodOK Savood is deleted

swagger:response unSavoodOK
*/
type UnSavoodOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessObject `json:"body,omitempty"`
}

// NewUnSavoodOK creates UnSavoodOK with default headers values
func NewUnSavoodOK() *UnSavoodOK {

	return &UnSavoodOK{}
}

// WithPayload adds the payload to the un savood o k response
func (o *UnSavoodOK) WithPayload(payload *models.SuccessObject) *UnSavoodOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un savood o k response
func (o *UnSavoodOK) SetPayload(payload *models.SuccessObject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnSavoodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnSavoodBadRequestCode is the HTTP code returned for type UnSavoodBadRequest
const UnSavoodBadRequestCode int = 400

/*UnSavoodBadRequest Invalid parameter input was passed

swagger:response unSavoodBadRequest
*/
type UnSavoodBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewUnSavoodBadRequest creates UnSavoodBadRequest with default headers values
func NewUnSavoodBadRequest() *UnSavoodBadRequest {

	return &UnSavoodBadRequest{}
}

// WithPayload adds the payload to the un savood bad request response
func (o *UnSavoodBadRequest) WithPayload(payload *models.InvalidParameterInput) *UnSavoodBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un savood bad request response
func (o *UnSavoodBadRequest) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnSavoodBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnSavoodForbiddenCode is the HTTP code returned for type UnSavoodForbidden
const UnSavoodForbiddenCode int = 403

/*UnSavoodForbidden Generic Error

swagger:response unSavoodForbidden
*/
type UnSavoodForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUnSavoodForbidden creates UnSavoodForbidden with default headers values
func NewUnSavoodForbidden() *UnSavoodForbidden {

	return &UnSavoodForbidden{}
}

// WithPayload adds the payload to the un savood forbidden response
func (o *UnSavoodForbidden) WithPayload(payload *models.ErrorModel) *UnSavoodForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un savood forbidden response
func (o *UnSavoodForbidden) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnSavoodForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnSavoodNotFoundCode is the HTTP code returned for type UnSavoodNotFound
const UnSavoodNotFoundCode int = 404

/*UnSavoodNotFound The specified resource was not found

swagger:response unSavoodNotFound
*/
type UnSavoodNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUnSavoodNotFound creates UnSavoodNotFound with default headers values
func NewUnSavoodNotFound() *UnSavoodNotFound {

	return &UnSavoodNotFound{}
}

// WithPayload adds the payload to the un savood not found response
func (o *UnSavoodNotFound) WithPayload(payload *models.ErrorModel) *UnSavoodNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un savood not found response
func (o *UnSavoodNotFound) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnSavoodNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnSavoodInternalServerErrorCode is the HTTP code returned for type UnSavoodInternalServerError
const UnSavoodInternalServerErrorCode int = 500

/*UnSavoodInternalServerError Generic Error

swagger:response unSavoodInternalServerError
*/
type UnSavoodInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUnSavoodInternalServerError creates UnSavoodInternalServerError with default headers values
func NewUnSavoodInternalServerError() *UnSavoodInternalServerError {

	return &UnSavoodInternalServerError{}
}

// WithPayload adds the payload to the un savood internal server error response
func (o *UnSavoodInternalServerError) WithPayload(payload *models.ErrorModel) *UnSavoodInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the un savood internal server error response
func (o *UnSavoodInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnSavoodInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
