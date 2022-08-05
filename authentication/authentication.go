package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"http"
	"io/ioutil"
	"log"
	"time"

	"github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/models"
	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

// sign with private key
// verify with public key
func init() {

	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("Could not access the private file")
	}

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("Could not access public file")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Could not parse private key")
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Could not parse public key")
	}

}

func generateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Could not sign the token")
	}
	return result
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "Error reading the user %s", err)
	}
	if validateUser(user) {

	}
}

func validateUser(user models.User) bool {
	return false
}
