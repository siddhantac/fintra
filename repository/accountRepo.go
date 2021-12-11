package repository

import (
	"github.com/siddhantac/fintra/domain"
)

type AccountStore interface {
	Store
	Update(id string, update interface{}) error
}

type AccountRepo struct {
	store AccountStore
}

func NewAccountRepository(store AccountStore) *AccountRepo {
	return &AccountRepo{store: store}
}

func (r *AccountRepo) GetByID(id string) (*domain.Account, error) {
	item, err := r.store.GetByID(id)
	if err != nil {
		return nil, err
	}

	acc := item.(*domain.Account)
	return acc, nil
}

/*
func (r *AccountRepo) CreditToAccount(id string, amount int64) (int64, error) {
	acc, err := r.GetByID(id)
	if err != nil {
		return 0, fmt.Errorf("failed to get account: %w", err)
	}

	acc.Balance += amount
	if err := r.store.Update(acc.ID, acc); err != nil {
		return 0, fmt.Errorf("failed to update: %w", err)
	}
	return acc.Balance, nil
}

func (r *AccountRepo) DebitFromAccount(id string, amount int64) (int64, error) {
	acc, err := r.GetByID(id)
	if err != nil {
		return 0, fmt.Errorf("failed to get account: %w", err)
	}

	acc.Balance -= amount
	if err := r.store.Update(acc.ID, acc); err != nil {
		return 0, fmt.Errorf("failed to update: %w", err)
	}
	return acc.Balance, nil
}
*/
