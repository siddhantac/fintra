package repository

import (
	"github.com/siddhantac/fintra/domain"
)

type AccountRepo struct {
	store Store
}

func NewAccountRepository(storage Store) *AccountRepo {
	return &AccountRepo{store: storage}
}

func (r *AccountRepo) Insert(account *domain.Account) error {
	r.store.Insert(account.Name(), account)
	return nil
}

func (r *AccountRepo) GetByID(name string) (*domain.Account, error) {
	item, err := r.store.GetByID(name)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, domain.ErrNotFound
	}

	acc := item.(*domain.Account)
	return acc, nil
}

func (r *AccountRepo) GetAll() ([]*domain.Account, error) {
	items := r.store.GetAll()
	accs := make([]*domain.Account, 0, len(items))
	for _, item := range items {
		acc := item.(*domain.Account)
		accs = append(accs, acc)
	}
	return accs, nil
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
