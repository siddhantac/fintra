package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicFlow(t *testing.T) {
	t.Run("listAll", func(t *testing.T) {
		url := baseURL + "/transactions"
		t.Logf("calling %s", url)

		resp, err := http.Get(url)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		fmt.Println(string(body))
		require.Len(t, body, 0)
	})

	// list all transactions, should be 0
	// create a new transaction
	// list the new transaction
	// create another transaction
	// list all transactions, should be 2
}
