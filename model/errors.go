package model

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
	ErrUnknownType = func(msg string) error {
		return fmt.Errorf("unknown type: %s", msg)
	}
)
