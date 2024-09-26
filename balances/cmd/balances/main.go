package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/joaoguimb/fc-ms-wallet/balances/internal/database"
	getbalance "github.com.br/joaoguimb/fc-ms-wallet/balances/internal/usecase/get_balance"
	updatebalance "github.com.br/joaoguimb/fc-ms-wallet/balances/internal/usecase/update_balance"
	"github.com.br/joaoguimb/fc-ms-wallet/balances/internal/web"
	"github.com.br/joaoguimb/fc-ms-wallet/balances/internal/web/webserver"
	"github.com.br/joaoguimb/fc-ms-wallet/balances/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

type MessagePayload struct {
	AccountIDFrom      string `json:"account_id_from"`
	AccountIDTo        string `json:"account_id_to"`
	BalanceAccountFrom int    `json:"balance_account_from"`
	BalanceAccountTo   int    `json:"balance_account_to"`
}

type TestMessage struct {
	Name    string
	Payload MessagePayload
}

func consumer() {
	fmt.Println("running consumer")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/balances?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balancesDb := database.NewBalanceDB(db)

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

	updateBalanceUseCase := updatebalance.NewUpdateBalanceUseCase(balancesDb)

	input := &updatebalance.UpdateBalanceInputDTO{
		AccountIDFrom:        testMessage.Payload.AccountIDFrom,
		AccountIDTo:          testMessage.Payload.AccountIDTo,
		BalanceAccountIDFrom: testMessage.Payload.BalanceAccountFrom,
		BalanceAccountIDTo:   testMessage.Payload.BalanceAccountTo,
	}
	output, err := updateBalanceUseCase.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

}

func main() {
	fmt.Println("Starting application")

	// Set up the database connection
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/balances?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balancesDb := database.NewBalanceDB(db)

	// Set up the web server
	getBalanceUseCase := getbalance.NewGetBalanceUseCase(balancesDb)
	webBalanceHandler := web.NewWebBalanceHanlder(*getBalanceUseCase)
	webserver := webserver.NewWebServer(":3003")

	// Start the web server in a goroutine
	go func() {
		webserver.AddHandler("/balances/", webBalanceHandler.GetBalance)
		webserver.Start()
		fmt.Println("Server started on port 3003")
	}()

	// Start the Kafka consumer
	consumer()

	// Prevent the main goroutine from exiting
	select {}
}
