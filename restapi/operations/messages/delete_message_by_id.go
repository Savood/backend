// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteMessageByIDHandlerFunc turns a function with the right signature into a delete message by Id handler
type DeleteMessageByIDHandlerFunc func(DeleteMessageByIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMessageByIDHandlerFunc) Handle(params DeleteMessageByIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteMessageByIDHandler interface for that can handle valid delete message by Id params
type DeleteMessageByIDHandler interface {
	Handle(DeleteMessageByIDParams, interface{}) middleware.Responder
}

// NewDeleteMessageByID creates a new http.Handler for the delete message by Id operation
func NewDeleteMessageByID(ctx *middleware.Context, handler DeleteMessageByIDHandler) *DeleteMessageByID {
	return &DeleteMessageByID{Context: ctx, Handler: handler}
}

/*DeleteMessageByID swagger:route DELETE /chats/{chatID}/messages/{id} messages deleteMessageById

Delete a message

*/
type DeleteMessageByID struct {
	Context *middleware.Context
	Handler DeleteMessageByIDHandler
}

func (o *DeleteMessageByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMessageByIDParams()

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
