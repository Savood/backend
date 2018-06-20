// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"git.dhbw.chd.cx/savood/backend/restapi/operations"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/health"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/messages"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/users"

	models "git.dhbw.chd.cx/savood/backend/models"
	"git.dhbw.chd.cx/savood/backend/auth"
	"git.dhbw.chd.cx/savood/backend/database"
)

//go:generate swagger generate server --target .. --name  --spec ../../api-definition/swagger.yml --principal models.Principal

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

	// Applies when the "Authorization" header is set
	api.BearerAuth = func(token string) (*models.Principal, error) {

		authFunc := auth.GetAuthorizer(nil)

		return authFunc(token)
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.HealthHealthcheckGetHandler = health.HealthcheckGetHandlerFunc(func(params health.HealthcheckGetParams) middleware.Responder {
		if database.GetDatabase().Session.Ping() != nil {
			code := int32(503)
			message := "database unhealthy"
			return health.NewHealthcheckGetServiceUnavailable().WithPayload(&models.ErrorModel{Code: &code, Message: &message})
		}
		return health.NewHealthcheckGetOK()
	})

	api.MessagesCreateNewMessageHandler = messages.CreateNewMessageHandlerFunc(func(params messages.CreateNewMessageParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation messages.CreateNewMessage has not yet been implemented")
	})
	api.OfferingsCreateNewOfferingHandler = offerings.CreateNewOfferingHandlerFunc(func(params offerings.CreateNewOfferingParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation offerings.CreateNewOffering has not yet been implemented")
	})
	api.UsersCreateNewUserHandler = users.CreateNewUserHandlerFunc(func(params users.CreateNewUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.CreateNewUser has not yet been implemented")
	})
	api.MessagesDeleteMessageByIDHandler = messages.DeleteMessageByIDHandlerFunc(func(params messages.DeleteMessageByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation messages.DeleteMessageByID has not yet been implemented")
	})
	api.OfferingsDeleteOfferingByIDHandler = offerings.DeleteOfferingByIDHandlerFunc(func(params offerings.DeleteOfferingByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation offerings.DeleteOfferingByID has not yet been implemented")
	})
	api.UsersDeleteUserByIDHandler = users.DeleteUserByIDHandlerFunc(func(params users.DeleteUserByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.DeleteUserByID has not yet been implemented")
	})
	api.MessagesGetAllChatsHandler = messages.GetAllChatsHandlerFunc(func(params messages.GetAllChatsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation messages.GetAllChats has not yet been implemented")
	})
	api.GetAllChatsForOfferingHandler = operations.GetAllChatsForOfferingHandlerFunc(func(params operations.GetAllChatsForOfferingParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation .GetAllChatsForOffering has not yet been implemented")
	})
	api.MessagesGetAllMessagesForChatHandler = messages.GetAllMessagesForChatHandlerFunc(func(params messages.GetAllMessagesForChatParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation messages.GetAllMessagesForChat has not yet been implemented")
	})
	api.OfferingsGetFeedHandler = offerings.GetFeedHandlerFunc(func(params offerings.GetFeedParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation offerings.GetFeed has not yet been implemented")
	})
	api.MessagesGetMessageByIDHandler = messages.GetMessageByIDHandlerFunc(func(params messages.GetMessageByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation messages.GetMessageByID has not yet been implemented")
	})
	api.OfferingsGetOfferingByIDHandler = offerings.GetOfferingByIDHandlerFunc(func(params offerings.GetOfferingByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation offerings.GetOfferingByID has not yet been implemented")
	})
	api.OfferingsGetOfferingsHandler = offerings.GetOfferingsHandlerFunc(func(params offerings.GetOfferingsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation offerings.GetOfferings has not yet been implemented")
	})
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
	})
	api.MessagesUpdateMessageByIDHandler = messages.UpdateMessageByIDHandlerFunc(func(params messages.UpdateMessageByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation messages.UpdateMessageByID has not yet been implemented")
	})
	api.OfferingsUpdateOfferingByIDHandler = offerings.UpdateOfferingByIDHandlerFunc(func(params offerings.UpdateOfferingByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation offerings.UpdateOfferingByID has not yet been implemented")
	})
	api.UsersUpdateUserByIDHandler = users.UpdateUserByIDHandlerFunc(func(params users.UpdateUserByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.UpdateUserByID has not yet been implemented")
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
func configureServer(s *http.Server, scheme, addr string) {
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
