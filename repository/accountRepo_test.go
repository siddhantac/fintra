package repository

import (
	"testing"

	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/infra/store"
	"github.com/stretchr/testify/require"
)

func TestGetAccountDetails(t *testing.T) {
	storage := store.NewMemStore()
	expectedAcc := model.NewAccount("Citibank", 150)
	// expectedAcc := &model.Account{
	// 	ID:              "AccID",
	// 	Balance:         150,
	// 	StartingBalance: 15,
	// 	Name:            "Citibank",
	// 	Created:         time.Now(),
	// 	Updated:         time.Now(),
	// }

	storage.Items = map[string]interface{}{
		"AccID": expectedAcc,
	}

	repo := NewAccountRepository(storage)
	gotAcc, err := repo.GetByID("AccID")
	require.NoError(t, err)
	require.Equal(t, expectedAcc, gotAcc)
}

/*
func TestCreditToAccount(t *testing.T) {
	storage := store.NewMemStore()
	initialState := &model.Account{
		ID:              "AccID",
		Balance:         150,
		StartingBalance: 15,
		Name:            "Citibank",
		Created:         time.Now(),
		Updated:         time.Now(),
	}

	storage.Items = map[string]interface{}{
		"AccID": initialState,
	}

	repo := NewAccountRepository(storage)

	balance, err := repo.CreditToAccount("AccID", 48)
	require.NoError(t, err)
	require.Equal(t, int64(198), balance)
}

func TestDebitFromAccount(t *testing.T) {
	storage := store.NewMemStore()
	initialState := &model.Account{
		ID:              "AccID",
		Balance:         150,
		StartingBalance: 15,
		Name:            "Citibank",
		Created:         time.Now(),
		Updated:         time.Now(),
	}

	storage.Items = map[string]interface{}{
		"AccID": initialState,
	}

	repo := NewAccountRepository(storage)

	balance, err := repo.DebitFromAccount("AccID", 48)
	require.NoError(t, err)
	require.Equal(t, int64(102), balance)
}

*/
