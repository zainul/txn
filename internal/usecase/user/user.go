package user

import (
	"fmt"

	"github.com/zainul/txn/internal/entity"
	"github.com/zainul/txn/internal/repository"
)

// Usecase ...
type Usecase struct {
	UserAccountRepo repository.UserAccount
}

// SeedUser is seed the user account
func (u *Usecase) SeedUser(n int) error {
	for i := 2; i < n+2; i++ {
		user := entity.UserAccount{
			AccountNumber: int64(i),
			Balance:       (1000 * 1000),
		}
		err := u.UserAccountRepo.Insert(user)

		if err != nil {
			fmt.Println(fmt.Sprintf("Failed store : %#v", user))
		}
	}

	return nil
}
