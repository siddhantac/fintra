package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ms := NewMemStore()
	assert.Equal(t, 0, ms.Len())

	tx, err := NewTransaction(23, time.Now(), true, string(TrCategoryEntertainment), string(TrTypeExpense), "desc", "Citibank")
	assert.NoError(t, err)
	ms.Insert(tx)

	assert.Equal(t, 1, ms.Len())

	tx2, err := NewTransaction(11, time.Now(), true, string(TrCategoryMeals), string(TrTypeExpense), "desc", "Citibank")
	assert.NoError(t, err)
	ms.Insert(tx2)

	txns := ms.GetAll()
	assert.Len(t, txns, 2)
	assert.Equal(t, tx, txns[0])
	assert.Equal(t, tx2, txns[1])
}
