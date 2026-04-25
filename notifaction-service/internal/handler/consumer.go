package handler

import (
	"log"
	"notifaction-service/internal/domain"

	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	ch        *amqp.Channel
	queueName string
	workers   int
	uc        domain.SMSUsecase
}

// DI constructor
func NewHandlerConsumer(ch *amqp.Channel, uc domain.SMSUsecase) *Handler {
	return &Handler{
		ch:        ch,
		queueName: os.Getenv("QUEUE_NAME"),
		workers:   10,
		uc:        uc,
	}
}

// start consuming
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

	select {} // block
}

// worker
func (h *Handler) worker(msgs <-chan amqp.Delivery, id int) {
	for msg := range msgs {

		err := h.uc.Handle(msg.Body)
		if err != nil {
		
			msg.Nack(false, true)
			continue
		}



		msg.Ack(false)
	}
}g