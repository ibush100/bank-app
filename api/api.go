package api

import (
	"bank-app/helpers"
	"bank-app/transaction"
	"bank-app/users"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Call successful dude")
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/", index)
	router.HandleFunc("/login", users.LoginUser).Methods("POST")
	router.HandleFunc("/user", users.RegisterUser).Methods("POST")
	router.HandleFunc("/user", users.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user", users.UpdateUserEmail).Methods("PUT")
	router.HandleFunc("/transaction", transaction.CreateTransactionHandler).Methods("POST")
	router.HandleFunc("/updateBalance", transaction.UpdateUserBalance).Methods("PUT")

	fmt.Println("App is working on port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
