package handlers

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventoInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Ordem criada: %v", event.GetValues())
	jsonOutput, _ := json.Marshal(event.GetValues())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
