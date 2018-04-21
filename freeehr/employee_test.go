package freeehr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	employeeTestDate0304 = time.Date(2018, 3, 4, 0, 0, 0, 0, time.UTC)
	employeeTestDate0305 = time.Date(2018, 3, 5, 0, 0, 0, 0, time.UTC)
	employeeTestDate0306 = time.Date(2018, 3, 6, 0, 0, 0, 0, time.UTC)
	employeeTestDateTime = time.Date(2018, 3, 4, 6, 31, 9, 895000000, time.UTC)
)

func TestEmployeeService_GetWorkRecord(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hr/api/v1/employees/999/work_records/2018-03-04", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		// {"break_records":[{"clock_in_at":"2018-04-18T12:00:00.000+09:00","clock_out_at":"2018-04-18T13:00:00.000+09:00"}],"clock_in_at":"2018-04-18T09:50:00.000+09:00","clock_out_at":"2018-04-18T21:37:00.000+09:00","date":"2018-04-18","day_pattern":"normal_day","early_leaving_mins":0,"is_absence":false,"is_editable":true,"lateness_mins":0,"normal_work_clock_in_at":"2018-04-18T08:45:00.000+09:00","normal_work_clock_out_at":"2018-04-18T17:45:00.000+09:00","normal_work_mins":480,"normal_work_mins_by_paid_holiday":0,"note":"","paid_holiday":0.0,"use_attendance_deduction":false,"use_default_work_pattern":true}
		fmt.Fprint(w, `{"break_records":[{"clock_in_at":"2018-03-04T06:31:09.895+09:00","clock_out_at":null}],"clock_in_at":"2018-03-04T06:31:09.895+09:00","clock_out_at":"2018-03-04T06:31:09.895+09:00","date":"2018-03-04","day_pattern":"string","early_leaving_mins":0,"is_absence":true,"is_editable":true,"lateness_mins":0,"normal_work_clock_in_at":"2018-03-04T06:31:09.895+09:00","normal_work_clock_out_at":"2018-03-04T06:31:09.895+09:00","normal_work_mins":0,"normal_work_mins_by_paid_holiday":0,"note":"string","paid_holiday":0.0,"use_attendance_deduction":true,"use_default_work_pattern":true}`)
	})

	workRecord, _, err := client.Employees.GetWorkRecord(999, FreeeDate{employeeTestDate0304})
	if err != nil {
		t.Errorf("Employees.GetWorkRecord returned error: %v", err)
	}

	want := &WorkRecord{
		BreakRecords: []BreakRecord{
			BreakRecord{
				ClockInAt:  &FreeeDateTime{employeeTestDateTime},
				ClockOutAt: nil,
			},
		},
		ClockInAt:                   &FreeeDateTime{employeeTestDateTime},
		ClockOutAt:                  &FreeeDateTime{employeeTestDateTime},
		Date:                        &FreeeDate{employeeTestDate0304},
		DayPattern:                  "string",
		EarlyLeavingMins:            0,
		IsAbsence:                   true,
		IsEditable:                  true,
		LatenessMins:                0,
		NormalWorkClockInAt:         &FreeeDateTime{employeeTestDateTime},
		NormalWorkClockOutAt:        &FreeeDateTime{employeeTestDateTime},
		NormalWorkMins:              0,
		NormalWorkMinsByPaidHoliday: 0,
		Note:                   "string",
		PaidHoliday:            0.0,
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
				ClockInAt:  &FreeeDateTime{employeeTestDateTime},
				ClockOutAt: nil,
			},
		},
		ClockInAt:                   &FreeeDateTime{employeeTestDateTime},
		ClockOutAt:                  &FreeeDateTime{employeeTestDateTime},
		Date:                        &FreeeDate{employeeTestDate0304},
		DayPattern:                  "string",
		EarlyLeavingMins:            0,
		IsAbsence:                   true,
		IsEditable:                  true,
		LatenessMins:                0,
		NormalWorkClockInAt:         &FreeeDateTime{employeeTestDateTime},
		NormalWorkClockOutAt:        &FreeeDateTime{employeeTestDateTime},
		NormalWorkMins:              0,
		NormalWorkMinsByPaidHoliday: 0,
		Note:                   "string",
		PaidHoliday:            0.0,
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
		fmt.Fprint(w, `{"break_records":[{"clock_in_at":"2018-03-04T06:31:09.895+09:00"}]}`)
	})

	workRecord, _, err := client.Employees.PutWorkRecord(999, FreeeDate{employeeTestDate0305}, input)
	if err != nil {
		t.Errorf("Employees.PutWorkRecord returned error: %v", err)
	}

	want := &WorkRecord{
		BreakRecords: []BreakRecord{
			BreakRecord{
				ClockInAt: &FreeeDateTime{employeeTestDateTime},
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

	err := client.Employees.DeleteWorkRecord(999, FreeeDate{employeeTestDate0306})
	if err != nil {
		t.Errorf("Employees.PutWorkRecord returned error: %v", err)
	}
}
