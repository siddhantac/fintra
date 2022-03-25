package model

import "time"

type Account struct {
	Created time.Time
	Updated time.Time

	Name            string
	StartingBalance int
	CurrentBalance  int
}

func NewAccount(name string, startingBalance int) *Account {
	now := time.Now()
	return &Account{
		Name:            name,
		StartingBalance: startingBalance,
		CurrentBalance:  startingBalance,
		Created:         now,
		Updated:         now,
	}
}

/* func (a *Account) Credit(amount int) int {
	a.currentBalance += amount
	return a.CurrentBalance()
}

func (a *Account) Debit(amount int) int {
	a.currentBalance -= amount
	return a.CurrentBalance()
} */
