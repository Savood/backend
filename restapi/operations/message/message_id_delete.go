// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// MessageIDDeleteHandlerFunc turns a function with the right signature into a message Id delete handler
type MessageIDDeleteHandlerFunc func(MessageIDDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MessageIDDeleteHandlerFunc) Handle(params MessageIDDeleteParams) middleware.Responder {
	return fn(params)
}

// MessageIDDeleteHandler interface for that can handle valid message Id delete params
type MessageIDDeleteHandler interface {
	Handle(MessageIDDeleteParams) middleware.Responder
}

// NewMessageIDDelete creates a new http.Handler for the message Id delete operation
func NewMessageIDDelete(ctx *middleware.Context, handler MessageIDDeleteHandler) *MessageIDDelete {
	return &MessageIDDelete{Context: ctx, Handler: handler}
}

/*MessageIDDelete swagger:route DELETE /message/{id} message messageIdDelete

Delete an message

*/
type MessageIDDelete struct {
	Context *middleware.Context
	Handler MessageIDDeleteHandler
}

func (o *MessageIDDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMessageIDDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
