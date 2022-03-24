package service

//go:generate moq -out transaction_mock_test.go . TransactionRepository AccountRepository

import (
	"fmt"
	"time"

	"github.com/siddhantac/fintra/infra/uid"
	"github.com/siddhantac/fintra/model"
)

const (
	dateLayout = "2006-01-02"
)

type RoundedTime string

func NewRoundedTime(t time.Time) RoundedTime {
	return RoundedTime(t.Round(time.Second).In(time.UTC).Format(dateLayout))
}

type TransactionService struct {
	txnRepo TransactionRepository
	accRepo AccountRepository
}

type TransactionRepository interface {
	InsertTransaction(id string, txn *model.Transaction) error
	GetTransactionByID(id string) (*model.Transaction, error)
	// GetAll() ([]*model.Transaction, error)
}

type AccountRepository interface {
	// GetAll() ([]*model.Account, error)
	InsertAccount(name string, txn *model.Account) error
	GetAccountByName(name string) (*model.Account, error)
}

func NewTransactionService(txnRepo TransactionRepository, accRepo AccountRepository) *TransactionService {
	return &TransactionService{
		txnRepo: txnRepo,
		accRepo: accRepo,
	}
}

func (s *TransactionService) GetTransaction(id string) (*model.Transaction, error) {
	return s.txnRepo.GetTransactionByID(id)
}

func (s *TransactionService) GetAllTransactions() ([]*model.Transaction, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TransactionService) NewTransaction(amount float64, isDebit bool, date, category, transactionType, description, account string) (*model.Transaction, error) {
	d, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	transaction := model.NewTransaction(uid.NewID(), amount, d, isDebit, category, transactionType, description, account)

	if err := validateTransaction(*transaction); err != nil {
		return nil, err
	}

	if _, err = s.accRepo.GetAccountByName(transaction.Account); err != nil {
		return nil, fmt.Errorf("error in account %s: %w", transaction.Account, err)
	}

	if err := s.txnRepo.InsertTransaction(transaction.ID, transaction); err != nil {
		return nil, fmt.Errorf("repo.Insert: %w", err)
	}
	// TODO update account balance

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
