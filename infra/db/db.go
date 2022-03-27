package db

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var (
	bucketAccounts     = []byte("accounts")
	bucketTransactions = []byte("transactions")
	bucketCategories   = []byte("categories")
)

type BoltDB struct {
	*bolt.DB
}

func New(dbname string) (*BoltDB, error) {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	if err := createBucket(db, bucketAccounts); err != nil {
		return nil, fmt.Errorf("failed to create bucket '%s': %w", bucketAccounts, err)
	}
	if err := createBucket(db, bucketTransactions); err != nil {
		return nil, fmt.Errorf("failed to create bucket '%s': %w", bucketTransactions, err)
	}

	if err != nil {
		db.Close()
		return nil, err
	}

	return &BoltDB{DB: db}, nil
}

func (b *BoltDB) put(bucket, key, item []byte) error {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).Put(key, item)
	})
	return err
}

func (b *BoltDB) get(key, bucket []byte) []byte {
	var obj []byte
	b.DB.View(func(tx *bolt.Tx) error {
		obj = tx.Bucket(bucket).Get(key)
		return nil
	})
	return obj
}

func (b *BoltDB) getAll(bucket []byte) ([][]byte, error) {
	items := make([][]byte, 0)

	err := b.DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucket).Cursor()

		for k, item := c.First(); k != nil; k, item = c.Next() {
			items = append(items, item)
		}
		return nil
	})

	return items, err
}

func createBucket(db *bolt.DB, bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
			return err
		}
		return nil
	})
}
