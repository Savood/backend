// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersIDBackgroundimageJpegParams creates a new GetUsersIDBackgroundimageJpegParams object
// no default values defined in spec.
func NewGetUsersIDBackgroundimageJpegParams() GetUsersIDBackgroundimageJpegParams {

	return GetUsersIDBackgroundimageJpegParams{}
}

// GetUsersIDBackgroundimageJpegParams contains all the bound params for the get users ID backgroundimage jpeg operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetUsersIDBackgroundimageJpeg
type GetUsersIDBackgroundimageJpegParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*If either width or height is set to 0, it will be set to an aspect ratio preserving value.
	  In: query
	*/
	Height *float64
	/*
	  Required: true
	  In: path
	*/
	ID string
	/*If either width or height is set to 0, it will be set to an aspect ratio preserving value.
	  In: query
	*/
	Width *float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUsersIDBackgroundimageJpegParams() beforehand.
func (o *GetUsersIDBackgroundimageJpegParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qHeight, qhkHeight, _ := qs.GetOK("height")
	if err := o.bindHeight(qHeight, qhkHeight, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qWidth, qhkWidth, _ := qs.GetOK("width")
	if err := o.bindWidth(qWidth, qhkWidth, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHeight binds and validates parameter Height from query.
func (o *GetUsersIDBackgroundimageJpegParams) bindHeight(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("height", "query", "float64", raw)
	}
	o.Height = &value

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetUsersIDBackgroundimageJpegParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

// bindWidth binds and validates parameter Width from query.
func (o *GetUsersIDBackgroundimageJpegParams) bindWidth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("width", "query", "float64", raw)
	}
	o.Width = &value

	return nil
}
