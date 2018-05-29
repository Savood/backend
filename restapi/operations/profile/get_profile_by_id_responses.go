// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetProfileByIDOKCode is the HTTP code returned for type GetProfileByIDOK
const GetProfileByIDOKCode int = 200

/*GetProfileByIDOK Object found and returned

swagger:response getProfileByIdOK
*/
type GetProfileByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetProfileByIDOK creates GetProfileByIDOK with default headers values
func NewGetProfileByIDOK() *GetProfileByIDOK {

	return &GetProfileByIDOK{}
}

// WithPayload adds the payload to the get profile by Id o k response
func (o *GetProfileByIDOK) WithPayload(payload *models.Profile) *GetProfileByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile by Id o k response
func (o *GetProfileByIDOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProfileByIDNotFoundCode is the HTTP code returned for type GetProfileByIDNotFound
const GetProfileByIDNotFoundCode int = 404

/*GetProfileByIDNotFound Invalid parameter input was passed

swagger:response getProfileByIdNotFound
*/
type GetProfileByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.InvalidParameterInput `json:"body,omitempty"`
}

// NewGetProfileByIDNotFound creates GetProfileByIDNotFound with default headers values
func NewGetProfileByIDNotFound() *GetProfileByIDNotFound {

	return &GetProfileByIDNotFound{}
}

// WithPayload adds the payload to the get profile by Id not found response
func (o *GetProfileByIDNotFound) WithPayload(payload *models.InvalidParameterInput) *GetProfileByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile by Id not found response
func (o *GetProfileByIDNotFound) SetPayload(payload *models.InvalidParameterInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
