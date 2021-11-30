package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ms := NewMemStore()
	assert.Equal(t, 0, ms.Count())

	// tx, err := domain.NewTransaction(23, time.Now(), true, string(domain.TrCategoryEntertainment), string(domain.TrTypeExpense), "desc", "Citibank")
	// assert.NoError(t, err)
	item := 23
	ms.Insert("id23", item)

	assert.Equal(t, 1, ms.Count())

	gotItem, err := ms.GetByID("id23")
	assert.NoError(t, err)
	assert.Equal(t, item, gotItem)

	// tx2, err := domain.NewTransaction(11, time.Now(), true, string(domain.TrCategoryMeals), string(domain.TrTypeExpense), "desc", "Citibank")
	// assert.NoError(t, err)
	city := map[string]interface{}{
		"id":   23,
		"name": "Singapore",
	}
	ms.Insert("id1", city)
	assert.Equal(t, 2, ms.Count())

	txns := ms.GetAll()
	assert.Len(t, txns, 2)
	assert.Equal(t, item, txns[0])
	assert.Equal(t, city, txns[1])
}
