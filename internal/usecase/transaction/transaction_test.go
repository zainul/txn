package transaction_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zainul/txn/internal/contract"
	"github.com/zainul/txn/internal/entity"
	"github.com/zainul/txn/internal/pkg/error/deliveryerror"
	"github.com/zainul/txn/internal/usecase"
)

func TestGetBalance(t *testing.T) {

	u := usecase.NewTransaction(mHist, mLog, mUser)

	res, err := u.GetBalance(111)

	errDeli := deliveryerror.GetError("errCode", nil)
	assert.Equal(t, err, errDeli)
	assert.Equal(t, res.LastBalance, float64(10))
}

func TestValidBalancePreTransaction(t *testing.T) {
	u := usecase.NewTransaction(mHist, mLog, mUser)

	valid := u.ValidBalancePreTransaction(entity.UserAccount{
		Balance: 10,
	}, 9)

	InValid := u.ValidBalancePreTransaction(entity.UserAccount{
		Balance: 10,
	}, 10.1)

	assert.Equal(t, true, valid)
	assert.Equal(t, false, InValid)
}

func TestInternalTransfer(t *testing.T) {
	u := usecase.NewTransaction(mHist, mLog, mUser)

	res, err := u.InternalTransfer(contract.TransactionRequest{
		Amount:            10,
		FromAccountNumber: 10,
		ToAccountNumber:   3,
	})

	errDeli := deliveryerror.GetError("errCode", nil)
	assert.Equal(t, err, errDeli)
	assert.NotEqual(t, "", res.TransactionID)
}
