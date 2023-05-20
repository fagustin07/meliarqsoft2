package notification

import (
	"github.com/streadway/amqp"
	"os"
)

type RabbitMQRepository struct {
	client  *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQRepository(client *amqp.Connection) *RabbitMQRepository {

	queueName := os.Getenv("AMQP_QUEUE_NAME")

	ch, errCh := client.Channel()

	if errCh != nil {
		panic(errCh.Error())
	}

	queue, errQ := ch.QueueDeclare(
		queueName, // name
		false,
		false,
		false,
		false,
		nil,
	)

	if errQ != nil {
		panic(errQ.Error())
	}

	repo := &RabbitMQRepository{
		client:  client,
		channel: ch,
		queue:   queue,
	}

	return repo
}
