package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/db"
	"github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/models"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

// sign with private key
// verify with public key
func init() {
	db.CreateAndWriteWithConstants()
	privateBytes, err := ioutil.ReadFile("./should_be_outside_respository/private.rsa")
	LogFatalIfError(err, "Could not access the private file")

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	LogFatalIfError(err, "Could not access public file")

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	LogFatalIfError(err, "Could not parse private key")

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	LogFatalIfError(err, "Could not parse public key")

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
	LogFatalIfError(err, "Could not sign the token")

	return result
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error reading the user %s", err)
	}
	if validateUser(user) {
		// set password to empty
		// so it wont show up on the JTW
		user.Password = ""
		user.Role = getRoleFromDB(user)

		token := generateJWT(user)
		result := models.ResponseToken{token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintln(w, "Error generating JWT JSON")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Email or password not valid")
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Fprintln(w, "The token has expired")
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Fprintln(w, "The token has been tampered")
				return
			default:
				fmt.Fprintln(w, "The token is not valid")
				return
			}
		default:
			fmt.Fprintln(w, "The token is not valid")
			break
		}
	}
	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Welcome! The token is valid for 60 minutes")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "The token is not valid")
	}
}

// replace func with searching in json/database
// and adjusting the functions accordingly
//
func getRoleFromDB(user models.User) string {
	return "admin"
}

func LogFatalIfError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func validateUser(user models.User) bool {
	return user.Email == "email" && user.Password == "pswd"
}
