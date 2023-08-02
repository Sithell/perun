package restapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sithell/perun/api/models"
	"github.com/sithell/perun/api/restapi/operations"
	"github.com/sithell/perun/internal/database"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
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
	ch, err := rs.app.MQ.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Printf("failed to close channel: %v", err)
		}
	}(ch)

	q, err := ch.QueueDeclare(
		"run-container",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	job := database.Job{Image: *rs.params.Job.Image, Command: rs.params.Job.Command}
	result := rs.app.DB.Save(&job)
	if result.Error != nil {
		log.Fatalf("failed to save job into database: %v", result.Error)
	}

	apiJob := dbModelToApiModel(job, nil)
	body, err := json.Marshal(apiJob)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("failed to publish a message: %v", err)
	}
	log.Printf(" [x] Sent %s\n", string(body))

	operations.NewCreateJobOK().WithPayload(apiJob).WriteResponse(rw, producer)
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

type getJobStdoutResponder struct {
	params operations.GetJobStdoutParams
	app    *App
}

func (rs getJobStdoutResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("WARN: failed to fetch runs for job #%d: %v", job.ID, result.Error)
		}
	}
	operations.NewGetJobStdoutOK().WithPayload(run.Stdout).WriteResponse(rw, producer)
}
