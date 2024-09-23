package create_transaction

import (
	"context"
	"testing"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/event"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/usecase/mocks"
	"github.com.br/joaoguimb/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john@doe.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Jane Doe", "jane@doe.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	accountGatewayMock := &mocks.AccountGatewayMock{}
	accountGatewayMock.On("FindByID", account1.ID).Return(account1, nil)
	accountGatewayMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionGatewayMock := &mocks.TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)

	input := &CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	output, err := uc.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
