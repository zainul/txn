package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zainul/txn/internal/entity"
)

func TestMakeRule(t *testing.T) {
	txLog := makeRule(entity.UserAccount{
		AccountNumber: 10,
	}, entity.UserAccount{
		AccountNumber: 100,
	}, "TMM", 100, "SomeTXID")

	txLogNoRow := makeRule(entity.UserAccount{
		AccountNumber: 10,
	}, entity.UserAccount{
		AccountNumber: 100,
	}, "TMMI", 100, "SomeTXID")

	assert.Equal(t, 2, len(txLog))
	assert.Equal(t, 0, len(txLogNoRow))
}
