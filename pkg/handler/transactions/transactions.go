package transactions

import (
	"encoding/json"
	"fmt"
	"gorm-with-generics/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var transaction models.TransactionLedger
	json.Unmarshal(body, &transaction)

	if errAdd := h.Service.CreateTransaction(&transaction); errAdd != nil {
		fmt.Println(errAdd)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h handler) SearchTransaction(w http.ResponseWriter, r *http.Request) {
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := h.Service.SearchTransactionByAmount(amount)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
