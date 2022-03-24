package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	*bolt.DB
	testbucket *bolt.Bucket
}

func New() (*BoltDB, error) {
	db, err := bolt.Open("fintra.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("testbucket")); err != nil {
			return fmt.Errorf("failed to open bucket: %w", err)
		}
		return nil
	})

	if err != nil {
		db.Close()
		return nil, err
	}

	return &BoltDB{
		DB: db,
	}, nil
}

func (b *BoltDB) Count() int {
	return 0
}

func (b *BoltDB) Insert(id string, item interface{}) error {
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}

	b.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("testbucket")).Put([]byte(id), j)
	})
	return nil
}
func (b *BoltDB) GetByID2(id string, item interface{}) error {
	var obj []byte
	err := b.DB.View(func(tx *bolt.Tx) error {
		obj = tx.Bucket([]byte("testbucket")).Get([]byte(id))
		return nil
	})

	err = json.Unmarshal(obj, item)
	return err
}

func (b *BoltDB) GetByID(id string) (interface{}, error) {
	var item interface{}
	err := b.DB.View(func(tx *bolt.Tx) error {
		item = tx.Bucket([]byte("testbucket")).Get([]byte(id))
		return nil
	})
	return item, err
}

func (b *BoltDB) GetAll() []interface{} {
	return nil
}
