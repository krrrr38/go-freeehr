package freeehr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	companyTestDateTime1 = time.Date(2018, 3, 4, 6, 38, 11, 261000000, time.UTC)
	companyTestDateTime2 = time.Date(2018, 3, 4, 6, 38, 11, 262000000, time.UTC)
)

func TestCompanyService_GetEmployees(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/companies/999/employees", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":100,"num":"numnum","display_name":"displayname","entry_date": "2018-03-04T06:38:11.261+09:00","retire_date":"2018-03-04T06:38:11.262+09:00","user_id":200,"email":300}]`)
	})

	employees, _, err := client.Companies.GetEmployees(999, nil)
	if err != nil {
		t.Errorf("Companies.GetEmployees returned error: %v", err)
	}

	want := &[]Employee{
		Employee{
			ID:          100,
			Num:         "numnum",
			DisplayName: "displayname",
			EntryDate:   FreeeDateTime{companyTestDateTime1},
			RetireDate:  FreeeDateTime{companyTestDateTime2},
			UserID:      200,
			Email:       300,
		},
	}
	if !reflect.DeepEqual(employees, want) {
		t.Errorf("Companies.GetEmployees returned %+v, want %+v", employees, want)
	}
}

func TestCompanyService_GetEmployees_Paging(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/companies/999/employees", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("per") != "10" {
			t.Errorf("Companies.GetEmployees Paging `per` query doesnt matched: query=%v", r.URL.Query())
		}
		if r.URL.Query().Get("page") != "2" {
			t.Errorf("Companies.GetEmployees Paging `page` query doesnt matched: query=%v", r.URL.Query())
		}

		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":111,"num":"numnum","display_name":"displayname","entry_date": "2018-03-04T06:38:11.261+09:00","retire_date":"2018-03-04T06:38:11.262+09:00","user_id":222,"email":333}]`)
	})

	employees, _, err := client.Companies.GetEmployees(999, &PagingOption{Per: 10, Page: 2})
	if err != nil {
		t.Errorf("Companies.GetEmployees returned error: %v", err)
	}

	want := &[]Employee{
		Employee{
			ID:          111,
			Num:         "numnum",
			DisplayName: "displayname",
			EntryDate:   FreeeDateTime{companyTestDateTime1},
			RetireDate:  FreeeDateTime{companyTestDateTime2},
			UserID:      222,
			Email:       333,
		},
	}
	if !reflect.DeepEqual(employees, want) {
		t.Errorf("Companies.GetEmployees returned %+v, want %+v", employees, want)
	}
}
