package main

import (
	"log"
	"net/http"

	"github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/authentication"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", authentication.Login)
	mux.HandleFunc("/validate", authentication.ValidateToken)

	log.Println("Listening api in http://localhost" + PORT)
	http.ListenAndServe(PORT, mux)
}
