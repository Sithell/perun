package restapi

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/sithell/perun/api/models"
	"github.com/sithell/perun/api/restapi/operations"
	"github.com/sithell/perun/internal/database"
	"gorm.io/gorm"
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

type getJobByIDResponder struct {
	params operations.GetJobByIDParams
	app    *App
}

func (rs getJobByIDResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	job := database.Job{}
	result := rs.app.DB.First(&job, rs.params.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			operations.NewGetJobByIDNotFound().WithPayload(&models.Error{
				Message: fmt.Sprintf("No job with id=%d", rs.params.ID),
			}).WriteResponse(rw, producer)
			return
		}
		log.Fatalf("failed to find job with id=%d in database: %v", &rs.params.ID, result.Error)
	}

	operations.NewCreateJobOK().WithPayload(&models.Job{
		ID:      int64(job.ID),
		Image:   &job.Image,
		Command: job.Command,
		Status:  "created",
	}).WriteResponse(rw, producer)
}
