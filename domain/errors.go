package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")

	ErrMustBeDebit  error = errors.New("expense or investment must be debit")
	ErrMustBeCredit error = errors.New("income must be credit")
	ErrIsNegative   error = errors.New("transaction amount cannot be negative")
	ErrEmpty              = func(s string) error {
		return fmt.Errorf("%s cannot be empty", s)
	}
	ErrUnknownTransactionType = func(s string) error {
		return fmt.Errorf("unknown transaction type: %s", s)
	}
)
