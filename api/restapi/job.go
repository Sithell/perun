package restapi

import (
	"github.com/go-openapi/runtime"
	"github.com/sithell/perun/api/models"
	"github.com/sithell/perun/api/restapi/operations"
	"github.com/sithell/perun/internal/database"
	"log"
	"net/http"
)

type createJobResponder struct {
	params operations.CreateJobParams
	app    *App
}

func (rs createJobResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	job := database.Job{Image: *rs.params.Job.Image, Command: rs.params.Job.Command}
	result := rs.app.DB.Save(&job)
	if result.Error != nil {
		log.Fatalf("failed to save job into database: %v", result.Error)
	}

	operations.NewCreateJobOK().WithPayload(&models.Job{
		ID:      int64(job.ID),
		Image:   &job.Image,
		Command: job.Command,
		Status:  "created",
	}).WriteResponse(rw, producer)
}
