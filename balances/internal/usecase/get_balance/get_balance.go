package getbalance

import "github.com.br/joaoguimb/fc-ms-wallet/balances/internal/gateway"

type GetBalanceInputDTO struct {
	AccountID string `json:"account_id"`
}

type GetBalanceOutputDTO struct {
	AccountID string
	Amount    int
}

type GetBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(balanceGateway gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *GetBalanceUseCase) Execute(input *GetBalanceInputDTO) (*GetBalanceOutputDTO, error) {
	balance, err := uc.BalanceGateway.Get(input.AccountID)
	if err != nil {
		return nil, err
	}

	return &GetBalanceOutputDTO{
		AccountID: input.AccountID,
		Amount:    balance.Amount,
	}, nil
}
