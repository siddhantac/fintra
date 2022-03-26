package money

import (
	"encoding/json"
	"fmt"
)

type Money struct {
	intAmount   int
	floatAmount float32
}

// NewMoney creates and returns an instance of Money.
func NewMoney(amount float32) Money {
	return Money{
		intAmount:   int(amount * 100),
		floatAmount: amount,
	}
}

func (m Money) Amount() float32 {
	return m.floatAmount
}

func (m Money) AmountAsInt() int {
	return m.intAmount
}

func (m Money) Add(money Money) Money {
	return newMoneyFromInt(m.intAmount + money.intAmount)
}

func (m Money) Subtract(money Money) Money {
	return newMoneyFromInt(m.intAmount - money.intAmount)
}

func (m Money) String() string {
	return fmt.Sprintf("%v", m.floatAmount)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.floatAmount)
}

func newMoneyFromInt(amount int) Money {
	return Money{
		intAmount:   amount,
		floatAmount: float32(amount) / 100.0,
	}
}
