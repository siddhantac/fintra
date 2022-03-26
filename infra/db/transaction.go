package db

import (
	"encoding/json"
	"fmt"

	"github.com/siddhantac/fintra/model"
)

func (b *BoltDB) GetTransactionByID(id string) (*model.Transaction, error) {
	object := b.get([]byte(id), bucketTransactions)
	if object == nil {
		return nil, model.ErrNotFound
	}

	var txn model.Transaction
	err := json.Unmarshal(object, &txn)
	return &txn, err
}

func (b *BoltDB) InsertTransaction(id string, txn *model.Transaction) error {
	j, err := json.Marshal(txn)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = b.put(bucketTransactions, []byte(id), j)
	return err
}

func (b *BoltDB) GetAllTransactions() ([]*model.Transaction, error) {
	var txns []*model.Transaction

	items, err := b.getAll(bucketTransactions)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		var txn model.Transaction
		if err := json.Unmarshal(item, &txn); err != nil {
			return nil, err
		}
		txns = append(txns, &txn)
	}

	return txns, err
}
