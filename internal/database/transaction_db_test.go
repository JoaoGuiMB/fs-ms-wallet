package database

import (
	"database/sql"
	"testing"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountTo     *entity.Account
	accountFrom   *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	s.transactionDB = NewTransactionDB(db)
	s.client, _ = entity.NewClient("John Doe", "john@doe.com")
	s.client2, _ = entity.NewClient("John Doe", "john@doe.com")
	s.accountTo = entity.NewAccount(s.client)
	s.accountFrom = entity.NewAccount(s.client2)
	s.accountFrom.Balance = 1000
	s.accountTo.Balance = 100
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, _ := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	err := s.transactionDB.Create(transaction)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.NotNil(err)
}
