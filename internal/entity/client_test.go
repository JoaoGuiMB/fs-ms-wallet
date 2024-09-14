package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@doe.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
	assert.NotEmpty(t, client.CreateAt)
	assert.NotEmpty(t, client.UpdateAt)
}

func TestCreateNewClientArgsAreInvalid(t *testing.T) {
	_, err := NewClient("", "")
	assert.NotNil(t, err)
}
