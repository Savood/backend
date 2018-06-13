// Code generated by go-swagger; DO NOT EDIT.

package offering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteOfferingByIDHandlerFunc turns a function with the right signature into a delete offering by Id handler
type DeleteOfferingByIDHandlerFunc func(DeleteOfferingByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteOfferingByIDHandlerFunc) Handle(params DeleteOfferingByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteOfferingByIDHandler interface for that can handle valid delete offering by Id params
type DeleteOfferingByIDHandler interface {
	Handle(DeleteOfferingByIDParams) middleware.Responder
}

// NewDeleteOfferingByID creates a new http.Handler for the delete offering by Id operation
func NewDeleteOfferingByID(ctx *middleware.Context, handler DeleteOfferingByIDHandler) *DeleteOfferingByID {
	return &DeleteOfferingByID{Context: ctx, Handler: handler}
}

/*DeleteOfferingByID swagger:route DELETE /offering/{id} offering deleteOfferingById

Delete an offering

*/
type DeleteOfferingByID struct {
	Context *middleware.Context
	Handler DeleteOfferingByIDHandler
}

func (o *DeleteOfferingByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteOfferingByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}