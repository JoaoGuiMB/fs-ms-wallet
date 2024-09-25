package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/joaoguimb/fc-ms-wallet/balances/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

type MessagePayload struct {
	AccountIDFrom      string `json:"account_id_from"`
	AccountIDTo        string `json:"account_id_to"`
	BalanceAccountFrom int64  `json:"balance_account_from"`
	BalanceAccountTo   int64  `json:"balance_account_to"`
}

type TestMessage struct {
	Name    string
	Payload MessagePayload
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/balances?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "wallet",
	}

	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})

	messages := make(chan *ckafka.Message)
	go kafkaConsumer.Consume(messages)

	var testMessage TestMessage

	msg := <-messages
	fmt.Println(msg)
	value := msg.Value
	err = json.Unmarshal(value, &testMessage)
	if err != nil {
		panic(err)
	}
	fmt.Println("aqui")
	fmt.Println(testMessage.Payload)
}
