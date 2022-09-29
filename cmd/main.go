package main

import (
	"github.com/gorilla/mux"
	"gorm-with-generics/pkg/handler/transactions"
	"gorm-with-generics/pkg/handler/user"
	"log"
	"net/http"
)

func main() {
	userHandler := user.New()
	transactionHandler := transactions.New()

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", userHandler.SearchUsersByFirstName).Methods(http.MethodGet)

	router.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods(http.MethodPost)
	router.HandleFunc("/transactions", transactionHandler.SearchTransaction).Methods(http.MethodGet)
	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
