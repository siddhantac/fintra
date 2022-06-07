package transfer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/siddhantac/fintra/http/rest"
)

type Handler struct {
}

type CreateTransferRequest struct {
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) RecordTransfer(w http.ResponseWriter, r *http.Request) {
	var ctr CreateTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&ctr); err != nil {
		log.Println(err)
		msg := fmt.Sprintf("invalid json: %v", err)
		http.Error(w, rest.NewErrorResponse(msg), http.StatusBadRequest)
		return
	}
}
