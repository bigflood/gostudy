// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/bigflood/gostudy/swagger/models"
	"github.com/bigflood/gostudy/swagger/restapi/operations"
	"github.com/bigflood/gostudy/swagger/restapi/operations/admins"
	"github.com/bigflood/gostudy/swagger/restapi/operations/developers"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
)

//go:generate swagger generate server --target ../../swagger --name SimpleInventory --spec ../swagger.yaml

var inventoryItems = make(map[string]*models.InventoryItem)

func configureFlags(api *operations.SimpleInventoryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SimpleInventoryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AdminsAddInventoryHandler = admins.AddInventoryHandlerFunc(func(params admins.AddInventoryParams) middleware.Responder {
		if _, ok := inventoryItems[*params.InventoryItem.ID]; ok {
			return admins.NewAddInventoryConflict()
		}

		inventoryItems[*params.InventoryItem.ID] = params.InventoryItem

		return admins.NewAddInventoryCreated()
	})
	api.DevelopersSearchInventoryHandler = developers.SearchInventoryHandlerFunc(func(params developers.SearchInventoryParams) middleware.Responder {
		items := make([]*models.InventoryItem, 0, len(inventoryItems))

		for _, item := range inventoryItems {
			items = append(items, item)
		}

		return developers.NewSearchInventoryOK().WithPayload(items)
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
	return handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Referer", "User-Agent", "Accept-Encoding"}),
	)(handler)
}
