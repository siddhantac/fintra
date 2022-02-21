package service

//go:generate moq -out service_mock_test.go . Repository

import (
	"fmt"
	"time"

	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/infra/uid"
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
	Insert(*model.Transaction) error
	GetByID(string) (*model.Transaction, error)
	GetAll() ([]*model.Transaction, error)
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTransaction(id string) (*model.Transaction, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllTransactions() ([]*model.Transaction, error) {
	return s.repo.GetAll()
}

func (s *Service) NewTransaction(amount float64, isDebit bool, date, category, transactionType, description, account string) (*model.Transaction, error) {
	d, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	transaction := model.NewTransaction(uid.NewID(), amount, d, isDebit, category, transactionType, description, account)

	if err := validateTransaction(*transaction); err != nil {
		return nil, err
	}

	if err := s.repo.Insert(transaction); err != nil {
		return nil, fmt.Errorf("repo.Insert: %w", err)
	}

	return transaction, nil
}

func validateTransaction(txn model.Transaction) error {
	if err := validateType(txn.Type, txn.IsDebit); err != nil {
		return err
	}

	if txn.Currency == "" {
		return model.ErrEmpty("currency")
	}

	if txn.Description == "" {
		return model.ErrEmpty("description")
	}

	if txn.Date.IsZero() {
		return model.ErrEmpty("date")
	}

	if txn.Category == "" {
		return model.ErrEmpty("category")
	}

	if txn.Account == "" {
		return model.ErrEmpty("account")
	}

	return nil
}

func validateType(txTyp model.TransactionType, isDebit bool) error {
	if txTyp == "" {
		return model.ErrEmpty("transaction type")
	}

	if _, ok := model.TransactionTypes[txTyp]; !ok {
		return model.ErrUnknownTransactionType(string(txTyp))
	}

	switch txTyp {
	case model.TrTypeExpense, model.TrTypeInvestment:
		if !isDebit {
			return model.ErrMustBeDebit
		}
	case model.TrTypeIncome:
		if isDebit {
			return model.ErrMustBeCredit
		}
	}
	return nil
}
