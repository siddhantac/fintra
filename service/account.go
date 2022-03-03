package service

import (
	"fmt"

	"github.com/siddhantac/fintra/model"
)

type AccountService struct {
	accRepo AccountRepository
}

func NewAccountService(accountRepo AccountRepository) *AccountService {
	return &AccountService{accRepo: accountRepo}
}

func (s *AccountService) NewAccount(name string, startingBalance int) (*model.Account, error) {
	acc := model.NewAccount(name, startingBalance)
	if err := validateAccount(acc); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	if err := s.accRepo.Insert(acc); err != nil {
		return nil, fmt.Errorf("repository insert failed: %w", err)
	}

	return acc, nil
}

func (s *AccountService) GetAllAccounts() ([]*model.Account, error) {
	return s.accRepo.GetAll()
}

func (s *AccountService) GetAccountByName(name string) (*model.Account, error) {
	return s.accRepo.GetByName(name)
}

func validateAccount(account *model.Account) error {
	if account.Name() == "" {
		return model.ErrEmpty("name")
	}
	return nil
}
