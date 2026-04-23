package rabbitmq

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"encoding/json"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(r *config.Rabbit) *Publisher {
	return &Publisher{
		ch:       r.Channel,
		exchange: os.Getenv("RABBITMQ_EXCHANGE"),
	}
}


func (p *Publisher) PublishAuthMessage(routingKey string, msg models.AuthMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.ch.Publish(
		p.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

