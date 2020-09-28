package api

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"bank-app/transaction"
	"bank-app/users"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Call successful dude")
}

func PrepareToken(ID uint) string {
	tokenContent := jwt.MapClaims{
		"user_id": ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func checkPass(email string, password string) bool {
	db := database.ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	passCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passCheck == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/", index)
	router.HandleFunc("/login", users.LoginUser).Methods("POST")
	router.HandleFunc("/register", users.RegisterUser).Methods("POST")
	router.HandleFunc("/transaction", transaction.CreateTransactionHandler).Methods("Post")
	router.HandleFunc("/updateEmail", users.UpdateUserEmail).Methods("PUT")
	router.HandleFunc("/updateBalance", users.UpdateUserEmail).Methods("PUT")

	fmt.Println("App is working on port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
