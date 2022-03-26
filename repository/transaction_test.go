package repository

import (
	"testing"

	"github.com/siddhantac/fintra/infra/store"
	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/money"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	storage := store.NewMemStore()
	item := &model.Transaction{
		ID:      "23",
		Amount:  money.NewMoney(100),
		Account: "Citibank",
	}
	storage.Items = map[string]interface{}{
		"23": item,
	}

	repo := NewTransactionRepository(storage)
	gotItem, err := repo.GetByID("23")
	assert.NoError(t, err)
	assert.Equal(t, item, gotItem)

}
