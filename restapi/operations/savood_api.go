// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"git.dhbw.chd.cx/savood/backend/restapi/operations/health"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/messages"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offerings"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/users"

	models "git.dhbw.chd.cx/savood/backend/models"
)

// NewSavoodAPI creates a new Savood instance
func NewSavoodAPI(spec *loads.Document) *SavoodAPI {
	return &SavoodAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		ImageJpegProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("imageJpeg producer has not yet been implemented")
		}),
		GetOfferingsIDImageJpegHandler: GetOfferingsIDImageJpegHandlerFunc(func(params GetOfferingsIDImageJpegParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetOfferingsIDImageJpeg has not yet been implemented")
		}),
		GetUsersIDBackgroundimageJpegHandler: GetUsersIDBackgroundimageJpegHandlerFunc(func(params GetUsersIDBackgroundimageJpegParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetUsersIDBackgroundimageJpeg has not yet been implemented")
		}),
		GetUsersIDImageJpegHandler: GetUsersIDImageJpegHandlerFunc(func(params GetUsersIDImageJpegParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetUsersIDImageJpeg has not yet been implemented")
		}),
		PostOfferingsIDImageJpegHandler: PostOfferingsIDImageJpegHandlerFunc(func(params PostOfferingsIDImageJpegParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PostOfferingsIDImageJpeg has not yet been implemented")
		}),
		PostUsersIDBackgroundimageJpegHandler: PostUsersIDBackgroundimageJpegHandlerFunc(func(params PostUsersIDBackgroundimageJpegParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PostUsersIDBackgroundimageJpeg has not yet been implemented")
		}),
		PostUsersIDImageJpegHandler: PostUsersIDImageJpegHandlerFunc(func(params PostUsersIDImageJpegParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PostUsersIDImageJpeg has not yet been implemented")
		}),
		MessagesCreateNewMessageHandler: messages.CreateNewMessageHandlerFunc(func(params messages.CreateNewMessageParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation MessagesCreateNewMessage has not yet been implemented")
		}),
		OfferingsCreateNewOfferingHandler: offerings.CreateNewOfferingHandlerFunc(func(params offerings.CreateNewOfferingParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation OfferingsCreateNewOffering has not yet been implemented")
		}),
		UsersCreateNewUserHandler: users.CreateNewUserHandlerFunc(func(params users.CreateNewUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UsersCreateNewUser has not yet been implemented")
		}),
		MessagesDeleteMessageByIDHandler: messages.DeleteMessageByIDHandlerFunc(func(params messages.DeleteMessageByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation MessagesDeleteMessageByID has not yet been implemented")
		}),
		OfferingsDeleteOfferingByIDHandler: offerings.DeleteOfferingByIDHandlerFunc(func(params offerings.DeleteOfferingByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation OfferingsDeleteOfferingByID has not yet been implemented")
		}),
		UsersDeleteUserByIDHandler: users.DeleteUserByIDHandlerFunc(func(params users.DeleteUserByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UsersDeleteUserByID has not yet been implemented")
		}),
		MessagesGetAllChatsHandler: messages.GetAllChatsHandlerFunc(func(params messages.GetAllChatsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation MessagesGetAllChats has not yet been implemented")
		}),
		GetAllChatsForOfferingHandler: GetAllChatsForOfferingHandlerFunc(func(params GetAllChatsForOfferingParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetAllChatsForOffering has not yet been implemented")
		}),
		MessagesGetAllMessagesForChatHandler: messages.GetAllMessagesForChatHandlerFunc(func(params messages.GetAllMessagesForChatParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation MessagesGetAllMessagesForChat has not yet been implemented")
		}),
		OfferingsGetFeedHandler: offerings.GetFeedHandlerFunc(func(params offerings.GetFeedParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation OfferingsGetFeed has not yet been implemented")
		}),
		MessagesGetMessageByIDHandler: messages.GetMessageByIDHandlerFunc(func(params messages.GetMessageByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation MessagesGetMessageByID has not yet been implemented")
		}),
		OfferingsGetOfferingByIDHandler: offerings.GetOfferingByIDHandlerFunc(func(params offerings.GetOfferingByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation OfferingsGetOfferingByID has not yet been implemented")
		}),
		OfferingsGetOfferingsHandler: offerings.GetOfferingsHandlerFunc(func(params offerings.GetOfferingsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation OfferingsGetOfferings has not yet been implemented")
		}),
		UsersGetUserByIDHandler: users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UsersGetUserByID has not yet been implemented")
		}),
		HealthHealthcheckGetHandler: health.HealthcheckGetHandlerFunc(func(params health.HealthcheckGetParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthHealthcheckGet has not yet been implemented")
		}),
		PlaceSavoodHandler: PlaceSavoodHandlerFunc(func(params PlaceSavoodParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PlaceSavood has not yet been implemented")
		}),
		MessagesUpdateMessageByIDHandler: messages.UpdateMessageByIDHandlerFunc(func(params messages.UpdateMessageByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation MessagesUpdateMessageByID has not yet been implemented")
		}),
		OfferingsUpdateOfferingByIDHandler: offerings.UpdateOfferingByIDHandlerFunc(func(params offerings.UpdateOfferingByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation OfferingsUpdateOfferingByID has not yet been implemented")
		}),
		UsersUpdateUserByIDHandler: users.UpdateUserByIDHandlerFunc(func(params users.UpdateUserByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation UsersUpdateUserByID has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (bearer) Authorization from header param [Authorization] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*SavoodAPI denn nur lebendiges food tut gut */
type SavoodAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// ImageJpegProducer registers a producer for a "image/jpeg" mime type
	ImageJpegProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GetOfferingsIDImageJpegHandler sets the operation handler for the get offerings ID image jpeg operation
	GetOfferingsIDImageJpegHandler GetOfferingsIDImageJpegHandler
	// GetUsersIDBackgroundimageJpegHandler sets the operation handler for the get users ID backgroundimage jpeg operation
	GetUsersIDBackgroundimageJpegHandler GetUsersIDBackgroundimageJpegHandler
	// GetUsersIDImageJpegHandler sets the operation handler for the get users ID image jpeg operation
	GetUsersIDImageJpegHandler GetUsersIDImageJpegHandler
	// PostOfferingsIDImageJpegHandler sets the operation handler for the post offerings ID image jpeg operation
	PostOfferingsIDImageJpegHandler PostOfferingsIDImageJpegHandler
	// PostUsersIDBackgroundimageJpegHandler sets the operation handler for the post users ID backgroundimage jpeg operation
	PostUsersIDBackgroundimageJpegHandler PostUsersIDBackgroundimageJpegHandler
	// PostUsersIDImageJpegHandler sets the operation handler for the post users ID image jpeg operation
	PostUsersIDImageJpegHandler PostUsersIDImageJpegHandler
	// MessagesCreateNewMessageHandler sets the operation handler for the create new message operation
	MessagesCreateNewMessageHandler messages.CreateNewMessageHandler
	// CreateNewOfferingHandler sets the operation handler for the create new offering operation
	OfferingsCreateNewOfferingHandler offerings.CreateNewOfferingHandler
	// UsersCreateNewUserHandler sets the operation handler for the create new user operation
	UsersCreateNewUserHandler users.CreateNewUserHandler
	// MessagesDeleteMessageByIDHandler sets the operation handler for the delete message by Id operation
	MessagesDeleteMessageByIDHandler messages.DeleteMessageByIDHandler
	// OfferingsDeleteOfferingByIDHandler sets the operation handler for the delete offering by Id operation
	OfferingsDeleteOfferingByIDHandler offerings.DeleteOfferingByIDHandler
	// UsersDeleteUserByIDHandler sets the operation handler for the delete user by Id operation
	UsersDeleteUserByIDHandler users.DeleteUserByIDHandler
	// MessagesGetAllChatsHandler sets the operation handler for the get all chats operation
	MessagesGetAllChatsHandler messages.GetAllChatsHandler
	// GetAllChatsForOfferingHandler sets the operation handler for the get all chats for offering operation
	GetAllChatsForOfferingHandler GetAllChatsForOfferingHandler
	// MessagesGetAllMessagesForChatHandler sets the operation handler for the get all messages for chat operation
	MessagesGetAllMessagesForChatHandler messages.GetAllMessagesForChatHandler
	// OfferingsGetFeedHandler sets the operation handler for the get feed operation
	OfferingsGetFeedHandler offerings.GetFeedHandler
	// MessagesGetMessageByIDHandler sets the operation handler for the get message by Id operation
	MessagesGetMessageByIDHandler messages.GetMessageByIDHandler
	// OfferingsGetOfferingByIDHandler sets the operation handler for the get offering by Id operation
	OfferingsGetOfferingByIDHandler offerings.GetOfferingByIDHandler
	// OfferingsGetOfferingsHandler sets the operation handler for the get offerings operation
	OfferingsGetOfferingsHandler offerings.GetOfferingsHandler
	// UsersGetUserByIDHandler sets the operation handler for the get user by Id operation
	UsersGetUserByIDHandler users.GetUserByIDHandler
	// HealthHealthcheckGetHandler sets the operation handler for the healthcheck get operation
	HealthHealthcheckGetHandler health.HealthcheckGetHandler
	// PlaceSavoodHandler sets the operation handler for the place savood operation
	PlaceSavoodHandler PlaceSavoodHandler
	// MessagesUpdateMessageByIDHandler sets the operation handler for the update message by Id operation
	MessagesUpdateMessageByIDHandler messages.UpdateMessageByIDHandler
	// OfferingsUpdateOfferingByIDHandler sets the operation handler for the update offering by Id operation
	OfferingsUpdateOfferingByIDHandler offerings.UpdateOfferingByIDHandler
	// UsersUpdateUserByIDHandler sets the operation handler for the update user by Id operation
	UsersUpdateUserByIDHandler users.UpdateUserByIDHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SavoodAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SavoodAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SavoodAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SavoodAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SavoodAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SavoodAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SavoodAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SavoodAPI
func (o *SavoodAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ImageJpegProducer == nil {
		unregistered = append(unregistered, "ImageJpegProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.GetOfferingsIDImageJpegHandler == nil {
		unregistered = append(unregistered, "GetOfferingsIDImageJpegHandler")
	}

	if o.GetUsersIDBackgroundimageJpegHandler == nil {
		unregistered = append(unregistered, "GetUsersIDBackgroundimageJpegHandler")
	}

	if o.GetUsersIDImageJpegHandler == nil {
		unregistered = append(unregistered, "GetUsersIDImageJpegHandler")
	}

	if o.PostOfferingsIDImageJpegHandler == nil {
		unregistered = append(unregistered, "PostOfferingsIDImageJpegHandler")
	}

	if o.PostUsersIDBackgroundimageJpegHandler == nil {
		unregistered = append(unregistered, "PostUsersIDBackgroundimageJpegHandler")
	}

	if o.PostUsersIDImageJpegHandler == nil {
		unregistered = append(unregistered, "PostUsersIDImageJpegHandler")
	}

	if o.MessagesCreateNewMessageHandler == nil {
		unregistered = append(unregistered, "messages.CreateNewMessageHandler")
	}

	if o.OfferingsCreateNewOfferingHandler == nil {
		unregistered = append(unregistered, "offerings.CreateNewOfferingHandler")
	}

	if o.UsersCreateNewUserHandler == nil {
		unregistered = append(unregistered, "users.CreateNewUserHandler")
	}

	if o.MessagesDeleteMessageByIDHandler == nil {
		unregistered = append(unregistered, "messages.DeleteMessageByIDHandler")
	}

	if o.OfferingsDeleteOfferingByIDHandler == nil {
		unregistered = append(unregistered, "offerings.DeleteOfferingByIDHandler")
	}

	if o.UsersDeleteUserByIDHandler == nil {
		unregistered = append(unregistered, "users.DeleteUserByIDHandler")
	}

	if o.MessagesGetAllChatsHandler == nil {
		unregistered = append(unregistered, "messages.GetAllChatsHandler")
	}

	if o.GetAllChatsForOfferingHandler == nil {
		unregistered = append(unregistered, "GetAllChatsForOfferingHandler")
	}

	if o.MessagesGetAllMessagesForChatHandler == nil {
		unregistered = append(unregistered, "messages.GetAllMessagesForChatHandler")
	}

	if o.OfferingsGetFeedHandler == nil {
		unregistered = append(unregistered, "offerings.GetFeedHandler")
	}

	if o.MessagesGetMessageByIDHandler == nil {
		unregistered = append(unregistered, "messages.GetMessageByIDHandler")
	}

	if o.OfferingsGetOfferingByIDHandler == nil {
		unregistered = append(unregistered, "offerings.GetOfferingByIDHandler")
	}

	if o.OfferingsGetOfferingsHandler == nil {
		unregistered = append(unregistered, "offerings.GetOfferingsHandler")
	}

	if o.UsersGetUserByIDHandler == nil {
		unregistered = append(unregistered, "users.GetUserByIDHandler")
	}

	if o.HealthHealthcheckGetHandler == nil {
		unregistered = append(unregistered, "health.HealthcheckGetHandler")
	}

	if o.PlaceSavoodHandler == nil {
		unregistered = append(unregistered, "PlaceSavoodHandler")
	}

	if o.MessagesUpdateMessageByIDHandler == nil {
		unregistered = append(unregistered, "messages.UpdateMessageByIDHandler")
	}

	if o.OfferingsUpdateOfferingByIDHandler == nil {
		unregistered = append(unregistered, "offerings.UpdateOfferingByIDHandler")
	}

	if o.UsersUpdateUserByIDHandler == nil {
		unregistered = append(unregistered, "users.UpdateUserByIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SavoodAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SavoodAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "bearer":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.BearerAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *SavoodAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *SavoodAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SavoodAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "image/jpeg":
			result["image/jpeg"] = o.ImageJpegProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SavoodAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the savood API
func (o *SavoodAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SavoodAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offerings/{id}/image.jpeg"] = NewGetOfferingsIDImageJpeg(o.context, o.GetOfferingsIDImageJpegHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}/backgroundimage.jpeg"] = NewGetUsersIDBackgroundimageJpeg(o.context, o.GetUsersIDBackgroundimageJpegHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}/image.jpeg"] = NewGetUsersIDImageJpeg(o.context, o.GetUsersIDImageJpegHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/offerings/{id}/image.jpeg"] = NewPostOfferingsIDImageJpeg(o.context, o.PostOfferingsIDImageJpegHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{id}/backgroundimage.jpeg"] = NewPostUsersIDBackgroundimageJpeg(o.context, o.PostUsersIDBackgroundimageJpegHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{id}/image.jpeg"] = NewPostUsersIDImageJpeg(o.context, o.PostUsersIDImageJpegHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/chats/{chatID}/messages"] = messages.NewCreateNewMessage(o.context, o.MessagesCreateNewMessageHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/offerings"] = offerings.NewCreateNewOffering(o.context, o.OfferingsCreateNewOfferingHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = users.NewCreateNewUser(o.context, o.UsersCreateNewUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/chats/{chatID}/messages/{id}"] = messages.NewDeleteMessageByID(o.context, o.MessagesDeleteMessageByIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/offerings/{id}"] = offerings.NewDeleteOfferingByID(o.context, o.OfferingsDeleteOfferingByIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}"] = users.NewDeleteUserByID(o.context, o.UsersDeleteUserByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chats"] = messages.NewGetAllChats(o.context, o.MessagesGetAllChatsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offerings/{id}/chats"] = NewGetAllChatsForOffering(o.context, o.GetAllChatsForOfferingHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chats/{chatID}/messages"] = messages.NewGetAllMessagesForChat(o.context, o.MessagesGetAllMessagesForChatHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/feed"] = offerings.NewGetFeed(o.context, o.OfferingsGetFeedHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chats/{chatID}/messages/{id}"] = messages.NewGetMessageByID(o.context, o.MessagesGetMessageByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offerings/{id}"] = offerings.NewGetOfferingByID(o.context, o.OfferingsGetOfferingByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offerings"] = offerings.NewGetOfferings(o.context, o.OfferingsGetOfferingsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}"] = users.NewGetUserByID(o.context, o.UsersGetUserByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthcheck"] = health.NewHealthcheckGet(o.context, o.HealthHealthcheckGetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/placeSavood"] = NewPlaceSavood(o.context, o.PlaceSavoodHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/chats/{chatID}/messages/{id}"] = messages.NewUpdateMessageByID(o.context, o.MessagesUpdateMessageByIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/offerings/{id}"] = offerings.NewUpdateOfferingByID(o.context, o.OfferingsUpdateOfferingByIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/users/{id}"] = users.NewUpdateUserByID(o.context, o.UsersUpdateUserByIDHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SavoodAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SavoodAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SavoodAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SavoodAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
