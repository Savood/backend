// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetAllMessagesForChatHandlerFunc turns a function with the right signature into a get all messages for chat handler
type GetAllMessagesForChatHandlerFunc func(GetAllMessagesForChatParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllMessagesForChatHandlerFunc) Handle(params GetAllMessagesForChatParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAllMessagesForChatHandler interface for that can handle valid get all messages for chat params
type GetAllMessagesForChatHandler interface {
	Handle(GetAllMessagesForChatParams, *models.Principal) middleware.Responder
}

// NewGetAllMessagesForChat creates a new http.Handler for the get all messages for chat operation
func NewGetAllMessagesForChat(ctx *middleware.Context, handler GetAllMessagesForChatHandler) *GetAllMessagesForChat {
	return &GetAllMessagesForChat{Context: ctx, Handler: handler}
}

/*GetAllMessagesForChat swagger:route GET /chats/{chatID}/messages messages getAllMessagesForChat

Add a new message

*/
type GetAllMessagesForChat struct {
	Context *middleware.Context
	Handler GetAllMessagesForChatHandler
}

func (o *GetAllMessagesForChat) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllMessagesForChatParams()

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
