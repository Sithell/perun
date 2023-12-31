// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetJobByIDHandlerFunc turns a function with the right signature into a get job by ID handler
type GetJobByIDHandlerFunc func(GetJobByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetJobByIDHandlerFunc) Handle(params GetJobByIDParams) middleware.Responder {
	return fn(params)
}

// GetJobByIDHandler interface for that can handle valid get job by ID params
type GetJobByIDHandler interface {
	Handle(GetJobByIDParams) middleware.Responder
}

// NewGetJobByID creates a new http.Handler for the get job by ID operation
func NewGetJobByID(ctx *middleware.Context, handler GetJobByIDHandler) *GetJobByID {
	return &GetJobByID{Context: ctx, Handler: handler}
}

/*
	GetJobByID swagger:route GET /jobs/{id} getJobById

# Get a job by id

Returns a job with corresponding id
*/
type GetJobByID struct {
	Context *middleware.Context
	Handler GetJobByIDHandler
}

func (o *GetJobByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetJobByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
