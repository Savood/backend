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
	"git.dhbw.chd.cx/savood/backend/restapi/operations/health"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/message"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offering"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/profile"
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

	api.MessageCreateNewMessageHandler = message.CreateNewMessageHandlerFunc(func(params message.CreateNewMessageParams) middleware.Responder {
		return middleware.NotImplemented("operation message.CreateNewMessage has not yet been implemented")
	})
	api.OfferingCreateNewOfferingHandler = offering.CreateNewOfferingHandlerFunc(func(params offering.CreateNewOfferingParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.CreateNewOffering has not yet been implemented")
	})
	api.ProfileCreateNewProfileHandler = profile.CreateNewProfileHandlerFunc(func(params profile.CreateNewProfileParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.CreateNewProfile has not yet been implemented")
	})
	api.MessageDeleteMessageByIDHandler = message.DeleteMessageByIDHandlerFunc(func(params message.DeleteMessageByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation message.DeleteMessageByID has not yet been implemented")
	})
	api.OfferingDeleteOfferingByIDHandler = offering.DeleteOfferingByIDHandlerFunc(func(params offering.DeleteOfferingByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.DeleteOfferingByID has not yet been implemented")
	})
	api.ProfileDeleteProfileByIDHandler = profile.DeleteProfileByIDHandlerFunc(func(params profile.DeleteProfileByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.DeleteProfileByID has not yet been implemented")
	})
	api.FeedGetFeedHandler = feed.GetFeedHandlerFunc(func(params feed.GetFeedParams) middleware.Responder {
		return middleware.NotImplemented("operation feed.GetFeed has not yet been implemented")
	})
	api.MessageGetMessageByIDHandler = message.GetMessageByIDHandlerFunc(func(params message.GetMessageByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation message.GetMessageByID has not yet been implemented")
	})
	api.OfferingGetOfferingByIDHandler = offering.GetOfferingByIDHandlerFunc(func(params offering.GetOfferingByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.GetOfferingByID has not yet been implemented")
	})
	api.ProfileGetProfileByIDHandler = profile.GetProfileByIDHandlerFunc(func(params profile.GetProfileByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.GetProfileByID has not yet been implemented")
	})
	api.HealthHealthcheckGetHandler = health.HealthcheckGetHandlerFunc(func(params health.HealthcheckGetParams) middleware.Responder {
		return health.NewHealthcheckGetOK()
	})
	api.MessageUpdateMessageByIDHandler = message.UpdateMessageByIDHandlerFunc(func(params message.UpdateMessageByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation message.UpdateMessageByID has not yet been implemented")
	})
	api.OfferingUpdateOfferingByIDHandler = offering.UpdateOfferingByIDHandlerFunc(func(params offering.UpdateOfferingByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation offering.UpdateOfferingByID has not yet been implemented")
	})
	api.ProfileUpdateProfileByIDHandler = profile.UpdateProfileByIDHandlerFunc(func(params profile.UpdateProfileByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation profile.UpdateProfileByID has not yet been implemented")
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
