package create_transaction

import (
	"context"
	"fmt"

	"github.com.br/joaoguimb/fc-ms-wallet/internal/entity"
	"github.com.br/joaoguimb/fc-ms-wallet/internal/gateway"
	"github.com.br/joaoguimb/fc-ms-wallet/pkg/events"
	"github.com.br/joaoguimb/fc-ms-wallet/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_to"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispacther    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
	BalanceUpdated     events.EventInterface
}

func NewCreateTransactionUseCase(uow uow.UowInterface, eventDispacther events.EventDispatcherInterface, transactionCreated events.EventInterface, balanceUpdated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                uow,
		EventDispacther:    eventDispacther,
		TransactionCreated: transactionCreated,
		BalanceUpdated:     balanceUpdated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {

	output := &CreateTransactionOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}

	err := uc.Uow.Do(context.Background(), func(uow *uow.Uow) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)

		if err != nil {
			return err
		}

		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

		if err != nil {
			return err

		}

		err = accountRepository.UpdateBalance(accountFrom)

		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)

		if err != nil {
			fmt.Println(err)
			return err
		}
		output.TransactionID = transaction.ID

		balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
		balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
		balanceUpdatedOutput.BalanceAccountIDFrom = accountFrom.Balance
		balanceUpdatedOutput.BalanceAccountIDTo = accountTo.Balance

		return nil
	})

	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispacther.Dispatch(uc.TransactionCreated)

	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	uc.EventDispacther.Dispatch(uc.BalanceUpdated)

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repository, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repository.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repository, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repository.(gateway.TransactionGateway)
}
