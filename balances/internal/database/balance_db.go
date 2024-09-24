package database

import (
	"database/sql"

	"github.com.br/joaoguimb/fc-ms-wallet/balances/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (db *BalanceDB) Get(accountID string) (*entity.Balance, error) {
	var balance entity.Balance
	stmt, err := db.DB.Prepare("SELECT account_id, amount FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(accountID).Scan(&balance.AccountID, &balance.Amount)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (db *BalanceDB) Update(balance entity.Balance) error {
	stmt, err := db.DB.Prepare("UPDATE balances SET amount = ? WHERE account_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(balance.Amount, balance.AccountID)
	if err != nil {
		return err
	}

	return nil
}
