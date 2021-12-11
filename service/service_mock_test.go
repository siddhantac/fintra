// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"github.com/siddhantac/fintra/domain"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			GetByIDFunc: func(s string) (*domain.Transaction, error) {
// 				panic("mock out the GetByID method")
// 			},
// 			InsertFunc: func(transaction *domain.Transaction) error {
// 				panic("mock out the Insert method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(s string) (*domain.Transaction, error)

	// InsertFunc mocks the Insert method.
	InsertFunc func(transaction *domain.Transaction) error

	// calls tracks calls to the methods.
	calls struct {
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// S is the s argument value.
			S string
		}
		// Insert holds details about calls to the Insert method.
		Insert []struct {
			// Transaction is the transaction argument value.
			Transaction *domain.Transaction
		}
	}
	lockGetByID sync.RWMutex
	lockInsert  sync.RWMutex
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(s string) (*domain.Transaction, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByID was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(s)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//     len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// Insert calls InsertFunc.
func (mock *RepositoryMock) Insert(transaction *domain.Transaction) error {
	if mock.InsertFunc == nil {
		panic("RepositoryMock.InsertFunc: method is nil but Repository.Insert was just called")
	}
	callInfo := struct {
		Transaction *domain.Transaction
	}{
		Transaction: transaction,
	}
	mock.lockInsert.Lock()
	mock.calls.Insert = append(mock.calls.Insert, callInfo)
	mock.lockInsert.Unlock()
	return mock.InsertFunc(transaction)
}

// InsertCalls gets all the calls that were made to Insert.
// Check the length with:
//     len(mockedRepository.InsertCalls())
func (mock *RepositoryMock) InsertCalls() []struct {
	Transaction *domain.Transaction
} {
	var calls []struct {
		Transaction *domain.Transaction
	}
	mock.lockInsert.RLock()
	calls = mock.calls.Insert
	mock.lockInsert.RUnlock()
	return calls
}
