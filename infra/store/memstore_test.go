package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ms := NewMemStore()
	assert.Equal(t, 0, ms.Count())

	// tx, err := model.NewTransaction(23, time.Now(), true, string(model.TrCategoryEntertainment), string(model.TrTypeExpense), "desc", "Citibank")
	// assert.NoError(t, err)
	item := 23
	ms.Insert("id23", item)

	assert.Equal(t, 1, ms.Count())

	gotItem, err := ms.GetByID("id23")
	assert.NoError(t, err)
	assert.Equal(t, item, gotItem)

	// tx2, err := model.NewTransaction(11, time.Now(), true, string(model.TrCategoryMeals), string(model.TrTypeExpense), "desc", "Citibank")
	// assert.NoError(t, err)
	city := map[string]interface{}{
		"id":   23,
		"name": "Singapore",
	}
	ms.Insert("id1", city)
	assert.Equal(t, 2, ms.Count())

	txns := ms.GetAll()
	assert.Len(t, txns, 2)
	assert.Contains(t, txns, item)
	assert.Contains(t, txns, city)
}
