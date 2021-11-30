package store

import (
	"sync"
)

type MemStore struct {
	inited bool
	Items  map[string]interface{}
	mtx    sync.Mutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		inited: true,
		Items:  make(map[string]interface{}),
		mtx:    sync.Mutex{},
	}
}

func (ms *MemStore) Count() int {
	if !ms.inited {
		return 0
	}
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	return len(ms.Items)
}

func (ms *MemStore) Insert(id string, item interface{}) error {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	ms.Items[id] = item
	return nil
}

func (ms *MemStore) GetByID(id string) (interface{}, error) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	return ms.Items[id], nil
}

func (ms *MemStore) GetAll() []interface{} {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	all := make([]interface{}, 0, len(ms.Items))
	for _, tx := range ms.Items {
		all = append(all, tx)
	}
	return all
}
