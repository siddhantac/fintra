package rest

import (
	"net/http"

	"github.com/siddhantac/fintra/model"
)

type TransferService interface {
	NewTransaction(amount float32, isDebit bool, date, category, transactionType, description, account string) (*model.Transaction, error)
}

type TransferHandler struct {
	tfrsvc TransferService
}

type TransferRequest struct {
	Amount      float32 `json:"amount"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
}

func NewTransactionHandler(transferService TransferService) *TransferHandler {
	return &TransferHandler{
		tfrsvc: transferService,
	}
}

func (h *TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {

}
