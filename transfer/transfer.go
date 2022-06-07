package transfer

type TransferRequest struct {
	Amount      float32 `json:"amount"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
}
