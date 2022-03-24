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

	var txn *model.Transaction
	err := json.Unmarshal(object, txn)
	return txn, err
}

func (b *BoltDB) InsertTransaction(id string, txn *model.Transaction) error {
	j, err := json.Marshal(txn)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = b.put(bucketTransactions, []byte(id), j)
	return err
}
