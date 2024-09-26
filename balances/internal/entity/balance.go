package entity

type Balance struct {
	AccountID string
	Amount    int
}

func NewBalance(accountID string, amount int) *Balance {
	return &Balance{
		AccountID: accountID,
		Amount:    amount,
	}
}

func (b *Balance) SetAmount(amount int) *Balance {
	b.Amount = amount
	return b
}
