// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/rs/cors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"feather-ai/service-core/api/generated/restapi/operations"
	"feather-ai/service-core/lib/config"
)

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logrus.Info("In ExampleMiddleware")
		wrappedRequest := r.WithContext(
			config.CreateRequestContext(
				r.Context(),
				&config.RequestContext{
					UserID: uuid.Nil,
				}))

		next.ServeHTTP(w, wrappedRequest)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.RequestURI) != "/v1/health" {

			headers := ""
			if reqHeadersBytes, err := json.Marshal(r.Header); err == nil {
				headers = string(reqHeadersBytes)
			}

			logrus.WithField("URL", r.RequestURI).WithField("Headers", headers).Info("New Request")
		}
		next.ServeHTTP(w, r)
	})
}

//go:generate swagger generate server --target ../../generated --name Core --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.CoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return exampleMiddleware(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
			"http://localhost:3000",
			"https://feather-ai.com",
			"https://www.feather-ai.com",
			"https://petstore.swagger.io", // Used for Swagger UI
		},
		AllowCredentials:   true,
		AllowedMethods:     []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{},
		Debug:              false,
		OptionsPassthrough: false,
	})

	return loggerMiddleware(c.Handler(handler))
}
