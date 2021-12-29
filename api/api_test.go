package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/siddhantac/fintra/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockIDGenerator struct {
}

func (mockIDGenerator) NewID() string {
	return "1"
}

func TestGetAllTransactions(t *testing.T) {
	tests := map[string]struct {
		wantCode     int
		wantRespBody string
	}{
		"valid expense request": {
			wantCode: http.StatusOK,
			wantRespBody: `[{
				"id": "1",
				"amount": 23,
				"type": "expense",
				"description": "dinner",
				"date": "2021-08-17",
				"category": "meals",
				"is_debit": true,
				"account": "axis bank"
			},
			{
				"id": "2",
				"amount": 99,
				"type": "expense",
				"description": "mrt",
				"date": "2021-08-20",
				"category": "transport",
				"is_debit": true,
				"account": "credit card"
			}]`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/transactions", nil)
			w := httptest.NewRecorder()

			mockSvc := &ServiceMock{
				GetAllTransactionsFunc: func() ([]*domain.Transaction, error) {
					return []*domain.Transaction{
						{
							ID:          "1",
							Amount:      23,
							Type:        "expense",
							Description: "dinner",
							Date:        time.Date(2021, 8, 17, 0, 0, 0, 0, &time.Location{}),
							Category:    "meals",
							IsDebit:     true,
							Account:     "axis bank",
						},
						{
							ID:          "2",
							Amount:      99,
							Type:        "expense",
							Description: "mrt",
							Date:        time.Date(2021, 8, 20, 0, 0, 0, 0, &time.Location{}),
							Category:    "transport",
							IsDebit:     true,
							Account:     "credit card",
						},
					}, nil
				},
			}

			handler := NewHandler(mockSvc)
			handler.GetAllTransactions(w, r)

			assert.Equal(t, test.wantCode, w.Code)
			assert.JSONEq(t, test.wantRespBody, string(w.Body.Bytes()))
		})
	}
}

func TestGetTransactionByID(t *testing.T) {
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

			mockSvc := &ServiceMock{
				GetTransactionFunc: func(id string) (*domain.Transaction, error) {
					return &domain.Transaction{
						ID:          "1",
						Amount:      23,
						Type:        "expense",
						Description: "dinner",
						Date:        time.Date(2021, 8, 17, 0, 0, 0, 0, &time.Location{}),
						Category:    "meals",
						IsDebit:     true,
						Account:     "axis bank",
					}, nil
				},
			}

			handler := NewHandler(mockSvc)
			handler.GetTransactionByID(w, r)

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
		expectedResp map[string]interface{}
		compareResp  func(t *testing.T, m map[string]interface{})
		serviceCalls int
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
			wantCode:     http.StatusOK,
			serviceCalls: 1,
			compareResp: func(t *testing.T, m map[string]interface{}) {
				t.Helper()
				require.Contains(t, m, "id")

				require.Contains(t, m, "amount")
				require.Equal(t, m["amount"], 23.00)
				require.Contains(t, m, "type")
				require.Equal(t, m["type"], "expense")
				require.Contains(t, m, "description")
				require.Equal(t, m["description"], "dinner")
				require.Contains(t, m, "date")
				require.Equal(t, m["date"], "2021-08-17")
				require.Contains(t, m, "category")
				require.Equal(t, m["category"], "meals")
				require.Contains(t, m, "is_debit")
				require.Equal(t, m["is_debit"], true)
				require.Contains(t, m, "account")
				require.Equal(t, m["account"], "axis bank")
			},
		},
		"invalid json": {
			reqBody: `{
				"amount": 23,
			}`,
			wantCode:     http.StatusBadRequest,
			expectedResp: map[string]interface{}{"error": "invalid JSON"},
			compareResp: func(t *testing.T, m map[string]interface{}) {
				require.Contains(t, m, "error")
				require.Equal(t, m["error"], "invalid JSON")
			},
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
			serviceCalls: 1,
			wantCode:     http.StatusBadRequest,
			expectedResp: map[string]interface{}{"error": "income must be credit"},
			compareResp: func(t *testing.T, m map[string]interface{}) {
				require.Contains(t, m, "error")
				require.Equal(t, m["error"], "income must be credit")
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.reqBody))
			w := httptest.NewRecorder()

			mockSvc := &ServiceMock{
				NewTransactionFunc: func(amount int, isDebit bool, date, category, transactionType, description, account string) (*domain.Transaction, error) {
					d := time.Date(2021, 8, 17, 0, 0, 0, 0, time.UTC)
					return domain.NewTransaction(amount, d, isDebit, category, transactionType, description, account)
				},
			}
			handler := NewHandler(mockSvc)
			handler.CreateTransaction(w, r)

			assert.Equal(t, test.wantCode, w.Code)
			var m map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &m)
			require.NoError(t, err)
			test.compareResp(t, m)
			require.Len(t, mockSvc.NewTransactionCalls(), test.serviceCalls)
		})
	}
}
