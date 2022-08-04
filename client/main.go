package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("supersecretkey")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong %s", err.Error())

		return "", err
	}
	return tokenString, nil
}
func main() {
	fmt.Println("simple test")
	tokenString, err := GenerateJWT()
	if err != nil {

		fmt.Println("Error generating token string")
	}
	fmt.Println(tokenString)
}
