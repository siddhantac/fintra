package repository

import "github.com/siddhantac/fintra/domain"

type Store interface {
	Count() int
	Insert(id string, item interface{}) error
	GetByID(id string) (interface{}, error)
	GetAll() []interface{}
}

type TransactionRepository struct {
	store Store
}

func NewTransactionRepository(storageEngine Store) *TransactionRepository {
	return &TransactionRepository{
		store: storageEngine,
	}
}

func (r *TransactionRepository) Insert(txn *domain.Transaction) error {
	r.store.Insert(txn.ID, txn)
	return nil
}

func (r *TransactionRepository) GetByID(id string) (*domain.Transaction, error) {
	item, err := r.store.GetByID(id)
	if err != nil {
		return nil, err
	}

	txn := item.(*domain.Transaction)
	return txn, nil
}
