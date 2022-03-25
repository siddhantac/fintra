package service

import (
	"fmt"

	"github.com/siddhantac/fintra/model"
)

type AccountService struct {
	accRepo AccountRepository
}

type AccountRepository interface {
	GetAllAccounts() ([]*model.Account, error)
	InsertAccount(name string, txn *model.Account) error
	GetAccountByName(name string) (*model.Account, error)
	UpdateAccount(name string, account *model.Account) (*model.Account, error)
}

func NewAccountService(accountRepo AccountRepository) *AccountService {
	return &AccountService{accRepo: accountRepo}
}

func (s *AccountService) NewAccount(name string, startingBalance int) (*model.Account, error) {
	acc := model.NewAccount(name, startingBalance)
	if err := validateAccount(acc); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	if err := s.accRepo.InsertAccount(acc.Name, acc); err != nil {
		return nil, fmt.Errorf("repository insert failed: %w", err)
	}

	return acc, nil
}

func (s *AccountService) GetAllAccounts() ([]*model.Account, error) {
	return s.accRepo.GetAllAccounts()
}

func (s *AccountService) GetAccountByName(name string) (*model.Account, error) {
	return s.accRepo.GetAccountByName(name)
}

func (s *AccountService) UpdateAccountBalance(name string, txn *model.Transaction) (*model.Account, error) {
	acc, err := s.accRepo.GetAccountByName(name)
	if err != nil {
		return nil, err
	}

	if txn.IsDebit {
		acc.Debit(txn.IntAmount)
	} else {
		acc.Credit(txn.IntAmount)
	}

	updatedAccount, err := s.accRepo.UpdateAccount(acc.Name, acc)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

func validateAccount(account *model.Account) error {
	if account.Name == "" {
		return model.ErrEmpty("name")
	}
	return nil
}
