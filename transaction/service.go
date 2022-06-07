package transaction

//go:generate moq -out transaction_mock_test.go . TransactionRepository AccountRepository

import (
	"fmt"
	"time"

	"github.com/siddhantac/fintra/account"
	"github.com/siddhantac/fintra/infra/uid"
	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/money"
)

const (
// dateLayout = "2006-01-02"
)

type RoundedTime string

func NewRoundedTime(t time.Time) RoundedTime {
	return RoundedTime(t.Round(time.Second).In(time.UTC).Format(dateLayout))
}

type Service struct {
	txnRepo TransactionRepository
	accSvc  *account.Service
}

type TransactionRepository interface {
	InsertTransaction(id string, txn *model.Transaction) error
	GetTransactionByID(id string) (*model.Transaction, error)
	GetAllTransactions() ([]*model.Transaction, error)
}

func NewService(txnRepo TransactionRepository, accService *account.Service) *Service {
	return &Service{
		txnRepo: txnRepo,
		accSvc:  accService,
	}
}

func (s *Service) GetTransaction(id string) (*model.Transaction, error) {
	return s.txnRepo.GetTransactionByID(id)
}

func (s *Service) GetAllTransactions() ([]*model.Transaction, error) {
	return s.txnRepo.GetAllTransactions()
}

func (s *Service) NewTransaction(amount float32, isDebit bool, date, category, transactionType, description, account string) (*model.Transaction, error) {
	d, err := time.Parse(dateLayout, date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	transaction := model.NewTransaction(uid.NewID(), money.NewMoney(amount), d, isDebit, category, transactionType, description, account)

	if err := validateTransaction(*transaction); err != nil {
		return nil, err
	}

	acc, err := s.accSvc.GetAccountByName(transaction.Account)
	if err != nil {
		return nil, fmt.Errorf("error in account %s: %w", transaction.Account, err)
	}

	if err := s.txnRepo.InsertTransaction(transaction.ID, transaction); err != nil {
		return nil, fmt.Errorf("repo.Insert: %w", err)
	}

	if _, err := s.accSvc.UpdateAccountBalance(acc.Name, transaction); err != nil {
		return nil, fmt.Errorf("accountService.updateAccountBalance: %w", err)
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

	if err := validateCategory(txn.Category); err != nil {
		return err
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
		return fmt.Errorf("transaction: %w", model.ErrUnknownType(string(txTyp)))
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

func validateCategory(category model.TransactionCategory) error {
	if category == "" {
		return model.ErrEmpty("category")
	}

	if _, ok := model.ValidCategories[category]; !ok {
		return fmt.Errorf("category: %w", model.ErrUnknownType(string(category)))
	}

	return nil
}
