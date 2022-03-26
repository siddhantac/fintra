// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package rest

import (
	"github.com/siddhantac/fintra/model"
	"sync"
)

// Ensure, that TransactionServiceMock does implement TransactionService.
// If this is not the case, regenerate this file with moq.
var _ TransactionService = &TransactionServiceMock{}

// TransactionServiceMock is a mock implementation of TransactionService.
//
// 	func TestSomethingThatUsesTransactionService(t *testing.T) {
//
// 		// make and configure a mocked TransactionService
// 		mockedTransactionService := &TransactionServiceMock{
// 			GetAllTransactionsFunc: func() ([]*model.Transaction, error) {
// 				panic("mock out the GetAllTransactions method")
// 			},
// 			GetTransactionFunc: func(id string) (*model.Transaction, error) {
// 				panic("mock out the GetTransaction method")
// 			},
// 			NewTransactionFunc: func(amount float32, isDebit bool, date string, category string, transactionType string, description string, account string) (*model.Transaction, error) {
// 				panic("mock out the NewTransaction method")
// 			},
// 		}
//
// 		// use mockedTransactionService in code that requires TransactionService
// 		// and then make assertions.
//
// 	}
type TransactionServiceMock struct {
	// GetAllTransactionsFunc mocks the GetAllTransactions method.
	GetAllTransactionsFunc func() ([]*model.Transaction, error)

	// GetTransactionFunc mocks the GetTransaction method.
	GetTransactionFunc func(id string) (*model.Transaction, error)

	// NewTransactionFunc mocks the NewTransaction method.
	NewTransactionFunc func(amount float32, isDebit bool, date string, category string, transactionType string, description string, account string) (*model.Transaction, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetAllTransactions holds details about calls to the GetAllTransactions method.
		GetAllTransactions []struct {
		}
		// GetTransaction holds details about calls to the GetTransaction method.
		GetTransaction []struct {
			// ID is the id argument value.
			ID string
		}
		// NewTransaction holds details about calls to the NewTransaction method.
		NewTransaction []struct {
			// Amount is the amount argument value.
			Amount float32
			// IsDebit is the isDebit argument value.
			IsDebit bool
			// Date is the date argument value.
			Date string
			// Category is the category argument value.
			Category string
			// TransactionType is the transactionType argument value.
			TransactionType string
			// Description is the description argument value.
			Description string
			// Account is the account argument value.
			Account string
		}
	}
	lockGetAllTransactions sync.RWMutex
	lockGetTransaction     sync.RWMutex
	lockNewTransaction     sync.RWMutex
}

// GetAllTransactions calls GetAllTransactionsFunc.
func (mock *TransactionServiceMock) GetAllTransactions() ([]*model.Transaction, error) {
	if mock.GetAllTransactionsFunc == nil {
		panic("TransactionServiceMock.GetAllTransactionsFunc: method is nil but TransactionService.GetAllTransactions was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAllTransactions.Lock()
	mock.calls.GetAllTransactions = append(mock.calls.GetAllTransactions, callInfo)
	mock.lockGetAllTransactions.Unlock()
	return mock.GetAllTransactionsFunc()
}

// GetAllTransactionsCalls gets all the calls that were made to GetAllTransactions.
// Check the length with:
//     len(mockedTransactionService.GetAllTransactionsCalls())
func (mock *TransactionServiceMock) GetAllTransactionsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllTransactions.RLock()
	calls = mock.calls.GetAllTransactions
	mock.lockGetAllTransactions.RUnlock()
	return calls
}

// GetTransaction calls GetTransactionFunc.
func (mock *TransactionServiceMock) GetTransaction(id string) (*model.Transaction, error) {
	if mock.GetTransactionFunc == nil {
		panic("TransactionServiceMock.GetTransactionFunc: method is nil but TransactionService.GetTransaction was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetTransaction.Lock()
	mock.calls.GetTransaction = append(mock.calls.GetTransaction, callInfo)
	mock.lockGetTransaction.Unlock()
	return mock.GetTransactionFunc(id)
}

// GetTransactionCalls gets all the calls that were made to GetTransaction.
// Check the length with:
//     len(mockedTransactionService.GetTransactionCalls())
func (mock *TransactionServiceMock) GetTransactionCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetTransaction.RLock()
	calls = mock.calls.GetTransaction
	mock.lockGetTransaction.RUnlock()
	return calls
}

// NewTransaction calls NewTransactionFunc.
func (mock *TransactionServiceMock) NewTransaction(amount float32, isDebit bool, date string, category string, transactionType string, description string, account string) (*model.Transaction, error) {
	if mock.NewTransactionFunc == nil {
		panic("TransactionServiceMock.NewTransactionFunc: method is nil but TransactionService.NewTransaction was just called")
	}
	callInfo := struct {
		Amount          float32
		IsDebit         bool
		Date            string
		Category        string
		TransactionType string
		Description     string
		Account         string
	}{
		Amount:          amount,
		IsDebit:         isDebit,
		Date:            date,
		Category:        category,
		TransactionType: transactionType,
		Description:     description,
		Account:         account,
	}
	mock.lockNewTransaction.Lock()
	mock.calls.NewTransaction = append(mock.calls.NewTransaction, callInfo)
	mock.lockNewTransaction.Unlock()
	return mock.NewTransactionFunc(amount, isDebit, date, category, transactionType, description, account)
}

// NewTransactionCalls gets all the calls that were made to NewTransaction.
// Check the length with:
//     len(mockedTransactionService.NewTransactionCalls())
func (mock *TransactionServiceMock) NewTransactionCalls() []struct {
	Amount          float32
	IsDebit         bool
	Date            string
	Category        string
	TransactionType string
	Description     string
	Account         string
} {
	var calls []struct {
		Amount          float32
		IsDebit         bool
		Date            string
		Category        string
		TransactionType string
		Description     string
		Account         string
	}
	mock.lockNewTransaction.RLock()
	calls = mock.calls.NewTransaction
	mock.lockNewTransaction.RUnlock()
	return calls
}
