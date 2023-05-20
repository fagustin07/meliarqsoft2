package notification

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"meliarqsoft2/internal/domain/model"
)

func (repo RabbitMQRepository) Send(notification *model.Notification) error {
	body, errJ := json.Marshal(notification)
	if errJ != nil {
		return errJ
	}

	msg := amqp.Publishing{
		DeliveryMode: 1,
		ContentType:  "application/json",
		Body:         []byte(body),
	}

	err := repo.channel.Publish("", repo.queue.Name, false, false, msg)

	return err
}
