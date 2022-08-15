package json_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var client *http.Client

func GetUsersFromPage() {

}

const userApiPage = "https://reqres.in/api/users?page="

type Response struct {
	Data []User_dao `json:"data"`
}

type User_dao struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Role       string `json:"role"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

func GetJsonFromPage(page int) []User_dao {
	// url + page
	url := userApiPage + strconv.Itoa(page)
	// users data access object slice
	users := []User_dao{}
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return users
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
func mainP() {
	client = &http.Client{Timeout: 10 * time.Second}
}
