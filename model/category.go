package model

import (
	"time"

	"github.com/siddhantac/fintra/money"
)

type Category struct {
	Name    string
	Budget  money.Money
	Spent   money.Money
	Created time.Time
	Updated time.Time
}

func NewCategory(name string) *Category {
	now := time.Now()
	return &Category{
		Name:    name,
		Created: now,
		Updated: now,
	}
}

// TODO for future use
func (c *Category) SetBudget(budget money.Money) {
	c.Budget = budget
}

func (c *Category) RecordExpense(amount money.Money) {
	c.Spent = c.Spent.Add(amount)
}

func (c *Category) RemainingBudget() money.Money {
	return c.Budget.Subtract(c.Spent)
}
