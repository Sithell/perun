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

func dbModelToApiModel(job database.Job, run *database.Run) *models.Job {
	result := models.Job{
		Command: job.Command,
		ID:      int64(job.ID),
		Image:   &job.Image,
		Status:  "created",
	}
	if run != nil {
		result.Status = run.Status
	}
	return &result
}

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

	operations.NewCreateJobOK().WithPayload(dbModelToApiModel(job, nil)).WriteResponse(rw, producer)
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
	run := &database.Run{}
	result = rs.app.DB.Where(database.Run{Job: job}).Order("created_at DESC").First(run)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			run = nil
		} else {
			log.Printf("WARN: failed to fetch runs for job #%d: %v", job.ID, result.Error)
		}
	}
	operations.NewCreateJobOK().WithPayload(dbModelToApiModel(job, run)).WriteResponse(rw, producer)
}
