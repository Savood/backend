// Code generated by go-swagger; DO NOT EDIT.

package offerings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetOfferingByIDOKCode is the HTTP code returned for type GetOfferingByIDOK
const GetOfferingByIDOKCode int = 200

/*GetOfferingByIDOK Object found and returned

swagger:response getOfferingByIdOK
*/
type GetOfferingByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Offering `json:"body,omitempty"`
}

// NewGetOfferingByIDOK creates GetOfferingByIDOK with default headers values
func NewGetOfferingByIDOK() *GetOfferingByIDOK {

	return &GetOfferingByIDOK{}
}

// WithPayload adds the payload to the get offering by Id o k response
func (o *GetOfferingByIDOK) WithPayload(payload *models.Offering) *GetOfferingByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get offering by Id o k response
func (o *GetOfferingByIDOK) SetPayload(payload *models.Offering) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOfferingByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOfferingByIDNotFoundCode is the HTTP code returned for type GetOfferingByIDNotFound
const GetOfferingByIDNotFoundCode int = 404

/*GetOfferingByIDNotFound Invalid parameter input was passed

swagger:response getOfferingByIdNotFound
*/
type GetOfferingByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewGetOfferingByIDNotFound creates GetOfferingByIDNotFound with default headers values
func NewGetOfferingByIDNotFound() *GetOfferingByIDNotFound {

	return &GetOfferingByIDNotFound{}
}

// WithPayload adds the payload to the get offering by Id not found response
func (o *GetOfferingByIDNotFound) WithPayload(payload *models.InvalidParameterInput) *GetOfferingByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get offering by Id not found response
func (o *GetOfferingByIDNotFound) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOfferingByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
