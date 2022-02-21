package service

//go:generate moq -out service_mock_test.go . Repository

import (
	"fmt"
	"time"

	"github.com/siddhantac/fintra/domain"
)

const (
	dateLayout = "2006-01-02"
)

type Time string

func NewTime(t time.Time) Time {
	return Time(t.Round(time.Second).In(time.UTC).Format(dateLayout))
}

type Service struct {
	repo Repository
}

type Repository interface {
	Insert(*domain.Transaction) error
	GetByID(string) (*domain.Transaction, error)
	GetAll() ([]*domain.Transaction, error)
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTransaction(id string) (*domain.Transaction, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllTransactions() ([]*domain.Transaction, error) {
	return s.repo.GetAll()
}

func (s *Service) NewTransaction(amount float64, isDebit bool, date, category, transactionType, description, account string) (*domain.Transaction, error) {
	d, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	transaction := domain.NewTransaction(amount, d, isDebit, category, transactionType, description, account)

	if err := validateTransaction(*transaction); err != nil {
		return nil, err
	}

	if err := s.repo.Insert(transaction); err != nil {
		return nil, fmt.Errorf("repo.Insert: %w", err)
	}

	return transaction, nil
}

func validateTransaction(txn domain.Transaction) error {
	if err := validateType(txn.Type, txn.IsDebit); err != nil {
		return err
	}

	if txn.Currency == "" {
		return domain.ErrEmpty("currency")
	}

	if txn.Description == "" {
		return domain.ErrEmpty("description")
	}

	if txn.Date.IsZero() {
		return domain.ErrEmpty("date")
	}

	if txn.Category == "" {
		return domain.ErrEmpty("category")
	}

	if txn.Account == "" {
		return domain.ErrEmpty("account")
	}

	return nil
}

func validateType(txTyp domain.TransactionType, isDebit bool) error {
	if txTyp == "" {
		return domain.ErrEmpty("transaction type")
	}

	if _, ok := domain.TransactionTypes[txTyp]; !ok {
		return domain.ErrUnknownTransactionType(string(txTyp))
	}

	switch txTyp {
	case domain.TrTypeExpense, domain.TrTypeInvestment:
		if !isDebit {
			return domain.ErrMustBeDebit
		}
	case domain.TrTypeIncome:
		if isDebit {
			return domain.ErrMustBeCredit
		}
	}
	return nil
}
