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
// 			GetAllTransactionsFunc: func() ([]*model.Transaction, error) {
// 				panic("mock out the GetAllTransactions method")
// 			},
// 			GetTransactionByIDFunc: func(id string) (*model.Transaction, error) {
// 				panic("mock out the GetTransactionByID method")
// 			},
// 			InsertTransactionFunc: func(id string, txn *model.Transaction) error {
// 				panic("mock out the InsertTransaction method")
// 			},
// 		}
//
// 		// use mockedTransactionRepository in code that requires TransactionRepository
// 		// and then make assertions.
//
// 	}
type TransactionRepositoryMock struct {
	// GetAllTransactionsFunc mocks the GetAllTransactions method.
	GetAllTransactionsFunc func() ([]*model.Transaction, error)

	// GetTransactionByIDFunc mocks the GetTransactionByID method.
	GetTransactionByIDFunc func(id string) (*model.Transaction, error)

	// InsertTransactionFunc mocks the InsertTransaction method.
	InsertTransactionFunc func(id string, txn *model.Transaction) error

	// calls tracks calls to the methods.
	calls struct {
		// GetAllTransactions holds details about calls to the GetAllTransactions method.
		GetAllTransactions []struct {
		}
		// GetTransactionByID holds details about calls to the GetTransactionByID method.
		GetTransactionByID []struct {
			// ID is the id argument value.
			ID string
		}
		// InsertTransaction holds details about calls to the InsertTransaction method.
		InsertTransaction []struct {
			// ID is the id argument value.
			ID string
			// Txn is the txn argument value.
			Txn *model.Transaction
		}
	}
	lockGetAllTransactions sync.RWMutex
	lockGetTransactionByID sync.RWMutex
	lockInsertTransaction  sync.RWMutex
}

// GetAllTransactions calls GetAllTransactionsFunc.
func (mock *TransactionRepositoryMock) GetAllTransactions() ([]*model.Transaction, error) {
	if mock.GetAllTransactionsFunc == nil {
		panic("TransactionRepositoryMock.GetAllTransactionsFunc: method is nil but TransactionRepository.GetAllTransactions was just called")
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
//     len(mockedTransactionRepository.GetAllTransactionsCalls())
func (mock *TransactionRepositoryMock) GetAllTransactionsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllTransactions.RLock()
	calls = mock.calls.GetAllTransactions
	mock.lockGetAllTransactions.RUnlock()
	return calls
}

// GetTransactionByID calls GetTransactionByIDFunc.
func (mock *TransactionRepositoryMock) GetTransactionByID(id string) (*model.Transaction, error) {
	if mock.GetTransactionByIDFunc == nil {
		panic("TransactionRepositoryMock.GetTransactionByIDFunc: method is nil but TransactionRepository.GetTransactionByID was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetTransactionByID.Lock()
	mock.calls.GetTransactionByID = append(mock.calls.GetTransactionByID, callInfo)
	mock.lockGetTransactionByID.Unlock()
	return mock.GetTransactionByIDFunc(id)
}

// GetTransactionByIDCalls gets all the calls that were made to GetTransactionByID.
// Check the length with:
//     len(mockedTransactionRepository.GetTransactionByIDCalls())
func (mock *TransactionRepositoryMock) GetTransactionByIDCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetTransactionByID.RLock()
	calls = mock.calls.GetTransactionByID
	mock.lockGetTransactionByID.RUnlock()
	return calls
}

// InsertTransaction calls InsertTransactionFunc.
func (mock *TransactionRepositoryMock) InsertTransaction(id string, txn *model.Transaction) error {
	if mock.InsertTransactionFunc == nil {
		panic("TransactionRepositoryMock.InsertTransactionFunc: method is nil but TransactionRepository.InsertTransaction was just called")
	}
	callInfo := struct {
		ID  string
		Txn *model.Transaction
	}{
		ID:  id,
		Txn: txn,
	}
	mock.lockInsertTransaction.Lock()
	mock.calls.InsertTransaction = append(mock.calls.InsertTransaction, callInfo)
	mock.lockInsertTransaction.Unlock()
	return mock.InsertTransactionFunc(id, txn)
}

// InsertTransactionCalls gets all the calls that were made to InsertTransaction.
// Check the length with:
//     len(mockedTransactionRepository.InsertTransactionCalls())
func (mock *TransactionRepositoryMock) InsertTransactionCalls() []struct {
	ID  string
	Txn *model.Transaction
} {
	var calls []struct {
		ID  string
		Txn *model.Transaction
	}
	mock.lockInsertTransaction.RLock()
	calls = mock.calls.InsertTransaction
	mock.lockInsertTransaction.RUnlock()
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
// 			GetAccountByNameFunc: func(name string) (*model.Account, error) {
// 				panic("mock out the GetAccountByName method")
// 			},
// 			GetAllAccountsFunc: func() ([]*model.Account, error) {
// 				panic("mock out the GetAllAccounts method")
// 			},
// 			InsertAccountFunc: func(name string, txn *model.Account) error {
// 				panic("mock out the InsertAccount method")
// 			},
// 		}
//
// 		// use mockedAccountRepository in code that requires AccountRepository
// 		// and then make assertions.
//
// 	}
type AccountRepositoryMock struct {
	// GetAccountByNameFunc mocks the GetAccountByName method.
	GetAccountByNameFunc func(name string) (*model.Account, error)

	// GetAllAccountsFunc mocks the GetAllAccounts method.
	GetAllAccountsFunc func() ([]*model.Account, error)

	// InsertAccountFunc mocks the InsertAccount method.
	InsertAccountFunc func(name string, txn *model.Account) error

	// calls tracks calls to the methods.
	calls struct {
		// GetAccountByName holds details about calls to the GetAccountByName method.
		GetAccountByName []struct {
			// Name is the name argument value.
			Name string
		}
		// GetAllAccounts holds details about calls to the GetAllAccounts method.
		GetAllAccounts []struct {
		}
		// InsertAccount holds details about calls to the InsertAccount method.
		InsertAccount []struct {
			// Name is the name argument value.
			Name string
			// Txn is the txn argument value.
			Txn *model.Account
		}
	}
	lockGetAccountByName sync.RWMutex
	lockGetAllAccounts   sync.RWMutex
	lockInsertAccount    sync.RWMutex
}

// GetAccountByName calls GetAccountByNameFunc.
func (mock *AccountRepositoryMock) GetAccountByName(name string) (*model.Account, error) {
	if mock.GetAccountByNameFunc == nil {
		panic("AccountRepositoryMock.GetAccountByNameFunc: method is nil but AccountRepository.GetAccountByName was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockGetAccountByName.Lock()
	mock.calls.GetAccountByName = append(mock.calls.GetAccountByName, callInfo)
	mock.lockGetAccountByName.Unlock()
	return mock.GetAccountByNameFunc(name)
}

// GetAccountByNameCalls gets all the calls that were made to GetAccountByName.
// Check the length with:
//     len(mockedAccountRepository.GetAccountByNameCalls())
func (mock *AccountRepositoryMock) GetAccountByNameCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockGetAccountByName.RLock()
	calls = mock.calls.GetAccountByName
	mock.lockGetAccountByName.RUnlock()
	return calls
}

// GetAllAccounts calls GetAllAccountsFunc.
func (mock *AccountRepositoryMock) GetAllAccounts() ([]*model.Account, error) {
	if mock.GetAllAccountsFunc == nil {
		panic("AccountRepositoryMock.GetAllAccountsFunc: method is nil but AccountRepository.GetAllAccounts was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAllAccounts.Lock()
	mock.calls.GetAllAccounts = append(mock.calls.GetAllAccounts, callInfo)
	mock.lockGetAllAccounts.Unlock()
	return mock.GetAllAccountsFunc()
}

// GetAllAccountsCalls gets all the calls that were made to GetAllAccounts.
// Check the length with:
//     len(mockedAccountRepository.GetAllAccountsCalls())
func (mock *AccountRepositoryMock) GetAllAccountsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllAccounts.RLock()
	calls = mock.calls.GetAllAccounts
	mock.lockGetAllAccounts.RUnlock()
	return calls
}

// InsertAccount calls InsertAccountFunc.
func (mock *AccountRepositoryMock) InsertAccount(name string, txn *model.Account) error {
	if mock.InsertAccountFunc == nil {
		panic("AccountRepositoryMock.InsertAccountFunc: method is nil but AccountRepository.InsertAccount was just called")
	}
	callInfo := struct {
		Name string
		Txn  *model.Account
	}{
		Name: name,
		Txn:  txn,
	}
	mock.lockInsertAccount.Lock()
	mock.calls.InsertAccount = append(mock.calls.InsertAccount, callInfo)
	mock.lockInsertAccount.Unlock()
	return mock.InsertAccountFunc(name, txn)
}

// InsertAccountCalls gets all the calls that were made to InsertAccount.
// Check the length with:
//     len(mockedAccountRepository.InsertAccountCalls())
func (mock *AccountRepositoryMock) InsertAccountCalls() []struct {
	Name string
	Txn  *model.Account
} {
	var calls []struct {
		Name string
		Txn  *model.Account
	}
	mock.lockInsertAccount.RLock()
	calls = mock.calls.InsertAccount
	mock.lockInsertAccount.RUnlock()
	return calls
}
