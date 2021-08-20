package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
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

type Time string

const (
	dateLayout = "2006-01-02"
)

func NewTime(t time.Time) Time {
	return Time(t.Round(time.Second).In(time.UTC).Format(dateLayout))
}

type Repository interface {
	Insert(*Transaction)
	Get(string) *Transaction
}

func GetTransaction(repo Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/transactions/")
		transaction := repo.Get(id)
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
}

func CreateTransaction(repo Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctr CreateTransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&ctr); err != nil {
			log.Println(err)
			http.Error(w, newErrorResponse("invalid JSON"), http.StatusBadRequest)
			return
		}

		date, err := time.Parse(dateLayout, ctr.Date)
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid date", http.StatusBadRequest)
			return
		}

		transaction, err := NewTransaction(ctr.Amount, date, ctr.IsDebit, ctr.Category, ctr.Type, ctr.Description, ctr.Account)
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

func newTransactionResponse(t *Transaction) TransactionResponse {
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
