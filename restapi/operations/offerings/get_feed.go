// Code generated by go-swagger; DO NOT EDIT.

package offerings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFeedHandlerFunc turns a function with the right signature into a get feed handler
type GetFeedHandlerFunc func(GetFeedParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFeedHandlerFunc) Handle(params GetFeedParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFeedHandler interface for that can handle valid get feed params
type GetFeedHandler interface {
	Handle(GetFeedParams, interface{}) middleware.Responder
}

// NewGetFeed creates a new http.Handler for the get feed operation
func NewGetFeed(ctx *middleware.Context, handler GetFeedHandler) *GetFeed {
	return &GetFeed{Context: ctx, Handler: handler}
}

/*GetFeed swagger:route GET /feed offerings getFeed

Display a feed of nearby offerings

*/
type GetFeed struct {
	Context *middleware.Context
	Handler GetFeedHandler
}

func (o *GetFeed) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFeedParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
