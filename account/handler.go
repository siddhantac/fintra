package account

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/siddhantac/fintra/http/rest"
	"github.com/siddhantac/fintra/model"
)

type createAccountRequest struct {
	Name            string  `json:"name"`
	StartingBalance float32 `json:"starting_balance"`
}

type createAccountResponse struct {
	Name            string  `json:"name"`
	StartingBalance float32 `json:"starting_balance"`
	CurrentBalance  float32 `json:"current_balance"`
}

type AccountHandler struct {
	accsvc AccountService
}

type AccountService interface {
	NewAccount(name string, startingBalance float32) (*model.Account, error)
	GetAllAccounts() ([]*model.Account, error)
	GetAccountByName(name string) (*model.Account, error)
}

func NewHandler(accountService AccountService) *AccountHandler {
	return &AccountHandler{accsvc: accountService}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req createAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		msg := fmt.Sprintf("invalid json: %v", err)
		http.Error(w, rest.NewErrorResponse(msg), http.StatusBadRequest)
		return
	}

	account, err := h.accsvc.NewAccount(req.Name, req.StartingBalance)
	if err != nil {
		log.Println(err)
		http.Error(w, rest.NewErrorResponse(err.Error()), http.StatusBadRequest)
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
		http.Error(w, rest.NewErrorResponse(err.Error()), http.StatusInternalServerError)
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

func (h *AccountHandler) GetAccountByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	account, err := h.accsvc.GetAccountByName(name)
	if err != nil {
		log.Println(err)
		http.Error(w, rest.NewErrorResponse(err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newAccountResponse(account)); err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func newAccountResponse(account *model.Account) createAccountResponse {
	return createAccountResponse{
		Name:            account.Name,
		StartingBalance: account.StartingBalance.Amount(),
		CurrentBalance:  account.CurrentBalance.Amount(),
	}
}
