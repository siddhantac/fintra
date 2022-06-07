package account

import (
	"testing"
	"time"

	"github.com/siddhantac/fintra/money"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	now := time.Now()
	acc := NewAccount("Citibank", money.NewMoney(100))
	assert.Equal(t, "Citibank", acc.Name)
	assert.Equal(t, float32(100), acc.StartingBalance.Amount())
	assert.Equal(t, float32(100), acc.CurrentBalance.Amount())
	assert.WithinDuration(t, now, acc.Created, time.Second*1)
	assert.WithinDuration(t, now, acc.Updated, time.Second*1)
}
