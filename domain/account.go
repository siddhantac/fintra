package domain

import "time"

type Account struct {
	Created time.Time
	Updated time.Time

	name            string
	startingBalance int
	currentBalance  int
}

func NewAccount(name string, startingBalance int) *Account {
	now := time.Now()
	return &Account{
		name:            name,
		startingBalance: startingBalance,
		currentBalance:  startingBalance,
		Created:         now,
		Updated:         now,
	}
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) StartingBalance() int {
	return a.startingBalance
}

func (a *Account) CurrentBalance() int {
	return a.currentBalance
}

func (a *Account) Credit(amount int) int {
	a.currentBalance += amount
	return a.CurrentBalance()
}

func (a *Account) Debit(amount int) int {
	a.currentBalance -= amount
	return a.CurrentBalance()
}
