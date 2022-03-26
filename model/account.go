package model

import (
	"time"

	"github.com/siddhantac/fintra/money"
)

type Account struct {
	Created time.Time
	Updated time.Time

	Name            string
	StartingBalance money.Money
	CurrentBalance  money.Money
}

func NewAccount(name string, startingBalance money.Money) *Account {
	now := time.Now()
	return &Account{
		Name:            name,
		StartingBalance: startingBalance,
		CurrentBalance:  startingBalance,
		Created:         now,
		Updated:         now,
	}
}

func (a *Account) Credit(amount money.Money) {
	a.CurrentBalance = a.CurrentBalance.Add(amount)
}

func (a *Account) Debit(amount money.Money) {
	a.CurrentBalance = a.CurrentBalance.Subtract(amount)
}
