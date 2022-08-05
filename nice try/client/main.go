package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("supersecretkey")

const PORT = ":9001"

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, validToken)
}

// you should recieve a json with username and password, then compare here

//get expected password by email, then

// if not ok return a json saying credentials are not ok

// if not then return jwtoken

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorization"] = true
	claims["email"] = "email"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong %s", err.Error())

		return "", err
	}
	return tokenString, nil
}
func main() {
	fmt.Println("server running on port " + PORT)
	handleRequests()
}
