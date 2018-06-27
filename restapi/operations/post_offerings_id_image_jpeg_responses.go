// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// PostOfferingsIDImageJpegNoContentCode is the HTTP code returned for type PostOfferingsIDImageJpegNoContent
const PostOfferingsIDImageJpegNoContentCode int = 204

/*PostOfferingsIDImageJpegNoContent Accepted; No response.

swagger:response postOfferingsIdImageJpegNoContent
*/
type PostOfferingsIDImageJpegNoContent struct {
}

// NewPostOfferingsIDImageJpegNoContent creates PostOfferingsIDImageJpegNoContent with default headers values
func NewPostOfferingsIDImageJpegNoContent() *PostOfferingsIDImageJpegNoContent {

	return &PostOfferingsIDImageJpegNoContent{}
}

// WriteResponse to the client
func (o *PostOfferingsIDImageJpegNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostOfferingsIDImageJpegBadRequestCode is the HTTP code returned for type PostOfferingsIDImageJpegBadRequest
const PostOfferingsIDImageJpegBadRequestCode int = 400

/*PostOfferingsIDImageJpegBadRequest Invalid parameter input was passed

swagger:response postOfferingsIdImageJpegBadRequest
*/
type PostOfferingsIDImageJpegBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewPostOfferingsIDImageJpegBadRequest creates PostOfferingsIDImageJpegBadRequest with default headers values
func NewPostOfferingsIDImageJpegBadRequest() *PostOfferingsIDImageJpegBadRequest {

	return &PostOfferingsIDImageJpegBadRequest{}
}

// WithPayload adds the payload to the post offerings Id image jpeg bad request response
func (o *PostOfferingsIDImageJpegBadRequest) WithPayload(payload *models.InvalidParameterInput) *PostOfferingsIDImageJpegBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post offerings Id image jpeg bad request response
func (o *PostOfferingsIDImageJpegBadRequest) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOfferingsIDImageJpegBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOfferingsIDImageJpegForbiddenCode is the HTTP code returned for type PostOfferingsIDImageJpegForbidden
const PostOfferingsIDImageJpegForbiddenCode int = 403

/*PostOfferingsIDImageJpegForbidden Unauthorized

swagger:response postOfferingsIdImageJpegForbidden
*/
type PostOfferingsIDImageJpegForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewPostOfferingsIDImageJpegForbidden creates PostOfferingsIDImageJpegForbidden with default headers values
func NewPostOfferingsIDImageJpegForbidden() *PostOfferingsIDImageJpegForbidden {

	return &PostOfferingsIDImageJpegForbidden{}
}

// WithPayload adds the payload to the post offerings Id image jpeg forbidden response
func (o *PostOfferingsIDImageJpegForbidden) WithPayload(payload *models.ErrorModel) *PostOfferingsIDImageJpegForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post offerings Id image jpeg forbidden response
func (o *PostOfferingsIDImageJpegForbidden) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOfferingsIDImageJpegForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOfferingsIDImageJpegInternalServerErrorCode is the HTTP code returned for type PostOfferingsIDImageJpegInternalServerError
const PostOfferingsIDImageJpegInternalServerErrorCode int = 500

/*PostOfferingsIDImageJpegInternalServerError Generic Error

swagger:response postOfferingsIdImageJpegInternalServerError
*/
type PostOfferingsIDImageJpegInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewPostOfferingsIDImageJpegInternalServerError creates PostOfferingsIDImageJpegInternalServerError with default headers values
func NewPostOfferingsIDImageJpegInternalServerError() *PostOfferingsIDImageJpegInternalServerError {

	return &PostOfferingsIDImageJpegInternalServerError{}
}

// WithPayload adds the payload to the post offerings Id image jpeg internal server error response
func (o *PostOfferingsIDImageJpegInternalServerError) WithPayload(payload *models.ErrorModel) *PostOfferingsIDImageJpegInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post offerings Id image jpeg internal server error response
func (o *PostOfferingsIDImageJpegInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOfferingsIDImageJpegInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
