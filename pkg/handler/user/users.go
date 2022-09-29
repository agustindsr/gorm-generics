package user

import (
	"encoding/json"
	"fmt"
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	if errAdd := h.Service.CreateUser(&user); errAdd != nil {
		fmt.Println(errAdd)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h handler) SearchUser(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var searchOptions common.SearchOptions[models.User]
	json.Unmarshal(body, &searchOptions)

	resp, err := h.Service.SearchUsers(searchOptions)
	if err != nil {
		fmt.Println(err)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
