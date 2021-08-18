package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	body := `{
		"amount": 23,
		"type": "expense",
		"description": "dinner",
		"date": "2021-08-17",
		"category": "meals",
		"is_debit": true,
		"account": "axis bank"
	}`

	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	w := httptest.NewRecorder()

	CreateTransaction(w, r)

	fmt.Println(w.Code)
	fmt.Println(string(w.Body.Bytes()))
}
