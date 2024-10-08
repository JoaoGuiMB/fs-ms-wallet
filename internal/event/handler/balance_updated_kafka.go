package handler

import (
	"fmt"
	"sync"

	"github.com.br/joaoguimb/fc-ms-wallet/pkg/events"
	"github.com.br/joaoguimb/fc-ms-wallet/pkg/kafka"
)

type BalanceUpdatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *BalanceUpdatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	err := h.Kafka.Publish(message, nil, "balances")
	fmt.Println("BalanceUpdatedKafkaHandler: ", message)
	if err != nil {
		panic(err)
	}
}
