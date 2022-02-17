package domain

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

type Transaction struct {
	IntAmount   int
	Amount      float64
	ID          string
	Type        TransactionType
	Currency    Currency
	Description string
	Date        time.Time
	Category    TransactionCategory
	IsDebit     bool
	Account     string // TODO use strongly typed accounts
	Created     time.Time
	Updated     time.Time
}

func NewTransaction(amount float64, date time.Time, isDebit bool, category, transactionType, description, account string) (*Transaction, error) {
	now := time.Now()
	intAmount := int(amount * 100)

	tr := &Transaction{
		IntAmount:   intAmount,
		Amount:      amount,
		ID:          NewID(),
		Type:        TransactionType(transactionType),
		Category:    TransactionCategory(category),
		Date:        date,
		IsDebit:     isDebit,
		Description: description,
		Currency:    DefaultCurrency,
		Account:     account,
		Created:     now,
	}

	if err := tr.validate(); err != nil {
		return nil, err
	}

	return tr, nil
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%s\t%-10s\t%0.2f\t%s\t%s(%s)", t.Date.Format("02 Jan 2006"), t.Description, t.Amount, t.Account, t.Category, t.Type)

}

type TransactionType string

// TODO: type maybe unnecessary
const (
	TrTypeExpense    TransactionType = "expense"
	TrTypeIncome     TransactionType = "income"
	TrTypeTransfer   TransactionType = "transfer"
	TrTypeInvestment TransactionType = "investment"
)

var transactionTypes = map[TransactionType]struct{}{
	TrTypeExpense:    {},
	TrTypeIncome:     {},
	TrTypeTransfer:   {},
	TrTypeInvestment: {},
}

type Currency string

const (
	USD             Currency = "usd"
	SGD             Currency = "sgd"
	DefaultCurrency Currency = SGD
)

type TransactionCategory string

const (
	TrCategoryEntertainment TransactionCategory = "entertainment"
	TrCategoryGroceries     TransactionCategory = "groceries"
	TrCategoryHousehold     TransactionCategory = "household"
	TrCategoryInsurance     TransactionCategory = "insurance"
	TrCategoryMeals         TransactionCategory = "meals"
	TrCategoryMedical       TransactionCategory = "medical"
	TrCategoryOthers        TransactionCategory = "others"
	TrCategoryPersonal      TransactionCategory = "personal"
	TrCategoryRent          TransactionCategory = "rent"
	TrCategoryTax           TransactionCategory = "tax"
	TrCategoryTransport     TransactionCategory = "transport"
	TrCategoryTravel        TransactionCategory = "travel"
	TrCategoryUtilities     TransactionCategory = "utilities"
	TrCategoryRefund        TransactionCategory = "refund"
)

func (t *Transaction) validate() error {
	if err := validateType(t.Type, t.IsDebit); err != nil {
		return err
	}

	if t.Amount < 0 {
		return ErrIsNegative
	}

	if t.Currency == "" {
		return ErrEmpty("currency")
	}

	if t.Description == "" {
		return ErrEmpty("description")
	}

	if t.Date.IsZero() {
		return ErrEmpty("date")
	}

	if t.Category == "" {
		return ErrEmpty("category")
	}

	if t.Account == "" {
		return ErrEmpty("account")
	}

	return nil
}

func validateType(ty TransactionType, isDebit bool) error {
	if ty == "" {
		return ErrEmpty("transaction type")
	}

	if _, ok := transactionTypes[ty]; !ok {
		return fmt.Errorf("%s: %w", ty, ErrUnknownType)
	}

	switch ty {
	case TrTypeExpense, TrTypeInvestment:
		if !isDebit {
			return ErrMustBeDebit
		}
	case TrTypeIncome:
		if isDebit {
			return ErrMustBeCredit
		}
	}
	return nil
}

func NewID() string {
	return ksuid.New().String()
}
