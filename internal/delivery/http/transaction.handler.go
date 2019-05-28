package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zainul/txn/internal/contract"

	"github.com/zainul/txn/internal/pkg/error/deliveryerror"
	"github.com/zainul/txn/internal/pkg/sendoutput"
	"github.com/zainul/txn/internal/usecase"
)

// TxHandler ...
type TxHandler struct {
	TxUsecase usecase.Transaction
}

func (t *TxHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	response := sendoutput.Response{}

	vars := mux.Vars(r)
	id := vars["account_number"]

	idInt, err := strconv.Atoi(id)

	if err != nil {
		err := deliveryerror.GetError(deliveryerror.BadRequest, errors.New("bad request"))
		sendoutput.Send(r, w, response, err)
		return
	}

	res, errDeli := t.TxUsecase.GetBalance(int64(idInt))

	if errDeli != nil {
		response.Error = errDeli
		sendoutput.Send(r, w, response, errDeli)
		return
	}

	response.Data = res
	sendoutput.Send(r, w, response, nil)
	return
}

// InternalTx ...
func (t *TxHandler) InternalTx(w http.ResponseWriter, r *http.Request) {
	response := sendoutput.Response{}

	body := contract.TransactionRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	if err != nil {
		err := deliveryerror.GetError(deliveryerror.BadRequest, err)
		sendoutput.Send(r, w, response, err)
		return
	}

	res, errDeli := t.TxUsecase.InternalTransfer(body)

	if errDeli != nil {
		response.Error = errDeli
		sendoutput.Send(r, w, response, errDeli)
		return
	}

	response.Data = res
	sendoutput.Send(r, w, response, nil)
	return
}
