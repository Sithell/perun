// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RegisterProviderOKCode is the HTTP code returned for type RegisterProviderOK
const RegisterProviderOKCode int = 200

/*
RegisterProviderOK Provider successfully registered

swagger:response registerProviderOK
*/
type RegisterProviderOK struct {
}

// NewRegisterProviderOK creates RegisterProviderOK with default headers values
func NewRegisterProviderOK() *RegisterProviderOK {

	return &RegisterProviderOK{}
}

// WriteResponse to the client
func (o *RegisterProviderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RegisterProviderForbiddenCode is the HTTP code returned for type RegisterProviderForbidden
const RegisterProviderForbiddenCode int = 403

/*
RegisterProviderForbidden Failed to confirm provider reachability

swagger:response registerProviderForbidden
*/
type RegisterProviderForbidden struct {
}

// NewRegisterProviderForbidden creates RegisterProviderForbidden with default headers values
func NewRegisterProviderForbidden() *RegisterProviderForbidden {

	return &RegisterProviderForbidden{}
}

// WriteResponse to the client
func (o *RegisterProviderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
