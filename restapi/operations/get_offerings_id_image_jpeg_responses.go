// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetOfferingsIDImageJpegOKCode is the HTTP code returned for type GetOfferingsIDImageJpegOK
const GetOfferingsIDImageJpegOKCode int = 200

/*GetOfferingsIDImageJpegOK OK

swagger:response getOfferingsIdImageJpegOK
*/
type GetOfferingsIDImageJpegOK struct {

	/*
	  In: Body
	*/
	Payload models.GetOfferingsIDImageJpegOKBody `json:"body,omitempty"`
}

// NewGetOfferingsIDImageJpegOK creates GetOfferingsIDImageJpegOK with default headers values
func NewGetOfferingsIDImageJpegOK() *GetOfferingsIDImageJpegOK {

	return &GetOfferingsIDImageJpegOK{}
}

// WithPayload adds the payload to the get offerings Id image jpeg o k response
func (o *GetOfferingsIDImageJpegOK) WithPayload(payload models.GetOfferingsIDImageJpegOKBody) *GetOfferingsIDImageJpegOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get offerings Id image jpeg o k response
func (o *GetOfferingsIDImageJpegOK) SetPayload(payload models.GetOfferingsIDImageJpegOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOfferingsIDImageJpegOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetOfferingsIDImageJpegNotFoundCode is the HTTP code returned for type GetOfferingsIDImageJpegNotFound
const GetOfferingsIDImageJpegNotFoundCode int = 404

/*GetOfferingsIDImageJpegNotFound The specified resource was not found

swagger:response getOfferingsIdImageJpegNotFound
*/
type GetOfferingsIDImageJpegNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewGetOfferingsIDImageJpegNotFound creates GetOfferingsIDImageJpegNotFound with default headers values
func NewGetOfferingsIDImageJpegNotFound() *GetOfferingsIDImageJpegNotFound {

	return &GetOfferingsIDImageJpegNotFound{}
}

// WithPayload adds the payload to the get offerings Id image jpeg not found response
func (o *GetOfferingsIDImageJpegNotFound) WithPayload(payload *models.ErrorModel) *GetOfferingsIDImageJpegNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get offerings Id image jpeg not found response
func (o *GetOfferingsIDImageJpegNotFound) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOfferingsIDImageJpegNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOfferingsIDImageJpegInternalServerErrorCode is the HTTP code returned for type GetOfferingsIDImageJpegInternalServerError
const GetOfferingsIDImageJpegInternalServerErrorCode int = 500

/*GetOfferingsIDImageJpegInternalServerError Generic Error

swagger:response getOfferingsIdImageJpegInternalServerError
*/
type GetOfferingsIDImageJpegInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewGetOfferingsIDImageJpegInternalServerError creates GetOfferingsIDImageJpegInternalServerError with default headers values
func NewGetOfferingsIDImageJpegInternalServerError() *GetOfferingsIDImageJpegInternalServerError {

	return &GetOfferingsIDImageJpegInternalServerError{}
}

// WithPayload adds the payload to the get offerings Id image jpeg internal server error response
func (o *GetOfferingsIDImageJpegInternalServerError) WithPayload(payload *models.ErrorModel) *GetOfferingsIDImageJpegInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get offerings Id image jpeg internal server error response
func (o *GetOfferingsIDImageJpegInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOfferingsIDImageJpegInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
