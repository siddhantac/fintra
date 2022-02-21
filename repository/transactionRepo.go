package repository

import (
	"github.com/siddhantac/fintra/model"
)

type TransactionRepository struct {
	store Store
}

func NewTransactionRepository(storageEngine Store) *TransactionRepository {
	return &TransactionRepository{
		store: storageEngine,
	}
}

func (r *TransactionRepository) Insert(txn *model.Transaction) error {
	r.store.Insert(txn.ID, txn)
	return nil
}

func (r *TransactionRepository) GetByID(id string) (*model.Transaction, error) {
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

func (r *TransactionRepository) GetAll() ([]*model.Transaction, error) {
	items := r.store.GetAll()
	txns := make([]*model.Transaction, 0, len(items))
	for _, item := range items {
		txn := item.(*model.Transaction)
		txns = append(txns, txn)
	}
	return txns, nil
}
