package database

import (
	"database/sql"
	"testing"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	s.clientDB.Save(client)

	result, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, result.ID)
	s.Equal(client.Name, result.Name)
	s.Equal(client.Email, result.Email)
	s.NotEmpty(result.CreatedAt)
}

func (s *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	err := s.clientDB.Save(client)
	s.Nil(err)
	result, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, result.ID)
	s.Equal(client.Name, result.Name)
	s.Equal(client.Email, result.Email)
	s.NotEmpty(result.CreatedAt)
}
