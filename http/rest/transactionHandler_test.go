package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/siddhantac/fintra/model"
	"github.com/siddhantac/fintra/money"
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
		mockSvc      *TransactionServiceMock
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
				"amount": 99.5,
				"type": "expense",
				"description": "mrt",
				"date": "2021-08-20",
				"category": "transport",
				"is_debit": true,
				"account": "credit card"
			}]`,
			mockSvc: &TransactionServiceMock{
				GetAllTransactionsFunc: func() ([]*model.Transaction, error) {
					return []*model.Transaction{
						{
							ID:          "1",
							Amount:      money.NewMoney(23),
							Type:        "expense",
							Description: "dinner",
							Date:        time.Date(2021, 8, 17, 0, 0, 0, 0, &time.Location{}),
							Category:    "meals",
							IsDebit:     true,
							Account:     "axis bank",
						},
						{
							ID:          "2",
							Amount:      money.NewMoney(99.5),
							Type:        "expense",
							Description: "mrt",
							Date:        time.Date(2021, 8, 20, 0, 0, 0, 0, &time.Location{}),
							Category:    "transport",
							IsDebit:     true,
							Account:     "credit card",
						},
					}, nil
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/transactions", nil)
			w := httptest.NewRecorder()

			handler := NewTransactionHandler(test.mockSvc)
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
		mockSvc      *TransactionServiceMock
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
			mockSvc: &TransactionServiceMock{
				GetTransactionFunc: func(id string) (*model.Transaction, error) {
					return &model.Transaction{
						ID:          "1",
						Amount:      money.NewMoney(23),
						Type:        "expense",
						Description: "dinner",
						Date:        time.Date(2021, 8, 17, 0, 0, 0, 0, &time.Location{}),
						Category:    "meals",
						IsDebit:     true,
						Account:     "axis bank",
					}, nil
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/transactions/1", nil)
			w := httptest.NewRecorder()

			handler := NewTransactionHandler(test.mockSvc)
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
			wantCode:     http.StatusCreated,
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
			wantCode: http.StatusBadRequest,
			compareResp: func(t *testing.T, m map[string]interface{}) {
				require.Contains(t, m, "error")
				require.Contains(t, m["error"], "invalid json")
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.reqBody))
			w := httptest.NewRecorder()

			mockSvc := &TransactionServiceMock{
				NewTransactionFunc: func(amount float32, isDebit bool, date, category, transactionType, description, account string) (*model.Transaction, error) {
					d := time.Date(2021, 8, 17, 0, 0, 0, 0, time.UTC)
					return model.NewTransaction("1", money.NewMoney(amount), d, isDebit, category, transactionType, description, account), nil
				},
			}
			handler := NewTransactionHandler(mockSvc)
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
