package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/siddhantac/fintra/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTransaction(t *testing.T) {
	repo := &TransactionRepositoryMock{
		GetTransactionByIDFunc: func(id string) (*model.Transaction, error) {
			return &model.Transaction{}, nil
		},
	}
	s := NewTransactionService(repo, nil)

	_, err := s.GetTransaction("x")
	assert.NoError(t, err)
	assert.Len(t, repo.GetTransactionByIDCalls(), 1)
}

func TestNewTransaction(t *testing.T) {
	txnRepo := &TransactionRepositoryMock{
		InsertTransactionFunc: func(_ string, _ *model.Transaction) error {
			return nil
		},
	}
	accRepo := &AccountRepositoryMock{
		GetAccountByNameFunc: func(_ string) (*model.Account, error) {
			return &model.Account{Name: "some acc"}, nil
		},
		UpdateAccountFunc: func(_ string, _ *model.Account) (*model.Account, error) {
			return nil, nil
		},
	}

	accSvc := NewAccountService(accRepo)
	s := NewTransactionService(txnRepo, accSvc)

	txn, err := s.NewTransaction(
		12,
		true,
		"2021-10-11",
		string(model.TrCategoryEntertainment),
		string(model.TrTypeExpense),
		"desc",
		"Citibank",
	)
	assert.NoError(t, err)
	expectedTxn := &model.Transaction{
		IntAmount:   1200,
		Amount:      12,
		IsDebit:     true,
		Date:        time.Date(2021, time.October, 11, 0, 0, 0, 0, time.UTC),
		Category:    model.TrCategoryEntertainment,
		Type:        model.TrTypeExpense,
		Description: "desc",
		Account:     "Citibank",
		Currency:    "sgd",
	}
	require.Len(t, txnRepo.InsertTransactionCalls(), 1)
	require.Len(t, accRepo.GetAccountByNameCalls(), 2)
	require.Len(t, accRepo.UpdateAccountCalls(), 1)

	// don't compare date and ID as they are non-deterministic
	txn.Created = time.Time{}
	txn.Updated = time.Time{}
	txn.ID = ""
	require.Equal(t, expectedTxn, txn)
}

func TestTransactionValidation(t *testing.T) {

	tests := map[string]struct {
		expectedErr error
		txn         model.Transaction
	}{
		"valid transaction": {
			expectedErr: nil,
			txn: model.Transaction{
				Type:        model.TrTypeExpense,
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    model.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"empty transaction type": {
			expectedErr: fmt.Errorf("transaction type cannot be empty"),
			txn: model.Transaction{
				Type:        "",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    model.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"invalid transaction type": {
			expectedErr: fmt.Errorf("unknown transaction type: Loan"),
			txn: model.Transaction{
				Type:        "Loan",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    model.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"currency cannot be empty": {
			expectedErr: fmt.Errorf("currency cannot be empty"),
			txn: model.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "",
				Description: "some description",
				Date:        time.Now(),
				Category:    model.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"date cannot be empty": {
			expectedErr: fmt.Errorf("date cannot be empty"),
			txn: model.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Time{},
				Category:    model.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"category cannot be empty": {
			expectedErr: fmt.Errorf("category cannot be empty"),
			txn: model.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    "",
				Account:     "some account",
			},
		},
		"account cannot be empty": {
			expectedErr: fmt.Errorf("account cannot be empty"),
			txn: model.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    model.TrCategoryEntertainment,
				Account:     "",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateTransaction(test.txn)
			if test.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, test.expectedErr.Error())
			}
		})
	}
}
