package account

import (
	"fmt"

	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/money"
)

type Service struct {
	accRepo AccountRepository
}

type AccountRepository interface {
	GetAllAccounts() ([]*model.Account, error)
	InsertAccount(name string, txn *model.Account) error
	GetAccountByName(name string) (*model.Account, error)
	UpdateAccount(name string, account *model.Account) (*model.Account, error)
}

func NewService(accountRepo AccountRepository) *Service {
	return &Service{accRepo: accountRepo}
}

func (s *Service) NewAccount(name string, startingBalance float32) (*model.Account, error) {
	acc := model.NewAccount(name, money.NewMoney(startingBalance))
	if err := validateAccount(acc); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	if err := s.accRepo.InsertAccount(acc.Name, acc); err != nil {
		return nil, fmt.Errorf("repository insert failed: %w", err)
	}

	return acc, nil
}

func (s *Service) GetAllAccounts() ([]*model.Account, error) {
	return s.accRepo.GetAllAccounts()
}

func (s *Service) GetAccountByName(name string) (*model.Account, error) {
	return s.accRepo.GetAccountByName(name)
}

func (s *Service) UpdateAccountBalance(name string, txn *model.Transaction) (*model.Account, error) {
	acc, err := s.accRepo.GetAccountByName(name)
	if err != nil {
		return nil, err
	}

	if txn.IsDebit {
		acc.Debit(txn.Amount)
	} else {
		acc.Credit(txn.Amount)
	}

	updatedAccount, err := s.accRepo.UpdateAccount(acc.Name, acc)
	if err != nil {
		return nil, fmt.Errorf("accountRepository.UpdateAccount: %w", err)
	}

	return updatedAccount, nil
}

func validateAccount(account *model.Account) error {
	if account.Name == "" {
		return model.ErrEmpty("name")
	}
	return nil
}
