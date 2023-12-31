// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCoreAPI creates a new Core instance
func NewCoreAPI(spec *loads.Document) *CoreAPI {
	return &CoreAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		DeleteV1ClientApikeyHandler: DeleteV1ClientApikeyHandlerFunc(func(params DeleteV1ClientApikeyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DeleteV1ClientApikey has not yet been implemented")
		}),
		GetV1ClientApikeyHandler: GetV1ClientApikeyHandlerFunc(func(params GetV1ClientApikeyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetV1ClientApikey has not yet been implemented")
		}),
		GetV1ClientUploadsHandler: GetV1ClientUploadsHandlerFunc(func(params GetV1ClientUploadsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetV1ClientUploads has not yet been implemented")
		}),
		GetV1DebugExecuteRequestSchemaHandler: GetV1DebugExecuteRequestSchemaHandlerFunc(func(params GetV1DebugExecuteRequestSchemaParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetV1DebugExecuteRequestSchema has not yet been implemented")
		}),
		GetV1HealthHandler: GetV1HealthHandlerFunc(func(params GetV1HealthParams) middleware.Responder {
			return middleware.NotImplemented("operation GetV1Health has not yet been implemented")
		}),
		GetV1PublicSystemHandler: GetV1PublicSystemHandlerFunc(func(params GetV1PublicSystemParams) middleware.Responder {
			return middleware.NotImplemented("operation GetV1PublicSystem has not yet been implemented")
		}),
		GetV1PublicSystemSystemIDHandler: GetV1PublicSystemSystemIDHandlerFunc(func(params GetV1PublicSystemSystemIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetV1PublicSystemSystemID has not yet been implemented")
		}),
		GetV1PublicSystemsHandler: GetV1PublicSystemsHandlerFunc(func(params GetV1PublicSystemsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetV1PublicSystems has not yet been implemented")
		}),
		GetV1PublicUserUserNameHandler: GetV1PublicUserUserNameHandlerFunc(func(params GetV1PublicUserUserNameParams) middleware.Responder {
			return middleware.NotImplemented("operation GetV1PublicUserUserName has not yet been implemented")
		}),
		PutV1APISystemCompletePublishHandler: PutV1APISystemCompletePublishHandlerFunc(func(params PutV1APISystemCompletePublishParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PutV1APISystemCompletePublish has not yet been implemented")
		}),
		PutV1APISystemPreparePublishHandler: PutV1APISystemPreparePublishHandlerFunc(func(params PutV1APISystemPreparePublishParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PutV1APISystemPreparePublish has not yet been implemented")
		}),
		PutV1ClientApikeyHandler: PutV1ClientApikeyHandlerFunc(func(params PutV1ClientApikeyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PutV1ClientApikey has not yet been implemented")
		}),
		PutV1ClientLoginHandler: PutV1ClientLoginHandlerFunc(func(params PutV1ClientLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation PutV1ClientLogin has not yet been implemented")
		}),
		PutV1ClientRefreshHandler: PutV1ClientRefreshHandlerFunc(func(params PutV1ClientRefreshParams) middleware.Responder {
			return middleware.NotImplemented("operation PutV1ClientRefresh has not yet been implemented")
		}),
		PutV1PublicSystemSystemIDStepStepIndexHandler: PutV1PublicSystemSystemIDStepStepIndexHandlerFunc(func(params PutV1PublicSystemSystemIDStepStepIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation PutV1PublicSystemSystemIDStepStepIndex has not yet been implemented")
		}),

		// Applies when the "X-FEATHER-API-KEY" header is set
		APIKeyAuthAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (ApiKeyAuth) X-FEATHER-API-KEY from header param [X-FEATHER-API-KEY] has not yet been implemented")
		},
		// Applies when the "X-FEATHER-TOKEN" header is set
		FeatherTokenAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (FeatherToken) X-FEATHER-TOKEN from header param [X-FEATHER-TOKEN] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*CoreAPI the core API */
type CoreAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIKeyAuthAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-FEATHER-API-KEY provided in the header
	APIKeyAuthAuth func(string) (interface{}, error)

	// FeatherTokenAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-FEATHER-TOKEN provided in the header
	FeatherTokenAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// DeleteV1ClientApikeyHandler sets the operation handler for the delete v1 client apikey operation
	DeleteV1ClientApikeyHandler DeleteV1ClientApikeyHandler
	// GetV1ClientApikeyHandler sets the operation handler for the get v1 client apikey operation
	GetV1ClientApikeyHandler GetV1ClientApikeyHandler
	// GetV1ClientUploadsHandler sets the operation handler for the get v1 client uploads operation
	GetV1ClientUploadsHandler GetV1ClientUploadsHandler
	// GetV1DebugExecuteRequestSchemaHandler sets the operation handler for the get v1 debug execute request schema operation
	GetV1DebugExecuteRequestSchemaHandler GetV1DebugExecuteRequestSchemaHandler
	// GetV1HealthHandler sets the operation handler for the get v1 health operation
	GetV1HealthHandler GetV1HealthHandler
	// GetV1PublicSystemHandler sets the operation handler for the get v1 public system operation
	GetV1PublicSystemHandler GetV1PublicSystemHandler
	// GetV1PublicSystemSystemIDHandler sets the operation handler for the get v1 public system system ID operation
	GetV1PublicSystemSystemIDHandler GetV1PublicSystemSystemIDHandler
	// GetV1PublicSystemsHandler sets the operation handler for the get v1 public systems operation
	GetV1PublicSystemsHandler GetV1PublicSystemsHandler
	// GetV1PublicUserUserNameHandler sets the operation handler for the get v1 public user user name operation
	GetV1PublicUserUserNameHandler GetV1PublicUserUserNameHandler
	// PutV1APISystemCompletePublishHandler sets the operation handler for the put v1 API system complete publish operation
	PutV1APISystemCompletePublishHandler PutV1APISystemCompletePublishHandler
	// PutV1APISystemPreparePublishHandler sets the operation handler for the put v1 API system prepare publish operation
	PutV1APISystemPreparePublishHandler PutV1APISystemPreparePublishHandler
	// PutV1ClientApikeyHandler sets the operation handler for the put v1 client apikey operation
	PutV1ClientApikeyHandler PutV1ClientApikeyHandler
	// PutV1ClientLoginHandler sets the operation handler for the put v1 client login operation
	PutV1ClientLoginHandler PutV1ClientLoginHandler
	// PutV1ClientRefreshHandler sets the operation handler for the put v1 client refresh operation
	PutV1ClientRefreshHandler PutV1ClientRefreshHandler
	// PutV1PublicSystemSystemIDStepStepIndexHandler sets the operation handler for the put v1 public system system ID step step index operation
	PutV1PublicSystemSystemIDStepStepIndexHandler PutV1PublicSystemSystemIDStepStepIndexHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *CoreAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *CoreAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *CoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CoreAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CoreAPI
func (o *CoreAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuthAuth == nil {
		unregistered = append(unregistered, "XFEATHERAPIKEYAuth")
	}
	if o.FeatherTokenAuth == nil {
		unregistered = append(unregistered, "XFEATHERTOKENAuth")
	}

	if o.DeleteV1ClientApikeyHandler == nil {
		unregistered = append(unregistered, "DeleteV1ClientApikeyHandler")
	}
	if o.GetV1ClientApikeyHandler == nil {
		unregistered = append(unregistered, "GetV1ClientApikeyHandler")
	}
	if o.GetV1ClientUploadsHandler == nil {
		unregistered = append(unregistered, "GetV1ClientUploadsHandler")
	}
	if o.GetV1DebugExecuteRequestSchemaHandler == nil {
		unregistered = append(unregistered, "GetV1DebugExecuteRequestSchemaHandler")
	}
	if o.GetV1HealthHandler == nil {
		unregistered = append(unregistered, "GetV1HealthHandler")
	}
	if o.GetV1PublicSystemHandler == nil {
		unregistered = append(unregistered, "GetV1PublicSystemHandler")
	}
	if o.GetV1PublicSystemSystemIDHandler == nil {
		unregistered = append(unregistered, "GetV1PublicSystemSystemIDHandler")
	}
	if o.GetV1PublicSystemsHandler == nil {
		unregistered = append(unregistered, "GetV1PublicSystemsHandler")
	}
	if o.GetV1PublicUserUserNameHandler == nil {
		unregistered = append(unregistered, "GetV1PublicUserUserNameHandler")
	}
	if o.PutV1APISystemCompletePublishHandler == nil {
		unregistered = append(unregistered, "PutV1APISystemCompletePublishHandler")
	}
	if o.PutV1APISystemPreparePublishHandler == nil {
		unregistered = append(unregistered, "PutV1APISystemPreparePublishHandler")
	}
	if o.PutV1ClientApikeyHandler == nil {
		unregistered = append(unregistered, "PutV1ClientApikeyHandler")
	}
	if o.PutV1ClientLoginHandler == nil {
		unregistered = append(unregistered, "PutV1ClientLoginHandler")
	}
	if o.PutV1ClientRefreshHandler == nil {
		unregistered = append(unregistered, "PutV1ClientRefreshHandler")
	}
	if o.PutV1PublicSystemSystemIDStepStepIndexHandler == nil {
		unregistered = append(unregistered, "PutV1PublicSystemSystemIDStepStepIndexHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "ApiKeyAuth":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuthAuth)

		case "FeatherToken":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.FeatherTokenAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *CoreAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CoreAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
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

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *CoreAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
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
func (o *CoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the core API
func (o *CoreAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CoreAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/client/apikey"] = NewDeleteV1ClientApikey(o.context, o.DeleteV1ClientApikeyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/client/apikey"] = NewGetV1ClientApikey(o.context, o.GetV1ClientApikeyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/client/uploads"] = NewGetV1ClientUploads(o.context, o.GetV1ClientUploadsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/debug/executeRequestSchema"] = NewGetV1DebugExecuteRequestSchema(o.context, o.GetV1DebugExecuteRequestSchemaHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/health"] = NewGetV1Health(o.context, o.GetV1HealthHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/public/system"] = NewGetV1PublicSystem(o.context, o.GetV1PublicSystemHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/public/system/{systemId}"] = NewGetV1PublicSystemSystemID(o.context, o.GetV1PublicSystemSystemIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/public/systems"] = NewGetV1PublicSystems(o.context, o.GetV1PublicSystemsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/public/user/{userName}"] = NewGetV1PublicUserUserName(o.context, o.GetV1PublicUserUserNameHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/api/system/completePublish"] = NewPutV1APISystemCompletePublish(o.context, o.PutV1APISystemCompletePublishHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/api/system/preparePublish"] = NewPutV1APISystemPreparePublish(o.context, o.PutV1APISystemPreparePublishHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/client/apikey"] = NewPutV1ClientApikey(o.context, o.PutV1ClientApikeyHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/client/login"] = NewPutV1ClientLogin(o.context, o.PutV1ClientLoginHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/client/refresh"] = NewPutV1ClientRefresh(o.context, o.PutV1ClientRefreshHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/public/system/{systemId}/step/{stepIndex}"] = NewPutV1PublicSystemSystemIDStepStepIndex(o.context, o.PutV1PublicSystemSystemIDStepStepIndexHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CoreAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CoreAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CoreAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CoreAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CoreAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
