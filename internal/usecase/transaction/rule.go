package transaction

import (
	"strings"

	"github.com/zainul/txn/internal/constant"
	"github.com/zainul/txn/internal/entity"
)

func makeRule(fromAcc entity.UserAccount, toAcc entity.UserAccount, transactionCode string, amount float64, txID string) []entity.TransactionLog {
	txLog := make([]entity.TransactionLog, 0)
	txCode := strings.ToUpper(transactionCode)

	if txCode == constant.TransferMemberToMemberCode {
		fromTxLog := entity.TransactionLog{
			DrCr:                  constant.DR,
			PreviousBalance:       fromAcc.Balance,
			AccountNumber:         fromAcc.AccountNumber,
			TransactionCode:       transactionCode,
			Amount:                amount,
			EndBalance:            fromAcc.Balance,
			OpponentAccountNumber: toAcc.AccountNumber,
			TransactionStatus:     constant.TransactionStatusPending,
			TxID:                  txID,
		}

		txLog = append(txLog, fromTxLog)

		toTxLog := entity.TransactionLog{
			DrCr:                  constant.CR,
			PreviousBalance:       toAcc.Balance,
			AccountNumber:         toAcc.AccountNumber,
			TransactionCode:       transactionCode,
			Amount:                amount,
			EndBalance:            toAcc.Balance,
			OpponentAccountNumber: fromAcc.AccountNumber,
			TransactionStatus:     constant.TransactionStatusPending,
			TxID:                  txID,
		}

		txLog = append(txLog, toTxLog)

		return txLog
	}

	return txLog
}
