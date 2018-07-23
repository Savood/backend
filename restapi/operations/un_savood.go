// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// UnSavoodHandlerFunc turns a function with the right signature into a un savood handler
type UnSavoodHandlerFunc func(UnSavoodParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UnSavoodHandlerFunc) Handle(params UnSavoodParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UnSavoodHandler interface for that can handle valid un savood params
type UnSavoodHandler interface {
	Handle(UnSavoodParams, *models.Principal) middleware.Responder
}

// NewUnSavood creates a new http.Handler for the un savood operation
func NewUnSavood(ctx *middleware.Context, handler UnSavoodHandler) *UnSavood {
	return &UnSavood{Context: ctx, Handler: handler}
}

/*UnSavood swagger:route POST /unSavood rpc-calls offerings unSavood

Places a savood on an offering

*/
type UnSavood struct {
	Context *middleware.Context
	Handler UnSavoodHandler
}

func (o *UnSavood) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUnSavoodParams()

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