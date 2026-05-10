package rabbitmqproducer

import (
	"encoding/json"
	"mail-service/internal/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMeal(ch *amqp.Channel, meal *models.Meal) error {

	body, err := json.Marshal(meal)
	if err != nil {
		return err
	}

	return ch.Publish(
		"statik_exchange",
		"statik.success",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}