package entity

import "time"

type Account struct {
	ID       string
	Client   *Client
	Balance  float64
	CreateAt time.Time
	UpdateAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	return &Account{
		ID:       client.ID,
		Client:   client,
		Balance:  0,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.UpdateAt = time.Now()
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.UpdateAt = time.Now()
}
