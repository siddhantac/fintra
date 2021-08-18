package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Date        Time   `json:"date"`
	Category    string `json:"category"`
	IsDebit     bool   `json:"is_debit"`
	Account     string `json:"account"`
	Created     Time   `json:"created"`
}

type Time string

func NewTime(t time.Time) Time {
	return Time(t.Round(time.Second).In(time.UTC).Format(time.RFC3339))
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var ctr CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&ctr); err != nil {
		log.Println(err)
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, ctr.Date)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	transaction, err := NewTransaction(ctr.Amount, date, ctr.IsDebit, ctr.Category, ctr.Type, ctr.Description, ctr.Account)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := TransactionResponse{
		Amount:      transaction.Amount,
		Type:        string(transaction.Type),
		Currency:    string(transaction.Currency),
		Description: transaction.Description,
		Date:        NewTime(transaction.Date),
		Category:    string(transaction.Category),
		IsDebit:     transaction.IsDebit,
		Account:     transaction.Account,
		Created:     NewTime(transaction.Created),
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
