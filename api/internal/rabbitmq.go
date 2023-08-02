package internal

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitMQ(user, password, host string, port uint) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port))
}
