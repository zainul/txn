package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zainul/txn/internal/usecase"
)

// NewUserHanlder ...
func NewUserHanlder(route *mux.Router, user usecase.User) {

	handler := UserHandler{
		UserUsecase: user,
	}

	route.HandleFunc("/user/seed", handler.SeedOneMillion).Methods(http.MethodGet)
}

// NewTxHanlder ...
func NewTxHanlder(route *mux.Router, tx usecase.Transaction) {

	handler := TxHandler{
		TxUsecase: tx,
	}

	route.HandleFunc("/transfer/internal", handler.InternalTx).Methods(http.MethodPost)
}
