package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type ErrResponse struct {
	Message string
}

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ReadBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	HandleErr(err)

	return body
}

func HashAndSalt(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func WriteToJson(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		HandleErr(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()
			if error != nil {
				log.Println(error)

				resp := ErrResponse{Message: "Internal server error"}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=bankapp sslmode=disable")
	HandleErr(err)
	return db
}
