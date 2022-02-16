package service

import (
	"testing"
	"time"

	"github.com/siddhantac/fintra/domain"
	"github.com/stretchr/testify/assert"
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
