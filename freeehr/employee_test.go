package freeehr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEmployeeService_GetWorkRecord(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/employees/999/work_records/2018-03-04", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"break_records":[{"clock_in_at":"2018-03-04T06:31:09.895Z","clock_out_at":"2018-03-04T06:31:09.895Z"}],"clock_in_at":"2018-03-04T06:31:09.896Z","clock_out_at":"2018-03-04T06:31:09.896Z","date":"2018-03-04T06:31:09.896Z","day_pattern":"string","early_leaving_mins":0,"is_absence":true,"is_editable":true,"lateness_mins":0,"normal_work_clock_in_at":"2018-03-04T06:31:09.896Z","normal_work_clock_out_at":"2018-03-04T06:31:09.896Z","normal_work_mins":0,"normal_work_mins_by_paid_holiday":0,"note":"string","paid_holiday":0,"use_attendance_deduction":true,"use_default_work_pattern":true}`)
	})

	workRecord, _, err := client.Employees.GetWorkRecord(context.Background(), 999, "2018-03-04")
	if err != nil {
		t.Errorf("Employees.GetWorkRecord returned error: %v", err)
	}

	want := &WorkRecord{
		BreakRecords: []BreakRecord{
			BreakRecord{
				ClockInAt:  "2018-03-04T06:31:09.895Z",
				ClockOutAt: "2018-03-04T06:31:09.895Z",
			},
		},
		ClockInAt:                   "2018-03-04T06:31:09.896Z",
		ClockOutAt:                  "2018-03-04T06:31:09.896Z",
		Date:                        "2018-03-04T06:31:09.896Z",
		DayPattern:                  "string",
		EarlyLeavingMins:            0,
		IsAbsence:                   true,
		IsEditable:                  true,
		LatenessMins:                0,
		NormalWorkClockInAt:         "2018-03-04T06:31:09.896Z",
		NormalWorkClockOutAt:        "2018-03-04T06:31:09.896Z",
		NormalWorkMins:              0,
		NormalWorkMinsByPaidHoliday: 0,
		Note:                   "string",
		PaidHoliday:            0,
		UseAttendanceDeduction: true,
		UseDefaultWorkPattern:  true,
	}
	if !reflect.DeepEqual(workRecord, want) {
		t.Errorf("Employees.GetWorkRecord returned %+v, want %+v", workRecord, want)
	}
}

func TestEmployeeService_PutWorkRecord(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &WorkRecord{
		BreakRecords: []BreakRecord{
			BreakRecord{
				ClockInAt:  "2018-03-04T06:31:09.895Z",
				ClockOutAt: "2018-03-04T06:31:09.895Z",
			},
		},
		ClockInAt:                   "2018-03-04T06:31:09.896Z",
		ClockOutAt:                  "2018-03-04T06:31:09.896Z",
		Date:                        "2018-03-04T06:31:09.896Z",
		DayPattern:                  "string",
		EarlyLeavingMins:            0,
		IsAbsence:                   true,
		IsEditable:                  true,
		LatenessMins:                0,
		NormalWorkClockInAt:         "2018-03-04T06:31:09.896Z",
		NormalWorkClockOutAt:        "2018-03-04T06:31:09.896Z",
		NormalWorkMins:              0,
		NormalWorkMinsByPaidHoliday: 0,
		Note:                   "string",
		PaidHoliday:            0,
		UseAttendanceDeduction: true,
		UseDefaultWorkPattern:  true,
	}

	mux.HandleFunc("/hr/api/v1/employees/999/work_records/2018-03-05", func(w http.ResponseWriter, r *http.Request) {
		v := new(WorkRecord)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"break_records":[{"clock_in_at":"2040-03-04T06:31:09.895Z"}]}`)
	})

	workRecord, _, err := client.Employees.PutWorkRecord(context.Background(), 999, "2018-03-05", input)
	if err != nil {
		t.Errorf("Employees.PutWorkRecord returned error: %v", err)
	}

	want := &WorkRecord{
		BreakRecords: []BreakRecord{
			BreakRecord{
				ClockInAt: "2040-03-04T06:31:09.895Z",
			},
		},
	}

	if !reflect.DeepEqual(workRecord, want) {
		t.Errorf("Employees.PutWorkRecord returned %+v, want %+v", workRecord, want)
	}
}

func TestEmployeeService_DeleteWorkRecord(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/employees/999/work_records/2018-03-06", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{}`)
	})

	err := client.Employees.DeleteWorkRecord(context.Background(), 999, "2018-03-06")
	if err != nil {
		t.Errorf("Employees.PutWorkRecord returned error: %v", err)
	}
}
