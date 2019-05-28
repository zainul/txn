package transaction_test

import (
	"testing"

	"github.com/zainul/txn/internal/entity"
)

type MockUserAccount struct{}

func (m MockUserAccount) Insert(p entity.UserAccount) error {
	return nil
}

func (m MockUserAccount) UserAccountBy(field string, value interface{}) ([]entity.UserAccount, error) {
	acc := entity.UserAccount{
		AccountNumber: 111,
		Balance:       10,
	}

	acs := make([]entity.UserAccount, 0)

	acs = append(acs, acc)
	return acs, nil
}

// TransactionHistory ...
type MockTransactionHistory struct{}

// TransactionLog ...
type MockTransactionLog struct{}

func (m MockTransactionLog) Transfer(e []entity.TransactionLog) error {
	return nil
}

var (
	mHist MockTransactionHistory
	mLog  MockTransactionLog
	mUser MockUserAccount
)

func TestMain(m *testing.M) {
	mHist = MockTransactionHistory{}
	mLog = MockTransactionLog{}
	mUser = MockUserAccount{}
}
