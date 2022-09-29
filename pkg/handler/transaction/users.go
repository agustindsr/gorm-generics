package transaction

import (
	"encoding/json"
	"fmt"
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
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
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var searchOptions common.SearchOptions[models.TransactionLedger]
	json.Unmarshal(body, &searchOptions)

	resp, err := h.Service.SearchTransaction(searchOptions)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
