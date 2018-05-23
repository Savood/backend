// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ProfilePostHandlerFunc turns a function with the right signature into a profile post handler
type ProfilePostHandlerFunc func(ProfilePostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProfilePostHandlerFunc) Handle(params ProfilePostParams) middleware.Responder {
	return fn(params)
}

// ProfilePostHandler interface for that can handle valid profile post params
type ProfilePostHandler interface {
	Handle(ProfilePostParams) middleware.Responder
}

// NewProfilePost creates a new http.Handler for the profile post operation
func NewProfilePost(ctx *middleware.Context, handler ProfilePostHandler) *ProfilePost {
	return &ProfilePost{Context: ctx, Handler: handler}
}

/*ProfilePost swagger:route POST /profile profile profilePost

Add a new profile

*/
type ProfilePost struct {
	Context *middleware.Context
	Handler ProfilePostHandler
}

func (o *ProfilePost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProfilePostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
