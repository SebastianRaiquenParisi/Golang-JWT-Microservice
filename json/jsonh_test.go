package json_handler

import (
	"testing"
)

func TestGetJsonFromPage(t *testing.T) {
	collection := GetJsonFromPage(1)
	if len(collection) == 0 {
		t.Errorf("Empty slice")
	}
	for _, user_dao := range collection {
		if (user_dao != User_dao{}) {
			t.Logf("Id: %d Email: %s First name: %s Last name: %s\n ", user_dao.Id, user_dao.Email, user_dao.First_name, user_dao.Last_name)
		} else {
			t.Errorf("Empty object")
		}
	}
}
