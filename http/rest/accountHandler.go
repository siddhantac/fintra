package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/siddhantac/fintra/model"
)

type createAccountRequest struct {
	Name            string `json:"name"`
	StartingBalance int    `json:"starting_balance"`
}

type createAccountResponse struct {
	Name            string `json:"name"`
	StartingBalance int    `json:"starting_balance"`
	CurrentBalance  int    `json:"current_balance"`
}

type AccountHandler struct {
	accsvc AccountService
}

type AccountService interface {
	NewAccount(name string, startingBalance int) (*model.Account, error)
	GetAllAccounts() ([]*model.Account, error)
}

func NewAccountHandler(accountService AccountService) *AccountHandler {
	return &AccountHandler{accsvc: accountService}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req createAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		msg := fmt.Sprintf("invalid json: %v", err)
		http.Error(w, newErrorResponse(msg), http.StatusBadRequest)
		return
	}

	account, err := h.accsvc.NewAccount(req.Name, req.StartingBalance)
	if err != nil {
		log.Println(err)
		http.Error(w, newErrorResponse(err.Error()), http.StatusBadRequest)
		return
	}

	resp := newAccountResponse(account)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (h *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accsvc.GetAllAccounts()
	if err != nil {
		log.Println(err)
		http.Error(w, newErrorResponse(err.Error()), http.StatusInternalServerError)
		return
	}

	accResps := make([]createAccountResponse, 0, len(accounts))
	for _, acc := range accounts {
		accResps = append(accResps, newAccountResponse(acc))
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(accResps); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func newAccountResponse(account *model.Account) createAccountResponse {
	return createAccountResponse{
		Name:            account.Name(),
		StartingBalance: account.StartingBalance(),
		CurrentBalance:  account.CurrentBalance(),
	}
}
