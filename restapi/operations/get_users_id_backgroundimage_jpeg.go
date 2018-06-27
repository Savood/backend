// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// GetUsersIDBackgroundimageJpegHandlerFunc turns a function with the right signature into a get users ID backgroundimage jpeg handler
type GetUsersIDBackgroundimageJpegHandlerFunc func(GetUsersIDBackgroundimageJpegParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersIDBackgroundimageJpegHandlerFunc) Handle(params GetUsersIDBackgroundimageJpegParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUsersIDBackgroundimageJpegHandler interface for that can handle valid get users ID backgroundimage jpeg params
type GetUsersIDBackgroundimageJpegHandler interface {
	Handle(GetUsersIDBackgroundimageJpegParams, *models.Principal) middleware.Responder
}

// NewGetUsersIDBackgroundimageJpeg creates a new http.Handler for the get users ID backgroundimage jpeg operation
func NewGetUsersIDBackgroundimageJpeg(ctx *middleware.Context, handler GetUsersIDBackgroundimageJpegHandler) *GetUsersIDBackgroundimageJpeg {
	return &GetUsersIDBackgroundimageJpeg{Context: ctx, Handler: handler}
}

/*GetUsersIDBackgroundimageJpeg swagger:route GET /users/{id}/backgroundimage.jpeg users image getUsersIdBackgroundimageJpeg

Gets the avatar image.

*/
type GetUsersIDBackgroundimageJpeg struct {
	Context *middleware.Context
	Handler GetUsersIDBackgroundimageJpegHandler
}

func (o *GetUsersIDBackgroundimageJpeg) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersIDBackgroundimageJpegParams()

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
