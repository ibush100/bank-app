package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateUser(username string, email string, password string) (interfaces.User, bool) {
	user, result := database.CreateUser(username, email, password)
	// need to clean up returning true
	return user, result
}

func PrepareToken() string {
	tokenContent := jwt.MapClaims{
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

//Need to figure this out - still working it out
func VerifyToken(jwtToken string) bool {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	helpers.HandleErr(err)
	//	var userId, _ = strconv.ParseFloat(id, 8)
	if !token.Valid {
		return false
	}
	return true
}
