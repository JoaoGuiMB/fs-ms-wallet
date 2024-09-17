package database

import (
	"database/sql"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (db *AccountDB) FindByID(id int64) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := db.DB.Prepare("SELECT a.id, a.client_id, a.created_at, c.id, c.name, c.email, c.created_at FROM accounts a JOIN clients c ON a.client_id = c.id WHERE a.id = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&account.ID, &account.Client.ID, &account.CreatedAt, &client.ID, &client.Name, &client.Email, &client.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (db *AccountDB) Save(account *entity.Account) error {
	stmt, err := db.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Client.ID, account.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
