// Code generated by go-swagger; DO NOT EDIT.

package offerings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// CreateNewOfferingHandlerFunc turns a function with the right signature into a create new offering handler
type CreateNewOfferingHandlerFunc func(CreateNewOfferingParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateNewOfferingHandlerFunc) Handle(params CreateNewOfferingParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateNewOfferingHandler interface for that can handle valid create new offering params
type CreateNewOfferingHandler interface {
	Handle(CreateNewOfferingParams, *models.Principal) middleware.Responder
}

// NewCreateNewOffering creates a new http.Handler for the create new offering operation
func NewCreateNewOffering(ctx *middleware.Context, handler CreateNewOfferingHandler) *CreateNewOffering {
	return &CreateNewOffering{Context: ctx, Handler: handler}
}

/*CreateNewOffering swagger:route POST /offerings offerings createNewOffering

Add a new offering

*/
type CreateNewOffering struct {
	Context *middleware.Context
	Handler CreateNewOfferingHandler
}

func (o *CreateNewOffering) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateNewOfferingParams()

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
