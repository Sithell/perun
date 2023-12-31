// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/loads"
	flag "github.com/spf13/pflag"

	"github.com/sithell/perun/backend/api/restapi"
	"github.com/sithell/perun/backend/api/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  api-server [OPTIONS]\n\n")

		title := "Perun API"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := "Public Perun API for users"
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	flag.Parse()

	api := operations.NewAPIAPI(swaggerSpec)
	// get server with flag values filled out
	server = restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
