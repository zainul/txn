package contract

// TransactionRequest ...
type TransactionRequest struct {
	ToAccountNumber   int64   `json:"to_account_number"`
	FromAccountNumber int64   `json:"from_account_number"`
	Amount            float64 `json:"amount"`
}

// TransactionResponse ...
type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
}

// GetBalanceResponse ...
type GetBalanceResponse struct {
	AccountNumber int64   `json:"account_number"`
	LastBalance   float64 `json:"last_balance"`
}
