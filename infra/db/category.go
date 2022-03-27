package db

import (
	"encoding/json"
	"fmt"

	"github.com/siddhantac/fintra/model"
)

func (b *BoltDB) GetCategoryByName(id string) (*model.Category, error) {
	object := b.get([]byte(id), bucketCategories)
	if object == nil {
		return nil, model.ErrNotFound
	}

	var ctg model.Category
	err := json.Unmarshal(object, &ctg)
	return &ctg, err
}

func (b *BoltDB) InsertCategory(name string, category *model.Category) error {
	j, err := json.Marshal(category)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = b.put(bucketCategories, []byte(name), j)
	return err
}

func (b *BoltDB) UpdateCategory(name string, category *model.Category) error {
	_, err := b.GetCategoryByName(name)
	if err != nil {
		return err
	}

	return b.InsertCategory(name, category)
}
