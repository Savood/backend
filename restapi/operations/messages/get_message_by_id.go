// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/backend/models"
)

// GetMessageByIDHandlerFunc turns a function with the right signature into a get message by Id handler
type GetMessageByIDHandlerFunc func(GetMessageByIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMessageByIDHandlerFunc) Handle(params GetMessageByIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetMessageByIDHandler interface for that can handle valid get message by Id params
type GetMessageByIDHandler interface {
	Handle(GetMessageByIDParams, *models.Principal) middleware.Responder
}

// NewGetMessageByID creates a new http.Handler for the get message by Id operation
func NewGetMessageByID(ctx *middleware.Context, handler GetMessageByIDHandler) *GetMessageByID {
	return &GetMessageByID{Context: ctx, Handler: handler}
}

/*GetMessageByID swagger:route GET /chats/{chatID}/messages/{id} messages getMessageById

Display a message

*/
type GetMessageByID struct {
	Context *middleware.Context
	Handler GetMessageByIDHandler
}

func (o *GetMessageByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMessageByIDParams()

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
