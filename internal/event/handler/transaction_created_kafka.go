package handler

import (
	"fmt"
	"sync"

	"github.com.br/joaoguimb/fc-ms-wallet/pkg/events"
	"github.com.br/joaoguimb/fc-ms-wallet/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("adsad")
	err := h.Kafka.Publish(message, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler: ", message.GetPayload())
	if err != nil {
		panic(err)
	}
}
