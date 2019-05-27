package sendoutput

import (
	"encoding/json"
	"net/http"

	"github.com/zainul/txn/internal/pkg/error/deliveryerror"
)

// Response ...
type Response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

// Write is to make response of http
func Write(w http.ResponseWriter, response interface{}, statusCode ...int) {

	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	}
	json.NewEncoder(w).Encode(response)
}

// Send ...
func Send(r *http.Request, w http.ResponseWriter, response Response, errDelivery *deliveryerror.Error) {

	w.Header().Set("Content-Type", "application/json")

	if errDelivery != nil {
		response.Error = errDelivery
		Write(w, response, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	Write(w, response)
}
