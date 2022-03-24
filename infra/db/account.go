package db

import (
	"encoding/json"
	"fmt"

	"github.com/siddhantac/fintra/model"
)

func (b *BoltDB) GetAccountByName(name string) (*model.Account, error) {
	object := b.get([]byte(name), bucketTransactions)
	if object == nil {
		return nil, model.ErrNotFound
	}

	var account *model.Account
	err := json.Unmarshal(object, account)
	return account, err
}

func (b *BoltDB) InsertAccount(name string, account *model.Account) error {
	j, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = b.put(bucketTransactions, []byte(name), j)
	return err
}
