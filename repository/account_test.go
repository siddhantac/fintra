package repository

import (
	"testing"

	"github.com/siddhantac/fintra/infra/store"
	"github.com/siddhantac/fintra/model"
	"github.com/stretchr/testify/require"
)

func TestGetAccountDetails(t *testing.T) {
	storage := store.NewMemStore()
	expectedAcc := model.NewAccount("Citibank", 150)
	// expectedAcc := &model.Account{
	// 	ID:              "FakeBankAccount",
	// 	Balance:         150,
	// 	StartingBalance: 15,
	// 	Name:            "Citibank",
	// 	Created:         time.Now(),
	// 	Updated:         time.Now(),
	// }

	storage.Items = map[string]interface{}{
		"FakeBankAccount": expectedAcc,
	}

	repo := NewAccountRepository(storage)
	gotAcc, err := repo.GetByName("FakeBankAccount")
	require.NoError(t, err)
	require.Equal(t, expectedAcc, gotAcc)
}

/*
func TestCreditToAccount(t *testing.T) {
	storage := store.NewMemStore()
	initialState := &model.Account{
		ID:              "FakeBankAccount",
		Balance:         150,
		StartingBalance: 15,
		Name:            "Citibank",
		Created:         time.Now(),
		Updated:         time.Now(),
	}

	storage.Items = map[string]interface{}{
		"FakeBankAccount": initialState,
	}

	repo := NewAccountRepository(storage)

	balance, err := repo.CreditToAccount("FakeBankAccount", 48)
	require.NoError(t, err)
	require.Equal(t, int64(198), balance)
}

func TestDebitFromAccount(t *testing.T) {
	storage := store.NewMemStore()
	initialState := &model.Account{
		ID:              "FakeBankAccount",
		Balance:         150,
		StartingBalance: 15,
		Name:            "Citibank",
		Created:         time.Now(),
		Updated:         time.Now(),
	}

	storage.Items = map[string]interface{}{
		"FakeBankAccount": initialState,
	}

	repo := NewAccountRepository(storage)

	balance, err := repo.DebitFromAccount("FakeBankAccount", 48)
	require.NoError(t, err)
	require.Equal(t, int64(102), balance)
}

*/
