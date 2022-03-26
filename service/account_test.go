package service

import (
	"testing"

	"github.com/siddhantac/fintra/model"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	tests := map[string]struct {
		name            string
		startingBalance float32
		accRepo         *AccountRepositoryMock
		verifyAccount   func(*testing.T, *model.Account)
		verifyErr       func(*testing.T, error)
		verifyRepoCalls func(*testing.T, *AccountRepositoryMock)
	}{
		"valid account": {
			name:            "AwesomeBank",
			startingBalance: 12,
			accRepo: &AccountRepositoryMock{
				InsertAccountFunc: func(_ string, _ *model.Account) error {
					return nil
				},
			},
			verifyAccount: func(t *testing.T, acc *model.Account) {
				require.False(t, acc.Created.IsZero())
				require.Equal(t, "AwesomeBank", acc.Name)
				require.Equal(t, float32(12.00), acc.StartingBalance.Amount())
				require.Equal(t, float32(12.00), acc.CurrentBalance.Amount())
			},
			verifyErr: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
			verifyRepoCalls: func(t *testing.T, accRepo *AccountRepositoryMock) {
				require.Len(t, accRepo.InsertAccountCalls(), 1)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewAccountService(test.accRepo)
			acc, err := s.NewAccount(test.name, test.startingBalance)
			test.verifyErr(t, err)
			test.verifyAccount(t, acc)
			test.verifyRepoCalls(t, test.accRepo)
		})
	}
}
