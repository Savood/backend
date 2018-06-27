// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// PlaceSavoodHandlerFunc turns a function with the right signature into a place savood handler
type PlaceSavoodHandlerFunc func(PlaceSavoodParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PlaceSavoodHandlerFunc) Handle(params PlaceSavoodParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PlaceSavoodHandler interface for that can handle valid place savood params
type PlaceSavoodHandler interface {
	Handle(PlaceSavoodParams, *models.Principal) middleware.Responder
}

// NewPlaceSavood creates a new http.Handler for the place savood operation
func NewPlaceSavood(ctx *middleware.Context, handler PlaceSavoodHandler) *PlaceSavood {
	return &PlaceSavood{Context: ctx, Handler: handler}
}

/*PlaceSavood swagger:route POST /placeSavood rpc-calls offerings placeSavood

Places a savood on an offering

*/
type PlaceSavood struct {
	Context *middleware.Context
	Handler PlaceSavoodHandler
}

func (o *PlaceSavood) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPlaceSavoodParams()

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
