package repository

import (
	"github.com/siddhantac/fintra/model"
)

type TransactionRepository struct {
	store Store2
}

func NewTransactionRepository(storageEngine Store2) *TransactionRepository {
	return &TransactionRepository{
		store: storageEngine,
	}
}

func (r *TransactionRepository) Insert(txn *model.Transaction) error {
	r.store.InsertTransaction(txn.ID, txn)
	return nil
}

func (r *TransactionRepository) GetByID(id string) (*model.Transaction, error) {
	var txn model.Transaction
	err := r.store.GetTransactionByID(id, &txn)
	if err != nil {
		return nil, err
	}

	// if item == nil {
	//     return nil, model.ErrNotFound
	// }

	// txn := item.(*model.Transaction)
	return &txn, nil
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
