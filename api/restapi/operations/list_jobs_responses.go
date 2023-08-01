// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sithell/perun/api/models"
)

// ListJobsOKCode is the HTTP code returned for type ListJobsOK
const ListJobsOKCode int = 200

/*
ListJobsOK Successfully fetched jobs

swagger:response listJobsOK
*/
type ListJobsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Job `json:"body,omitempty"`
}

// NewListJobsOK creates ListJobsOK with default headers values
func NewListJobsOK() *ListJobsOK {

	return &ListJobsOK{}
}

// WithPayload adds the payload to the list jobs o k response
func (o *ListJobsOK) WithPayload(payload []*models.Job) *ListJobsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list jobs o k response
func (o *ListJobsOK) SetPayload(payload []*models.Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListJobsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Job, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
