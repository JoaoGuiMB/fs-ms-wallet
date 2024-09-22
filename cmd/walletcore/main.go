package main

import (
	"database/sql"
	"fmt"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/database"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/event"
	create_account "github.com.br/joaoguimb/fc-ms-wallet/internal/usecase/create_acount"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/usecase/create_client"
	createtransaction "github.com.br/joaoguimb/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/web"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/web/webserver"
	"github.com.br/joaoguimb/fc-ms-wallet/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallet?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()

	//eventDispatcher.Register("TransactionCreated", handler)

	cliendDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(cliendDB)

	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDB, cliendDB)

	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDB, accountDB, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":8080")

	webClientHandler := web.NewWebClientHandler(*createClientUseCase)
	webAccountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	webTransactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", webClientHandler.CreateClient)
	webserver.AddHandler("/accounts", webAccountHandler.CreateAccount)
	webserver.AddHandler("/transactions", webTransactionHandler.CreateTransaction)

	fmt.Println("Server started on port 8080")
	webserver.Start()
}
