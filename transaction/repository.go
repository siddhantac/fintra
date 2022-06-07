package transaction

import (
	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/repository"
)

type Repository struct {
	store repository.Store
}

func NewTransactionRepository(storageEngine repository.Store) *Repository {
	return &Repository{
		store: storageEngine,
	}
}

func (r *Repository) Insert(txn *model.Transaction) error {
	r.store.Insert(txn.ID, txn)
	return nil
}

func (r *Repository) GetByID(id string) (*model.Transaction, error) {
	item, err := r.store.GetByID(id)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, model.ErrNotFound
	}

	txn := item.(*model.Transaction)
	return txn, nil
}

func (r *Repository) GetAll() ([]*model.Transaction, error) {
	items := r.store.GetAll()
	txns := make([]*model.Transaction, 0, len(items))
	for _, item := range items {
		txn := item.(*model.Transaction)
		txns = append(txns, txn)
	}
	return txns, nil
}
