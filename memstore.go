package main

import "sync"

type MemStore struct {
	inited       bool
	transactions map[string]*Transaction
	mtx          sync.Mutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		inited:       true,
		transactions: make(map[string]*Transaction),
		mtx:          sync.Mutex{},
	}
}

func (ms *MemStore) Len() int {
	if !ms.inited {
		return 0
	}
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	return len(ms.transactions)
}

func (ms *MemStore) Insert(t *Transaction) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	ms.transactions[t.ID] = t
}

func (ms *MemStore) Get(id string) *Transaction {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	return ms.transactions[id]
}
