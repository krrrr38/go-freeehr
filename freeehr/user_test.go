package freeehr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserService_GetMe(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/users/me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":100,"companies": [{"id":200,"name":"namename","role": "rolerole","employee_id":300}]}`)
	})

	user, _, err := client.Users.GetMe()
	if err != nil {
		t.Errorf("Users.GetMe returned error: %v", err)
	}

	want := &User{
		ID: 100,
		Companies: []UserCompany{
			UserCompany{
				ID:         200,
				Name:       "namename",
				Role:       "rolerole",
				EmployeeID: 300,
			},
		},
	}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.GetMe returned %+v, want %+v", user, want)
	}
}
