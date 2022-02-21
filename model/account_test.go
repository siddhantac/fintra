package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	now := time.Now()
	acc := NewAccount("Citibank", 100)
	assert.Equal(t, "Citibank", acc.Name())
	assert.Equal(t, 100, acc.StartingBalance())
	assert.Equal(t, 100, acc.CurrentBalance())
	assert.WithinDuration(t, now, acc.Created, time.Second*1)
	assert.WithinDuration(t, now, acc.Updated, time.Second*1)
}
