package restapi

import (
	"github.com/go-openapi/runtime"
	"github.com/sithell/perun/api/models"
	"github.com/sithell/perun/api/restapi/operations"
	"net/http"
)

type createJobResponder struct {
	params operations.CreateJobParams
	app    *App
}

func (rs createJobResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	operations.NewCreateJobOK().WithPayload(&models.Job{
		Image: rs.params.Job.Image,
	}).WriteResponse(rw, producer)
}
