// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// PostUsersIDBackgroundimageJpegHandlerFunc turns a function with the right signature into a post users ID backgroundimage jpeg handler
type PostUsersIDBackgroundimageJpegHandlerFunc func(PostUsersIDBackgroundimageJpegParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUsersIDBackgroundimageJpegHandlerFunc) Handle(params PostUsersIDBackgroundimageJpegParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostUsersIDBackgroundimageJpegHandler interface for that can handle valid post users ID backgroundimage jpeg params
type PostUsersIDBackgroundimageJpegHandler interface {
	Handle(PostUsersIDBackgroundimageJpegParams, *models.Principal) middleware.Responder
}

// NewPostUsersIDBackgroundimageJpeg creates a new http.Handler for the post users ID backgroundimage jpeg operation
func NewPostUsersIDBackgroundimageJpeg(ctx *middleware.Context, handler PostUsersIDBackgroundimageJpegHandler) *PostUsersIDBackgroundimageJpeg {
	return &PostUsersIDBackgroundimageJpeg{Context: ctx, Handler: handler}
}

/*PostUsersIDBackgroundimageJpeg swagger:route POST /users/{id}/backgroundimage.jpeg users image postUsersIdBackgroundimageJpeg

Uploads avatar image.

*/
type PostUsersIDBackgroundimageJpeg struct {
	Context *middleware.Context
	Handler PostUsersIDBackgroundimageJpegHandler
}

func (o *PostUsersIDBackgroundimageJpeg) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostUsersIDBackgroundimageJpegParams()

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
