package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ms := NewMemStore()
	assert.Equal(t, 0, ms.Len())

	tx, err := NewTransaction(23, time.Now(), true, string(TrCategoryEntertainment), string(TrTypeExpense), "", "Citibank")
	assert.NoError(t, err)
	ms.Insert(tx)

	assert.Equal(t, 1, ms.Len())

	tx2, err := NewTransaction(11, time.Now(), true, string(TrCategoryMeals), string(TrTypeExpense), "", "Citibank")
	assert.NoError(t, err)
	ms.Insert(tx2)

}
