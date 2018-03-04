package freeehr

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCompanyService_GetEmployees(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/companies/999/employees", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":100,"num":"numnum","display_name":"displayname","entry_date": "2018-03-04T06:38:11.261Z","retire_date":"2018-03-04T06:38:11.262Z","user_id":200,"email":300}]`)
	})

	employees, _, err := client.Companies.GetEmployees(context.Background(), 999, nil)
	if err != nil {
		t.Errorf("Companies.GetEmployees returned error: %v", err)
	}

	want := &[]Employee{
		Employee{
			ID:          100,
			Num:         "numnum",
			DisplayName: "displayname",
			EntryDate:   "2018-03-04T06:38:11.261Z",
			RetireDate:  "2018-03-04T06:38:11.262Z",
			UserID:      200,
			Email:       300,
		},
	}
	if !reflect.DeepEqual(employees, want) {
		t.Errorf("Companies.GetEmployees returned %+v, want %+v", employees, want)
	}
}
