package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/siddhantac/fintra/domain"
	"github.com/siddhantac/fintra/repository"
	"github.com/stretchr/testify/assert"
)

type mockIDGenerator struct {
}

func (mockIDGenerator) NewID() string {
	return "1"
}

func TestGetTransaction(t *testing.T) {
	tests := map[string]struct {
		wantCode     int
		wantRespBody string
	}{
		"valid expense request": {
			wantCode: http.StatusOK,
			wantRespBody: `{
				"id": "1",
				"amount": 23,
				"type": "expense",
				"description": "dinner",
				"date": "2021-08-17",
				"category": "meals",
				"is_debit": true,
				"account": "axis bank"
			}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/transactions/1", nil)
			w := httptest.NewRecorder()

			mockRepo := &repository.MemStore{
				Transactions: map[string]*domain.Transaction{
					"1": {
						ID:          "1",
						Amount:      23,
						Type:        "expense",
						Description: "dinner",
						Date:        time.Date(2021, 8, 17, 0, 0, 0, 0, &time.Location{}),
						Category:    "meals",
						IsDebit:     true,
						Account:     "axis bank",
					},
				},
			}
			handler := GetTransaction(mockRepo)
			handler(w, r)

			assert.Equal(t, test.wantCode, w.Code)
			assert.JSONEq(t, test.wantRespBody, string(w.Body.Bytes()))
		})
	}
}

func TestCreateTransaction(t *testing.T) {

	tests := map[string]struct {
		reqBody      string
		wantCode     int
		wantRespBody string
	}{
		"valid expense request": {
			reqBody: `{
				"id": "1",
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
				"id": "1",
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

			handler := CreateTransaction(&repository.MemStore{}, mockIDGenerator{})
			handler(w, r)

			assert.Equal(t, test.wantCode, w.Code)
			assert.JSONEq(t, test.wantRespBody, string(w.Body.Bytes()))
		})
	}
}
