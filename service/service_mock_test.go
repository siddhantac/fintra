// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"github.com/siddhantac/fintra/model"
	"sync"
)

// Ensure, that TransactionRepositoryMock does implement TransactionRepository.
// If this is not the case, regenerate this file with moq.
var _ TransactionRepository = &TransactionRepositoryMock{}

// TransactionRepositoryMock is a mock implementation of TransactionRepository.
//
// 	func TestSomethingThatUsesTransactionRepository(t *testing.T) {
//
// 		// make and configure a mocked TransactionRepository
// 		mockedTransactionRepository := &TransactionRepositoryMock{
// 			GetAllFunc: func() ([]*model.Transaction, error) {
// 				panic("mock out the GetAll method")
// 			},
// 			GetByIDFunc: func(s string) (*model.Transaction, error) {
// 				panic("mock out the GetByID method")
// 			},
// 			InsertFunc: func(transaction *model.Transaction) error {
// 				panic("mock out the Insert method")
// 			},
// 		}
//
// 		// use mockedTransactionRepository in code that requires TransactionRepository
// 		// and then make assertions.
//
// 	}
type TransactionRepositoryMock struct {
	// GetAllFunc mocks the GetAll method.
	GetAllFunc func() ([]*model.Transaction, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(s string) (*model.Transaction, error)

	// InsertFunc mocks the Insert method.
	InsertFunc func(transaction *model.Transaction) error

	// calls tracks calls to the methods.
	calls struct {
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// S is the s argument value.
			S string
		}
		// Insert holds details about calls to the Insert method.
		Insert []struct {
			// Transaction is the transaction argument value.
			Transaction *model.Transaction
		}
	}
	lockGetAll  sync.RWMutex
	lockGetByID sync.RWMutex
	lockInsert  sync.RWMutex
}

// GetAll calls GetAllFunc.
func (mock *TransactionRepositoryMock) GetAll() ([]*model.Transaction, error) {
	if mock.GetAllFunc == nil {
		panic("TransactionRepositoryMock.GetAllFunc: method is nil but TransactionRepository.GetAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc()
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//     len(mockedTransactionRepository.GetAllCalls())
func (mock *TransactionRepositoryMock) GetAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *TransactionRepositoryMock) GetByID(s string) (*model.Transaction, error) {
	if mock.GetByIDFunc == nil {
		panic("TransactionRepositoryMock.GetByIDFunc: method is nil but TransactionRepository.GetByID was just called")
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
//     len(mockedTransactionRepository.GetByIDCalls())
func (mock *TransactionRepositoryMock) GetByIDCalls() []struct {
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
func (mock *TransactionRepositoryMock) Insert(transaction *model.Transaction) error {
	if mock.InsertFunc == nil {
		panic("TransactionRepositoryMock.InsertFunc: method is nil but TransactionRepository.Insert was just called")
	}
	callInfo := struct {
		Transaction *model.Transaction
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
//     len(mockedTransactionRepository.InsertCalls())
func (mock *TransactionRepositoryMock) InsertCalls() []struct {
	Transaction *model.Transaction
} {
	var calls []struct {
		Transaction *model.Transaction
	}
	mock.lockInsert.RLock()
	calls = mock.calls.Insert
	mock.lockInsert.RUnlock()
	return calls
}

// Ensure, that AccountRepositoryMock does implement AccountRepository.
// If this is not the case, regenerate this file with moq.
var _ AccountRepository = &AccountRepositoryMock{}

// AccountRepositoryMock is a mock implementation of AccountRepository.
//
// 	func TestSomethingThatUsesAccountRepository(t *testing.T) {
//
// 		// make and configure a mocked AccountRepository
// 		mockedAccountRepository := &AccountRepositoryMock{
// 			GetByNameFunc: func(s string) (*model.Account, error) {
// 				panic("mock out the GetByName method")
// 			},
// 		}
//
// 		// use mockedAccountRepository in code that requires AccountRepository
// 		// and then make assertions.
//
// 	}
type AccountRepositoryMock struct {
	// GetByNameFunc mocks the GetByName method.
	GetByNameFunc func(s string) (*model.Account, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetByName holds details about calls to the GetByName method.
		GetByName []struct {
			// S is the s argument value.
			S string
		}
	}
	lockGetByName sync.RWMutex
}

// GetByName calls GetByNameFunc.
func (mock *AccountRepositoryMock) GetByName(s string) (*model.Account, error) {
	if mock.GetByNameFunc == nil {
		panic("AccountRepositoryMock.GetByNameFunc: method is nil but AccountRepository.GetByName was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockGetByName.Lock()
	mock.calls.GetByName = append(mock.calls.GetByName, callInfo)
	mock.lockGetByName.Unlock()
	return mock.GetByNameFunc(s)
}

// GetByNameCalls gets all the calls that were made to GetByName.
// Check the length with:
//     len(mockedAccountRepository.GetByNameCalls())
func (mock *AccountRepositoryMock) GetByNameCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockGetByName.RLock()
	calls = mock.calls.GetByName
	mock.lockGetByName.RUnlock()
	return calls
}
