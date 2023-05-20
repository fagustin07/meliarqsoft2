package rabbitmq

import (
	"github.com/streadway/amqp"
	"os"
)

type Factory struct {
	mqClient *amqp.Connection
}

func NewFactory() *Factory {
	return &Factory{}
}

func (factory *Factory) InitRabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial(os.Getenv("AMQP_URI"))
	if err != nil {
		panic(err.Error())
	}

	return conn
}
