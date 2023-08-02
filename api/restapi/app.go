package restapi

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
	MQ *amqp.Connection
}
