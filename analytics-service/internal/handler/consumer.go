package handler

import (
	"encoding/json"
	"log"
	"os"

	

	amqp "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	ch        *amqp.Channel
	queueName string
	workers   int
	uc        *use.MealUsecase
}

// DI constructor
func NewHandlerConsumer(
	ch *amqp.Channel,
	uc *usecase.MealUsecase,
) *Handler {

	return &Handler{
		ch:        ch,
		queueName: os.Getenv("QUEUE_NAME"),
		workers:   10,
		uc:        uc,
	}
}

// start consumer
func (h *Handler) Start() {

	msgs, err := h.ch.Consume(
		h.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("consume error:", err)
	}

	for i := 0; i < h.workers; i++ {
		go h.worker(msgs, i)
	}

	select {}
}

// worker
func (h *Handler) worker(
	msgs <-chan amqp.Delivery,
	id int,
) {

	for msg := range msgs {

		var meal models.Meal

		// json parse
		err := json.Unmarshal(msg.Body, &meal)
		if err != nil {

			log.Println("json parse error:", err)

			msg.Nack(false, false)
			continue
		}

		// create meal
		err = h.uc.Create(&meal)
		if err != nil {

			log.Println("create meal error:", err)

			msg.Nack(false, true)
			continue
		}

		log.Printf(
			"[worker:%d] meal created: %s",
			id,
			meal.Name,
		)

		msg.Ack(false)
	}
}