package http

import (
	"net/http"

	"github.com/zainul/txn/internal/pkg/sendoutput"
	"github.com/zainul/txn/internal/usecase"
)

// UserHandler ...
type UserHandler struct {
	UserUsecase usecase.User
}

// SeedOneMillion ...
func (u *UserHandler) SeedOneMillion(w http.ResponseWriter, r *http.Request) {
	response := sendoutput.Response{}
	u.UserUsecase.SeedUser(1000 * 1000)
	sendoutput.Send(r, w, response, nil)
	return
}
