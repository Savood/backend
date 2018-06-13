// Code generated by go-swagger; DO NOT EDIT.

package feed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFeedParams creates a new GetFeedParams object
// no default values defined in spec.
func NewGetFeedParams() GetFeedParams {

	return GetFeedParams{}
}

// GetFeedParams contains all the bound params for the get feed operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFeed
type GetFeedParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Location string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFeedParams() beforehand.
func (o *GetFeedParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLocation, qhkLocation, _ := qs.GetOK("location")
	if err := o.bindLocation(qLocation, qhkLocation, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFeedParams) bindLocation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("location", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("location", "query", raw); err != nil {
		return err
	}

	o.Location = raw

	return nil
}