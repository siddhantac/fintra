package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/siddhantac/fintra/model"
)

type CreateTransactionRequest struct {
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	// Currency    string `json:"currency"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Category    string `json:"category"`
	IsDebit     bool   `json:"is_debit"`
	Account     string `json:"account"`
}

type TransactionResponse struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	// Currency    string `json:"currency"`
	Description string `json:"description"`
	Date        Time   `json:"date"`
	Category    string `json:"category"`
	IsDebit     bool   `json:"is_debit"`
	Account     string `json:"account"`
	// Created     Time   `json:"created"`
}

// type Repository interface {
//     Insert(*model.Transaction) error
//     GetByID(string) (*model.Transaction, error)
// }

type IDGenerator interface {
	NewID() string
}

type Handler struct {
	txnsvc TransactionService
}

//go:generate moq -out handler_mock_test.go . TransactionService
type TransactionService interface {
	GetAllTransactions() ([]*model.Transaction, error)
	GetTransaction(id string) (*model.Transaction, error)
	NewTransaction(
		amount float64,
		isDebit bool,
		date, category, transactionType, description, account string) (*model.Transaction, error)
}

func NewHandler(transactionService TransactionService) *Handler {
	return &Handler{txnsvc: transactionService}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"status": "healthy",
	}
	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.txnsvc.GetAllTransactions()
	if err != nil {
		log.Println(err)
		http.Error(w, newErrorResponse(err.Error()), http.StatusInternalServerError)
		return
	}

	txnResps := make([]TransactionResponse, 0, len(transactions))
	for _, txn := range transactions {
		txnResps = append(txnResps, newTransactionResponse(txn))
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(txnResps); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "txnID")
	transaction, err := h.txnsvc.GetTransaction(id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, model.ErrNotFound) {
			http.Error(w, newErrorResponse(err.Error()), http.StatusNotFound)
			return
		}
		http.Error(w, newErrorResponse(err.Error()), http.StatusBadRequest)
		return
	}

	if transaction == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	resp := newTransactionResponse(transaction)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var ctr CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&ctr); err != nil {
		log.Println(err)
		msg := fmt.Sprintf("invalid json: %v", err)
		http.Error(w, newErrorResponse(msg), http.StatusBadRequest)
		return
	}

	transaction, err := h.txnsvc.NewTransaction(ctr.Amount, ctr.IsDebit, ctr.Date, ctr.Category, ctr.Type, ctr.Description, ctr.Account)
	if err != nil {
		log.Println(err)
		http.Error(w, newErrorResponse(err.Error()), http.StatusBadRequest)
		return
	}

	resp := newTransactionResponse(transaction)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(msg string) string {
	resp := errorResponse{Error: msg}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return "internal error"
	}
	return string(b)
}

func newTransactionResponse(t *model.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:     t.ID,
		Amount: t.Amount,
		Type:   string(t.Type),
		// Currency:    string(transaction.Currency),
		Description: t.Description,
		Date:        NewTime(t.Date),
		Category:    string(t.Category),
		IsDebit:     t.IsDebit,
		Account:     t.Account,
		// Created:     NewTime(transaction.Created),
	}
}

const (
	dateLayout = "2006-01-02"
)

type Time string

func NewTime(t time.Time) Time {
	return Time(t.Round(time.Second).In(time.UTC).Format(dateLayout))
}
