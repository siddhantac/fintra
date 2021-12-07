package repository

import "github.com/siddhantac/fintra/domain"

type AccountRepo struct {
	store Store
}

func NewAccountRepository(store Store) *AccountRepo {
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
