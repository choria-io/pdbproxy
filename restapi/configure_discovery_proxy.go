package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/choria-io/discovery_proxy/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name discovery_proxy --spec ../schema.yaml

func configureFlags(api *operations.DiscoveryProxyAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DiscoveryProxyAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BinProducer = runtime.ByteStreamProducer()

	api.DeleteSetSetHandler = operations.DeleteSetSetHandlerFunc(func(params operations.DeleteSetSetParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteSetSet has not yet been implemented")
	})
	api.GetBackupHandler = operations.GetBackupHandlerFunc(func(params operations.GetBackupParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBackup has not yet been implemented")
	})
	api.GetDiscoverHandler = operations.GetDiscoverHandlerFunc(func(params operations.GetDiscoverParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetDiscover has not yet been implemented")
	})
	api.GetSetSetHandler = operations.GetSetSetHandlerFunc(func(params operations.GetSetSetParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetSetSet has not yet been implemented")
	})
	api.GetSetsHandler = operations.GetSetsHandlerFunc(func(params operations.GetSetsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetSets has not yet been implemented")
	})
	api.PostSetHandler = operations.PostSetHandlerFunc(func(params operations.PostSetParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostSet has not yet been implemented")
	})
	api.PutSetSetHandler = operations.PutSetSetHandlerFunc(func(params operations.PutSetSetParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutSetSet has not yet been implemented")
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
