package repository

type Store interface {
	Count() int
	Insert(id string, item interface{}) error
	GetByID(id string) (interface{}, error)
	GetAll() []interface{}
}
