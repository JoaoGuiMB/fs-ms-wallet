package updatebalance

import "github.com.br/joaoguimb/fc-ms-wallet/balances/internal/gateway"

type UpdateBalanceInputDTO struct {
	AccountIDFrom        string `json:"account_id_from"`
	AccountIDTo          string `json:"account_id_to"`
	BalanceAccountIDFrom int    `json:"balance_account_id_from"`
	BalanceAccountIDTo   int    `json:"balance_account_id_to"`
}

type UpdateBalanceOutputDTO struct {
	AccountID string
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateBalanceUseCase(balanceGateway gateway.BalanceGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *UpdateBalanceUseCase) Execute(input *UpdateBalanceInputDTO) (*UpdateBalanceOutputDTO, error) {
	balanceFrom, err := uc.BalanceGateway.Get(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	balanceTo, err := uc.BalanceGateway.Get(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	balanceFrom.SetAmount(input.BalanceAccountIDFrom)

	// TODO: if update fails rollback
	err = uc.BalanceGateway.Update(*balanceFrom)
	if err != nil {
		return nil, err
	}

	balanceTo.SetAmount(input.BalanceAccountIDTo)
	err = uc.BalanceGateway.Update(*balanceTo)
	if err != nil {
		return nil, err
	}

	return &UpdateBalanceOutputDTO{
		AccountID: input.AccountIDFrom,
	}, nil
}
