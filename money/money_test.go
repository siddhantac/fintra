package money

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoney(t *testing.T) {
	m := NewMoney(32.50)
	require.Equal(t, float32(32.5), m.Amount())
	require.Equal(t, 3250, m.AmountAsInt())
}

func TestMoneyAdd(t *testing.T) {
	m := NewMoney(32.50)
	added := m.Add(NewMoney(1.90))
	require.Equal(t, float32(34.4), added.Amount())
}

func TestMoneySubtract(t *testing.T) {
	m := NewMoney(32.50)
	subtracted := m.Subtract(NewMoney(1.90))
	require.Equal(t, float32(30.6), subtracted.Amount())
}

func TestMoneyString(t *testing.T) {
	m := NewMoney(32.5)
	require.Equal(t, "32.5", m.String())
}

func TestMoneyJSON(t *testing.T) {
	m := NewMoney(18.54)
	marshalled, err := json.Marshal(m)
	require.NoError(t, err)
	require.Equal(t, []byte(`18.54`), marshalled)

	var money Money
	err = json.Unmarshal(marshalled, &money)
	require.NoError(t, err)
	require.Equal(t, m, money)
}
