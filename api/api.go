package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/siddhantac/fintra/domain"
)

type CreateTransactionRequest struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
	// Currency    string `json:"currency"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Category    string `json:"category"`
	IsDebit     bool   `json:"is_debit"`
	Account     string `json:"account"`
}

type TransactionResponse struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
	// Currency    string `json:"currency"`
	Description string `json:"description"`
	Date        Time   `json:"date"`
	Category    string `json:"category"`
	IsDebit     bool   `json:"is_debit"`
	Account     string `json:"account"`
	// Created     Time   `json:"created"`
}

type Repository interface {
	Insert(*domain.Transaction) error
	GetByID(string) (*domain.Transaction, error)
}

type IDGenerator interface {
	NewID() string
}

type Handler struct {
	service Service
}

//go:generate moq -out api_mock_test.go . Service
type Service interface {
	GetTransaction(id string) (*domain.Transaction, error)
	NewTransaction(
		amount int,
		isDebit bool,
		date, category, transactionType, description, account string) (*domain.Transaction, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
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

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/transactions/")
	transaction, err := h.service.GetTransaction(id)
	if err != nil {
		log.Println(err)
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
		http.Error(w, newErrorResponse("invalid JSON"), http.StatusBadRequest)
		return
	}

	transaction, err := h.service.NewTransaction(ctr.Amount, ctr.IsDebit, ctr.Date, ctr.Category, ctr.Type, ctr.Description, ctr.Account)
	if err != nil {
		log.Println(err)
		http.Error(w, newErrorResponse(err.Error()), http.StatusBadRequest)
		return
	}

	resp := newTransactionResponse(transaction)
	w.WriteHeader(http.StatusOK)
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

func newTransactionResponse(t *domain.Transaction) TransactionResponse {
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
