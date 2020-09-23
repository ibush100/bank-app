package api

import (
	"bank-app/helpers"
	"bank-app/interfaces"
	"bank-app/transaction"
	"bank-app/users"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

// TODO
// swtich marshal to encoder

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var fomattedUser interfaces.Register
	err := json.Unmarshal(body, &fomattedUser)
	helpers.HandleErr(err)
	registerUser, result := users.CreateUser(fomattedUser.Username, fomattedUser.Email, fomattedUser.Password)
	if result {
		w.WriteHeader(http.StatusCreated)
		helpers.WriteToJson(w, registerUser)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func updateUserEmail(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody interfaces.UpdateUserEmail
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if isUserPresent(formattedBody.Email) {
		updateEmail(formattedBody.NewEmail, formattedBody.Email)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func updateEmail(newEmail string, email string) {
	db := helpers.ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	user.Email = newEmail
	db.Save(user)

}

func updateUserBalance(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody interfaces.UpdateUserBalance
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if isUserPresent(formattedBody.Email) {
		transaction.TopUpBalance(formattedBody.Email, formattedBody.TopUp)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody interfaces.Transaction
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if isUserPresent(formattedBody.PayorEmail) && isUserPresent(formattedBody.PayeeEmail) {
		transaction.CreateTransaction(formattedBody.PayeeEmail, formattedBody.PayorEmail, formattedBody.Amount)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Call successful dude")
}

// this will break until you put bcryt stuff in
func loginUser(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody interfaces.User
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	if checkPass(formattedBody.Email, formattedBody.Password) {
		token := PrepareToken(formattedBody.ID)
		w.WriteHeader(http.StatusCreated)
		// create token
		helpers.WriteToJson(w, token)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func isUserPresent(email string) bool {
	userResult := FindUser(email)
	if userResult <= 0 {
		return false
	}
	return true
}

func FindUser(email string) uint {
	db := helpers.ConnectDB()
	var user interfaces.User
	//db.Table("users").Select("user_id").Where("email = ? ", email).First(&user.ID)
	db.Where("email = ?", email).First(&user)

	return user.ID
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
	db := helpers.ConnectDB()
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
	router.HandleFunc("/login", loginUser).Methods("POST")
	router.HandleFunc("/register", registerUser).Methods("POST")
	router.HandleFunc("/updateEmail", updateUserEmail).Methods("PUT")
	router.HandleFunc("/updateBalance", updateUserEmail).Methods("PUT")

	fmt.Println("App is working on port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
