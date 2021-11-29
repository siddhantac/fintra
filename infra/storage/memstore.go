package repository

import (
	"sync"

	"github.com/siddhantac/fintra/domain"
)

type MemStore struct {
	inited       bool
	Transactions map[string]*domain.Transaction
	mtx          sync.Mutex
}

func NewMemStore() *MemStore {
	return &MemStore{
		inited:       true,
		Transactions: make(map[string]*domain.Transaction),
		mtx:          sync.Mutex{},
	}
}

func (ms *MemStore) Len() int {
	if !ms.inited {
		return 0
	}
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	return len(ms.Transactions)
}

func (ms *MemStore) Insert(t *domain.Transaction) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	ms.Transactions[t.ID] = t
}

func (ms *MemStore) Get(id string) *domain.Transaction {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	return ms.Transactions[id]
}

func (ms *MemStore) GetAll() []*domain.Transaction {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	allTx := make([]*domain.Transaction, 0, len(ms.Transactions))
	for _, tx := range ms.Transactions {
		allTx = append(allTx, tx)
	}
	return allTx
}
