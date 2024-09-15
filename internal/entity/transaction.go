package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	Client      *Client
	AccountTo   *Account
	AccountFrom *Account
	Amount      float64
	CreateAt    time.Time
	UpdateAt    time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		Client:      accountFrom.Client,
		AccountTo:   accountTo,
		AccountFrom: accountFrom,
		Amount:      amount,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	transaction.Commit(amount)
	return transaction, nil
}

func (t *Transaction) Commit(amount float64) {
	t.AccountFrom.Debit(amount)
	t.AccountTo.Credit(amount)
	t.UpdateAt = time.Now()
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if (t.AccountFrom.Balance - t.Amount) < 0 {
		return errors.New("insuficient funds")
	}

	return nil
}
