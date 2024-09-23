package event

import "time"

type BalanceUpdated struct {
	Name    string
	Payload interface{}
}

func NewBalanceUpdated() *BalanceUpdated {
	return &BalanceUpdated{
		Name: "balance_updated",
	}
}

func (e *BalanceUpdated) GetName() string {
	return e.Name
}

func (e *BalanceUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *BalanceUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *BalanceUpdated) GetDateTime() time.Time {
	return time.Now()
}
