package json_handler

import (
	"encoding/json"
	"net/http"
	"time"
)

var client *http.Client

func GetUsersFromPage() {

}

type Response struct {
	Data []User_dao `json:"data"`
}

type User_dao struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Role       string `json:"role, omitempty"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
}
