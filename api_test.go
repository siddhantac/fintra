package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	tests := map[string]struct {
		reqBody      string
		wantCode     int
		wantRespBody string
	}{
		"valid expense request": {
			reqBody: `{
				"amount": 23,
				"type": "expense",
				"description": "dinner",
				"date": "2021-08-17",
				"category": "meals",
				"is_debit": true,
				"account": "axis bank"
			}`,
			wantCode: http.StatusOK,
			wantRespBody: `{
				"amount": 23,
				"type": "expense",
				"description": "dinner",
				"date": "2021-08-17",
				"category": "meals",
				"is_debit": true,
				"account": "axis bank"
			}`,
		},
		"invalid json": {
			reqBody: `{
				"amount": 23,
			}`,
			wantCode:     http.StatusBadRequest,
			wantRespBody: `{"error": "invalid JSON"}`,
		},
		"invalid request": {
			reqBody: `{
				"amount": 23,
				"type": "income",
				"description": "dinner",
				"date": "2021-08-17",
				"category": "meals",
				"is_debit": true,
				"account": "axis bank"
			}`,
			wantCode:     http.StatusBadRequest,
			wantRespBody: `{"error": "income must be credit"}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.reqBody))
			w := httptest.NewRecorder()

			CreateTransaction(w, r)

			assert.Equal(t, test.wantCode, w.Code)
			assert.JSONEq(t, test.wantRespBody, string(w.Body.Bytes()))
		})
	}
}
