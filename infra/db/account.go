package db

import (
	"encoding/json"
	"fmt"

	"github.com/siddhantac/fintra/model"
)

func (b *BoltDB) GetAccountByName(name string) (*model.Account, error) {
	object := b.get([]byte(name), bucketAccounts)
	if object == nil {
		return nil, model.ErrNotFound
	}

	var account model.Account
	err := json.Unmarshal(object, &account)
	return &account, err
}

func (b *BoltDB) InsertAccount(name string, account *model.Account) error {
	j, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = b.put(bucketAccounts, []byte(name), j)
	return err
}

func (b *BoltDB) GetAllAccounts() ([]*model.Account, error) {
	var accs []*model.Account

	items, err := b.getAll(bucketAccounts)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		var acc model.Account
		if err := json.Unmarshal(item, &acc); err != nil {
			return nil, err
		}
		accs = append(accs, &acc)
	}

	return accs, err
}

func (b *BoltDB) UpdateAccount(name string, update *model.Account) (*model.Account, error) {
	return nil, fmt.Errorf("not implemented")
}
