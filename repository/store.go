package repository

type Store interface {
	Count() int
	Insert(id string, item interface{}) error
	GetByID(id string) (interface{}, error)
	GetAll() []interface{}
}

type Store2 interface {
	Store
	GetByID2(id string, item interface{}) error
}
