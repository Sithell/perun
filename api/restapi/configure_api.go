// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/sithell/perun/internal/database"
	flag "github.com/spf13/pflag"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/sithell/perun/api/restapi/operations"
)

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbPort     uint
	dbName     string
)

func init() {
	flag.StringVar(&dbHost, "db-host", "localhost", "database host")
	flag.StringVar(&dbUser, "db-user", "perun", "database user")
	flag.UintVar(&dbPort, "db-port", 5432, "database port")
	flag.StringVar(&dbName, "db-name", "perun", "database name")
	dbPassword = os.Getenv("DATABASE_PASSWORD")
}

var initAppFn = initApp

func initApp() (*App, error) {
	db, err := database.InitDB(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}
	return &App{DB: db}, nil
}

//goland:noinspection GoUnusedParameter
func configureFlags(api *operations.APIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.APIAPI) http.Handler {
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

	app, err := initAppFn()
	if err != nil {
		log.Fatalf("failed to init app: %v", app)
	}

	api.CreateJobHandler = operations.CreateJobHandlerFunc(func(params operations.CreateJobParams) middleware.Responder {
		return &createJobResponder{params: params, app: app}
	})

	if api.GetJobByIDHandler == nil {
		api.GetJobByIDHandler = operations.GetJobByIDHandlerFunc(func(params operations.GetJobByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetJobByID has not yet been implemented")
		})
	}
	if api.GetJobStdoutHandler == nil {
		api.GetJobStdoutHandler = operations.GetJobStdoutHandlerFunc(func(params operations.GetJobStdoutParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetJobStdout has not yet been implemented")
		})
	}
	if api.ListJobsHandler == nil {
		api.ListJobsHandler = operations.ListJobsHandlerFunc(func(params operations.ListJobsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListJobs has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
//
//goland:noinspection GoUnusedParameter
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
//
//goland:noinspection GoUnusedParameter
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
