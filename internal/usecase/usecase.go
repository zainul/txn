package usecase

import (
	"github.com/zainul/txn/internal/contract"
)

// TransactionUseCase ...
type TransactionUseCase interface {
	Transaction(txnParam contract.TransactionRequest) (contract.TransactionResponse, error)
	GetBalance(accountNumber int64) (contract.GetBalanceResponse, error)
}
