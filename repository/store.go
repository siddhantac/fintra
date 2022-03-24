package repository

import "github.com/siddhantac/fintra/model"

type Store interface {
	Count() int
	Insert(id string, item interface{}) error
	GetByID(id string) (interface{}, error)
	GetAll() []interface{}
}

type Store2 interface {
	Store
	InsertTransaction(id string, txn *model.Transaction) error
	GetTransactionByID(id string, txn *model.Transaction) error
}
