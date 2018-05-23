// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/feed"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/message"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offering"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/profile"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/health"
)

//go:generate swagger generate server --target .. --name  --spec ../../api-definition/swagger.yml

func configureFlags(api *operations.SavoodAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SavoodAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HealthHealthcheckGetHandler = health.HealthcheckGetHandlerFunc(func(params health.HealthcheckGetParams) middleware.Responder {
		return health.NewHealthcheckGetOK()
	})

	api.FeedFeedGetHandler = feed.FeedGetHandlerFunc(func(params feed.FeedGetParams) middleware.Responder {
		return middleware.NotImplemented("operation feed.FeedGet has not yet been implemented")
	})
	api.MessageMessageIDDeleteHandler = message.MessageIDDeleteHandlerFunc(func(params message.MessageIDDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation message.MessageIDDelete has not yet been implemented")
	})
	api.MessageMessageIDGetHandler = message.MessageIDGetHandlerFunc(func(params message.MessageIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation message.MessageIDGet has not yet been implemented")
	})
	api.MessageMessageIDPutHandler = message.MessageIDPutHandlerFunc(func(params message.MessageIDPutParams) middleware.Responder {
		return middleware.NotImplemented("operation message.MessageIDPut has not yet been implemented")
	})
	api.MessageMessagePostHandler = message.MessagePostHandlerFunc(func(params message.MessagePostParams) middleware.Responder {
		return middleware.NotImplemented("operation message.MessagePost has not yet been implemented")
	})
	api.OfferingOfferingIDDeleteHandler = offering.OfferingIDDeleteHandlerFunc(func(params offering.OfferingIDDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.OfferingIDDelete has not yet been implemented")
	})
	api.OfferingOfferingIDGetHandler = offering.OfferingIDGetHandlerFunc(func(params offering.OfferingIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.OfferingIDGet has not yet been implemented")
	})
	api.OfferingOfferingIDPutHandler = offering.OfferingIDPutHandlerFunc(func(params offering.OfferingIDPutParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.OfferingIDPut has not yet been implemented")
	})
	api.OfferingOfferingPostHandler = offering.OfferingPostHandlerFunc(func(params offering.OfferingPostParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.OfferingPost has not yet been implemented")
	})
	api.ProfileProfileIDDeleteHandler = profile.ProfileIDDeleteHandlerFunc(func(params profile.ProfileIDDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.ProfileIDDelete has not yet been implemented")
	})
	api.ProfileProfileIDGetHandler = profile.ProfileIDGetHandlerFunc(func(params profile.ProfileIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.ProfileIDGet has not yet been implemented")
	})
	api.ProfileProfileIDPutHandler = profile.ProfileIDPutHandlerFunc(func(params profile.ProfileIDPutParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.ProfileIDPut has not yet been implemented")
	})
	api.ProfileProfilePostHandler = profile.ProfilePostHandlerFunc(func(params profile.ProfilePostParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.ProfilePost has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
