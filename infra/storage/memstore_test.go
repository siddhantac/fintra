package repository

import (
	"testing"
	"time"

	"github.com/siddhantac/fintra/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ms := NewMemStore()
	assert.Equal(t, 0, ms.NumTransactions())

	tx, err := domain.NewTransaction(23, time.Now(), true, string(domain.TrCategoryEntertainment), string(domain.TrTypeExpense), "desc", "Citibank")
	assert.NoError(t, err)
	err = ms.Insert(tx)
	assert.NoError(t, err)

	assert.Equal(t, 1, ms.NumTransactions())

	tx2, err := domain.NewTransaction(11, time.Now(), true, string(domain.TrCategoryMeals), string(domain.TrTypeExpense), "desc", "Citibank")
	assert.NoError(t, err)
	err = ms.Insert(tx2)
	assert.NoError(t, err)

	txns := ms.GetAll()
	assert.Len(t, txns, 2)
	assert.Equal(t, tx, txns[0])
	assert.Equal(t, tx2, txns[1])
}
