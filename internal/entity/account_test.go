package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("JohnDoe", "john@doe.com")
	Account := NewAccount(client)
	assert.NotNil(t, Account)
	assert.Equal(t, client.ID, Account.ID)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("JohnDoe", "john@doe.com")
	Account := NewAccount(client)
	Account.Credit(100)
	assert.Equal(t, float64(100), Account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("JohnDoe", "john@doe.com")
	Account := NewAccount(client)
	Account.Credit(100)
	Account.Debit(50)
	assert.Equal(t, float64(50), Account.Balance)
}
