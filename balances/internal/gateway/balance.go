package gateway

import "github.com.br/joaoguimb/fc-ms-wallet/balances/internal/entity"

type BalanceGateway interface {
	Get(accountID string) (*entity.Balance, error)
	Update(balance entity.Balance) error
}
