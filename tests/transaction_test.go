//go:build integration

package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicFlow(t *testing.T) {
	var txnID string

	t.Run("listAll", func(t *testing.T) {
		url := baseURL + "/transactions"
		t.Logf("calling %s", url)

		resp, err := http.Get(url)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var m []interface{}
		err = json.Unmarshal(body, &m)
		require.NoError(t, err)
		require.Equal(t, []interface{}{}, m)
	})

	t.Run("create new", func(t *testing.T) {
		url := baseURL + "/transactions"
		t.Logf("calling %s", url)

		reqBody := strings.NewReader(`{"amount": 13.5, "description": "Tasty Restaurant", "category": "meals", "date": "2021-12-23", "type": "expense", "is_debit": true, "account": "awesome bank"}`)
		resp, err := http.Post(url, "application/json", reqBody)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)

		var m map[string]interface{}
		err = json.Unmarshal(body, &m)
		require.NoError(t, err)

		txnID = m["id"].(string)
		isEqualTransactionResponse(t, newTransactionResponse(), m, "id")
	})

	t.Run("list the new transaction", func(t *testing.T) {
		url := baseURL + "/transactions/" + txnID
		t.Logf("calling %s", url)

		resp, err := http.Get(url)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		var m map[string]interface{}
		err = json.Unmarshal(body, &m)
		require.NoError(t, err)
		isEqualTransactionResponse(t, newTransactionResponse(), m, "id")
	})

	t.Run("list all transactions 2", func(t *testing.T) {
		url := baseURL + "/transactions"
		t.Logf("calling %s", url)

		// create another transaction
		reqBody := strings.NewReader(`{"amount": 13.5, "description": "Tasty Restaurant", "category": "meals", "date": "2021-12-23", "type": "expense", "is_debit": true, "account": "awesome bank"}`)
		resp, err := http.Post(url, "application/json", reqBody)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)

		resp, err = http.Get(url)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var m []interface{}
		err = json.Unmarshal(body, &m)
		require.NoError(t, err)
		require.Len(t, m, 2)

		m0 := m[0].(map[string]interface{})
		m1 := m[1].(map[string]interface{})
		isEqualTransactionResponse(t, newTransactionResponse(), m0, "id")
		isEqualTransactionResponse(t, newTransactionResponse(), m1, "id")
	})
}

func isEqualTransactionResponse(t *testing.T, expected, got map[string]interface{}, ignoreFields ...string) {
	t.Helper()

	// assert the ignored field is not empty, then clear it
	for _, f := range ignoreFields {
		require.NotEmpty(t, got[f])
		got[f] = ""
	}
	require.Equal(t, expected, got)
}

func newTransactionResponse() map[string]interface{} {
	return map[string]interface{}{
		"account":     "awesome bank",
		"amount":      13.50,
		"category":    "meals",
		"date":        "2021-12-23",
		"id":          "",
		"description": "Tasty Restaurant",
		"is_debit":    true,
		"type":        "expense",
	}
}
