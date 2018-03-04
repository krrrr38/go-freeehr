package freeehr

import (
	"context"
	"fmt"
)

// EmployeeService manage employee related resources
type EmployeeService service

// Employee is company employee
type Employee struct {
	ID          int           `json:"id,omitempty"`
	Num         string        `json:"num,omitempty"`
	DisplayName string        `json:"display_name,omitempty"`
	EntryDate   FreeeDateTime `json:"entry_date,omitempty"`
	RetireDate  FreeeDateTime `json:"retire_date,omitempty"`
	UserID      int           `json:"user_id,omitempty"`
	Email       int           `json:"email,omitempty"`
}

// BreakRecord represents break time
type BreakRecord struct {
	ClockInAt  FreeeDateTime `json:"clock_in_at,omitempty"`
	ClockOutAt FreeeDateTime `json:"clock_out_at,omitempty"`
}

// WorkRecord represents work record by date
type WorkRecord struct {
	BreakRecords                []BreakRecord `json:"break_records,omitempty"`
	ClockInAt                   FreeeDateTime `json:"clock_in_at,omitempty"`
	ClockOutAt                  FreeeDateTime `json:"clock_out_at,omitempty"`
	Date                        FreeeDateTime `json:"date,omitempty"`
	DayPattern                  string        `json:"day_pattern,omitempty"`
	EarlyLeavingMins            int           `json:"early_leaving_mins,omitempty"`
	IsAbsence                   bool          `json:"is_absence,omitempty"`
	IsEditable                  bool          `json:"is_editable,omitempty"`
	LatenessMins                int           `json:"lateness_mins,omitempty"`
	NormalWorkClockInAt         FreeeDateTime `json:"normal_work_clock_in_at,omitempty"`
	NormalWorkClockOutAt        FreeeDateTime `json:"normal_work_clock_out_at,omitempty"`
	NormalWorkMins              int           `json:"normal_work_mins,omitempty"`
	NormalWorkMinsByPaidHoliday int           `json:"normal_work_mins_by_paid_holiday,omitempty"`
	Note                        string        `json:"note,omitempty"`
	PaidHoliday                 int           `json:"paid_holiday,omitempty"`
	UseAttendanceDeduction      bool          `json:"use_attendance_deduction,omitempty"`
	UseDefaultWorkPattern       bool          `json:"use_default_work_pattern,omitempty"`
}

// GetWorkRecord can fetch target date work record
// https://www.freee.co.jp/hr/api/#/%E5%8B%A4%E6%80%A0/show2
func (s *EmployeeService) GetWorkRecord(ctx context.Context, employeeID int, date FreeeDate) (*WorkRecord, *Response, error) {
	path := fmt.Sprintf("/hr/api/v1/employees/%d/work_records/%s", employeeID, date)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	workRecordResponse := new(WorkRecord)
	resp, err := s.client.Do(ctx, req, workRecordResponse)
	if err != nil {
		return nil, resp, err
	}

	return workRecordResponse, resp, nil
}

// PutWorkRecord can register/overwrite target date work record
// https://www.freee.co.jp/hr/api/#/%E5%8B%A4%E6%80%A0/update
func (s *EmployeeService) PutWorkRecord(ctx context.Context, employeeID int, date FreeeDate, workRecord *WorkRecord) (*WorkRecord, *Response, error) {
	path := fmt.Sprintf("/hr/api/v1/employees/%d/work_records/%s", employeeID, date)
	req, err := s.client.NewRequest("PUT", path, workRecord)
	if err != nil {
		return nil, nil, err
	}

	workRecordResponse := new(WorkRecord)
	resp, err := s.client.Do(ctx, req, workRecordResponse)
	if err != nil {
		return nil, resp, err
	}

	return workRecordResponse, resp, nil
}

// DeleteWorkRecord https://www.freee.co.jp/hr/api/#/%E5%8B%A4%E6%80%A0/destroy
func (s *EmployeeService) DeleteWorkRecord(ctx context.Context, employeeID int, date FreeeDate) error {
	path := fmt.Sprintf("/hr/api/v1/employees/%d/work_records/%s", employeeID, date)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
