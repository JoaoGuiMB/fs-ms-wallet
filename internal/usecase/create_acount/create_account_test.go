package create_account

import (
	"testing"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	clientGatewayMock := &mocks.ClientGatewayMock{}
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	account := entity.NewAccount(client)
	accountGatewayMock := &mocks.AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)

	input := &CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, account.ID, output.ID)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
