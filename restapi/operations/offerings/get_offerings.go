// Code generated by go-swagger; DO NOT EDIT.

package offerings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetOfferingsHandlerFunc turns a function with the right signature into a get offerings handler
type GetOfferingsHandlerFunc func(GetOfferingsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOfferingsHandlerFunc) Handle(params GetOfferingsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetOfferingsHandler interface for that can handle valid get offerings params
type GetOfferingsHandler interface {
	Handle(GetOfferingsParams, *models.Principal) middleware.Responder
}

// NewGetOfferings creates a new http.Handler for the get offerings operation
func NewGetOfferings(ctx *middleware.Context, handler GetOfferingsHandler) *GetOfferings {
	return &GetOfferings{Context: ctx, Handler: handler}
}

/*GetOfferings swagger:route GET /offerings offerings getOfferings

Display owned or requested offerings

*/
type GetOfferings struct {
	Context *middleware.Context
	Handler GetOfferingsHandler
}

func (o *GetOfferings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOfferingsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
