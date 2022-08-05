package main

// the server should be the .NET
import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":9000"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Super Secret Information")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}
