package balances

import (
	"database/sql"

	"github.com.br/joaoguimb/fc-ms-wallet/balances/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/balances?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "balances",
	}

	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})

	kafkaConsumer.Consume(make(chan *ckafka.Message))

}
