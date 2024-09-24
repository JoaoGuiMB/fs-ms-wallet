package entity

type Balance struct {
	AccountID string
	Amount    float64
}

func NewBalance(accountID string, amount float64) *Balance {
	return &Balance{
		AccountID: accountID,
		Amount:    amount,
	}
}

func (b *Balance) SetAmount(amount float64) *Balance {
	b.Amount = amount
	return b
}
