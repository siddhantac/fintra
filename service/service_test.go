package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/siddhantac/fintra/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTransaction(t *testing.T) {
	repo := &RepositoryMock{
		GetByIDFunc: func(id string) (*domain.Transaction, error) {
			return &domain.Transaction{}, nil
		},
	}
	s := NewService(repo)

	_, err := s.GetTransaction("x")
	assert.NoError(t, err)
	assert.Len(t, repo.GetByIDCalls(), 1)
}

func TestNewTransaction(t *testing.T) {
	repo := &RepositoryMock{
		InsertFunc: func(_ *domain.Transaction) error {
			return nil
		},
	}
	s := NewService(repo)

	txn, err := s.NewTransaction(
		12,
		true,
		"2021-10-11",
		string(domain.TrCategoryEntertainment),
		string(domain.TrTypeExpense),
		"desc",
		"Citibank",
	)
	assert.NoError(t, err)
	assert.Len(t, repo.InsertCalls(), 1)
	assert.Equal(t, 12.00, txn.Amount)
	assert.Equal(t, true, txn.IsDebit)
	assert.Equal(t, time.Date(2021, time.October, 11, 0, 0, 0, 0, time.UTC), txn.Date)
	assert.Equal(t, domain.TrCategoryEntertainment, txn.Category)
	assert.Equal(t, domain.TrTypeExpense, txn.Type)
	assert.Equal(t, "desc", txn.Description)
	assert.Equal(t, "Citibank", txn.Account)
}

func TestTransactionValidation(t *testing.T) {

	tests := map[string]struct {
		expectedErr error
		txn         domain.Transaction
	}{
		"valid transaction": {
			expectedErr: nil,
			txn: domain.Transaction{
				Type:        domain.TrTypeExpense,
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    domain.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"empty transaction type": {
			expectedErr: fmt.Errorf("transaction type cannot be empty"),
			txn: domain.Transaction{
				Type:        "",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    domain.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"invalid transaction type": {
			expectedErr: fmt.Errorf("unknown transaction type: Loan"),
			txn: domain.Transaction{
				Type:        "Loan",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    domain.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"currency cannot be empty": {
			expectedErr: fmt.Errorf("currency cannot be empty"),
			txn: domain.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "",
				Description: "some description",
				Date:        time.Now(),
				Category:    domain.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"date cannot be empty": {
			expectedErr: fmt.Errorf("date cannot be empty"),
			txn: domain.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Time{},
				Category:    domain.TrCategoryEntertainment,
				Account:     "some account",
			},
		},
		"category cannot be empty": {
			expectedErr: fmt.Errorf("category cannot be empty"),
			txn: domain.Transaction{
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
			txn: domain.Transaction{
				Type:        "expense",
				IsDebit:     true,
				Amount:      10,
				Currency:    "SGD",
				Description: "some description",
				Date:        time.Now(),
				Category:    domain.TrCategoryEntertainment,
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
