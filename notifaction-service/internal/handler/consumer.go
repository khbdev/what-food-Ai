package handler

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// =====================
// HANDLER = CONSUMER
// =====================

type Handler struct {
	ch        *amqp.Channel
	queueName string
	workers   int
}

// DI constructor
func NewHandlerConsumer(ch *amqp.Channel) *Handler {
	return &Handler{
		ch:        ch,
		queueName: os.Getenv("QUEUE_NAME"),
		workers:   10,
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

		log.Printf("[worker-%d] %s\n", id, string(msg.Body))

		msg.Ack(false)
	}
}