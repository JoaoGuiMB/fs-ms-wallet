package database

import (
	"database/sql"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John Doe", "john@doe.com")
}

func (s *AccountDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", "1", "John Doe", "john@doe.com", s.client.CreatedAt)
	s.db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, ?)", "1", "1", 100, s.client.CreatedAt)

	accountDb, err := s.accountDB.FindByID(1)
	s.Nil(err)
	s.Equal(accountDb.ID, "1")
	s.Equal(accountDb.Client.ID, "1")
	s.Equal(accountDb.Client.Name, "John Doe")
	s.Equal(accountDb.Client.Email, "john@doe.com")
	s.Equal(accountDb.Client.CreatedAt, s.client.CreatedAt)
	s.Equal(accountDb.Balance, 100)
	s.NotEmpty(accountDb.CreatedAt)
}
