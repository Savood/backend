// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
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

	"git.dhbw.chd.cx/savood/backend/restapi/operations/feed"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/health"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/message"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/offering"
	"git.dhbw.chd.cx/savood/backend/restapi/operations/profile"
)

// NewSavoodAPI creates a new Savood instance
func NewSavoodAPI(spec *loads.Document) *SavoodAPI {
	return &SavoodAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		MessageCreateNewMessageHandler: message.CreateNewMessageHandlerFunc(func(params message.CreateNewMessageParams) middleware.Responder {
			return middleware.NotImplemented("operation MessageCreateNewMessage has not yet been implemented")
		}),
		OfferingCreateNewOfferingHandler: offering.CreateNewOfferingHandlerFunc(func(params offering.CreateNewOfferingParams) middleware.Responder {
			return middleware.NotImplemented("operation OfferingCreateNewOffering has not yet been implemented")
		}),
		ProfileCreateNewProfileHandler: profile.CreateNewProfileHandlerFunc(func(params profile.CreateNewProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileCreateNewProfile has not yet been implemented")
		}),
		MessageDeleteMessageByIDHandler: message.DeleteMessageByIDHandlerFunc(func(params message.DeleteMessageByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation MessageDeleteMessageByID has not yet been implemented")
		}),
		OfferingDeleteOfferingByIDHandler: offering.DeleteOfferingByIDHandlerFunc(func(params offering.DeleteOfferingByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation OfferingDeleteOfferingByID has not yet been implemented")
		}),
		ProfileDeleteProfileByIDHandler: profile.DeleteProfileByIDHandlerFunc(func(params profile.DeleteProfileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileDeleteProfileByID has not yet been implemented")
		}),
		FeedGetFeedHandler: feed.GetFeedHandlerFunc(func(params feed.GetFeedParams) middleware.Responder {
			return middleware.NotImplemented("operation FeedGetFeed has not yet been implemented")
		}),
		MessageGetMessageByIDHandler: message.GetMessageByIDHandlerFunc(func(params message.GetMessageByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation MessageGetMessageByID has not yet been implemented")
		}),
		OfferingGetOfferingByIDHandler: offering.GetOfferingByIDHandlerFunc(func(params offering.GetOfferingByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation OfferingGetOfferingByID has not yet been implemented")
		}),
		ProfileGetProfileByIDHandler: profile.GetProfileByIDHandlerFunc(func(params profile.GetProfileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileGetProfileByID has not yet been implemented")
		}),
		HealthHealthcheckGetHandler: health.HealthcheckGetHandlerFunc(func(params health.HealthcheckGetParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthHealthcheckGet has not yet been implemented")
		}),
		MessageUpdateMessageByIDHandler: message.UpdateMessageByIDHandlerFunc(func(params message.UpdateMessageByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation MessageUpdateMessageByID has not yet been implemented")
		}),
		OfferingUpdateOfferingByIDHandler: offering.UpdateOfferingByIDHandlerFunc(func(params offering.UpdateOfferingByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation OfferingUpdateOfferingByID has not yet been implemented")
		}),
		ProfileUpdateProfileByIDHandler: profile.UpdateProfileByIDHandlerFunc(func(params profile.UpdateProfileByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileUpdateProfileByID has not yet been implemented")
		}),
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

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// MessageCreateNewMessageHandler sets the operation handler for the create new message operation
	MessageCreateNewMessageHandler message.CreateNewMessageHandler
	// OfferingCreateNewOfferingHandler sets the operation handler for the create new offering operation
	OfferingCreateNewOfferingHandler offering.CreateNewOfferingHandler
	// ProfileCreateNewProfileHandler sets the operation handler for the create new profile operation
	ProfileCreateNewProfileHandler profile.CreateNewProfileHandler
	// MessageDeleteMessageByIDHandler sets the operation handler for the delete message by Id operation
	MessageDeleteMessageByIDHandler message.DeleteMessageByIDHandler
	// OfferingDeleteOfferingByIDHandler sets the operation handler for the delete offering by Id operation
	OfferingDeleteOfferingByIDHandler offering.DeleteOfferingByIDHandler
	// ProfileDeleteProfileByIDHandler sets the operation handler for the delete profile by Id operation
	ProfileDeleteProfileByIDHandler profile.DeleteProfileByIDHandler
	// FeedGetFeedHandler sets the operation handler for the get feed operation
	FeedGetFeedHandler feed.GetFeedHandler
	// MessageGetMessageByIDHandler sets the operation handler for the get message by Id operation
	MessageGetMessageByIDHandler message.GetMessageByIDHandler
	// OfferingGetOfferingByIDHandler sets the operation handler for the get offering by Id operation
	OfferingGetOfferingByIDHandler offering.GetOfferingByIDHandler
	// ProfileGetProfileByIDHandler sets the operation handler for the get profile by Id operation
	ProfileGetProfileByIDHandler profile.GetProfileByIDHandler
	// HealthHealthcheckGetHandler sets the operation handler for the healthcheck get operation
	HealthHealthcheckGetHandler health.HealthcheckGetHandler
	// MessageUpdateMessageByIDHandler sets the operation handler for the update message by Id operation
	MessageUpdateMessageByIDHandler message.UpdateMessageByIDHandler
	// OfferingUpdateOfferingByIDHandler sets the operation handler for the update offering by Id operation
	OfferingUpdateOfferingByIDHandler offering.UpdateOfferingByIDHandler
	// ProfileUpdateProfileByIDHandler sets the operation handler for the update profile by Id operation
	ProfileUpdateProfileByIDHandler profile.UpdateProfileByIDHandler

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

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.MessageCreateNewMessageHandler == nil {
		unregistered = append(unregistered, "message.CreateNewMessageHandler")
	}

	if o.OfferingCreateNewOfferingHandler == nil {
		unregistered = append(unregistered, "offering.CreateNewOfferingHandler")
	}

	if o.ProfileCreateNewProfileHandler == nil {
		unregistered = append(unregistered, "profile.CreateNewProfileHandler")
	}

	if o.MessageDeleteMessageByIDHandler == nil {
		unregistered = append(unregistered, "message.DeleteMessageByIDHandler")
	}

	if o.OfferingDeleteOfferingByIDHandler == nil {
		unregistered = append(unregistered, "offering.DeleteOfferingByIDHandler")
	}

	if o.ProfileDeleteProfileByIDHandler == nil {
		unregistered = append(unregistered, "profile.DeleteProfileByIDHandler")
	}

	if o.FeedGetFeedHandler == nil {
		unregistered = append(unregistered, "feed.GetFeedHandler")
	}

	if o.MessageGetMessageByIDHandler == nil {
		unregistered = append(unregistered, "message.GetMessageByIDHandler")
	}

	if o.OfferingGetOfferingByIDHandler == nil {
		unregistered = append(unregistered, "offering.GetOfferingByIDHandler")
	}

	if o.ProfileGetProfileByIDHandler == nil {
		unregistered = append(unregistered, "profile.GetProfileByIDHandler")
	}

	if o.HealthHealthcheckGetHandler == nil {
		unregistered = append(unregistered, "health.HealthcheckGetHandler")
	}

	if o.MessageUpdateMessageByIDHandler == nil {
		unregistered = append(unregistered, "message.UpdateMessageByIDHandler")
	}

	if o.OfferingUpdateOfferingByIDHandler == nil {
		unregistered = append(unregistered, "offering.UpdateOfferingByIDHandler")
	}

	if o.ProfileUpdateProfileByIDHandler == nil {
		unregistered = append(unregistered, "profile.UpdateProfileByIDHandler")
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

	return nil

}

// Authorizer returns the registered authorizer
func (o *SavoodAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SavoodAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

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

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/message"] = message.NewCreateNewMessage(o.context, o.MessageCreateNewMessageHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/offering"] = offering.NewCreateNewOffering(o.context, o.OfferingCreateNewOfferingHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/profile"] = profile.NewCreateNewProfile(o.context, o.ProfileCreateNewProfileHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/message/{id}"] = message.NewDeleteMessageByID(o.context, o.MessageDeleteMessageByIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/offering/{id}"] = offering.NewDeleteOfferingByID(o.context, o.OfferingDeleteOfferingByIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/profile/{id}"] = profile.NewDeleteProfileByID(o.context, o.ProfileDeleteProfileByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/feed"] = feed.NewGetFeed(o.context, o.FeedGetFeedHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/message/{id}"] = message.NewGetMessageByID(o.context, o.MessageGetMessageByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offering/{id}"] = offering.NewGetOfferingByID(o.context, o.OfferingGetOfferingByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile/{id}"] = profile.NewGetProfileByID(o.context, o.ProfileGetProfileByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthcheck"] = health.NewHealthcheckGet(o.context, o.HealthHealthcheckGetHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/message/{id}"] = message.NewUpdateMessageByID(o.context, o.MessageUpdateMessageByIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/offering/{id}"] = offering.NewUpdateOfferingByID(o.context, o.OfferingUpdateOfferingByIDHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/profile/{id}"] = profile.NewUpdateProfileByID(o.context, o.ProfileUpdateProfileByIDHandler)

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
