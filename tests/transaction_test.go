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
		require.Equal(t, http.StatusOK, resp.StatusCode)

		var m map[string]interface{}
		err = json.Unmarshal(body, &m)
		require.NoError(t, err)
		require.Equal(t, []interface{}{}, m)

	})
	// list the new transaction
	// create another transaction
	// list all transactions, should be 2
}
