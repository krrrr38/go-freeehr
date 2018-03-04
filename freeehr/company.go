package freeehr

import (
	"context"
	"fmt"
)

// TODO
type CompanyService service

// TODO
type Company struct {
	Id         *int    `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Role       *string `json:"role,omitempty"`
	EmployeeId *int    `json:"employee_id,omitempty"`
}

// https://www.freee.co.jp/hr/api/#/%E5%BE%93%E6%A5%AD%E5%93%A1/index
func (s *CompanyService) GetEmployees(ctx context.Context, companyId int) (*[]Employee, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("/hr/api/v1/companies/%v/employees", companyId), nil)
	if err != nil {
		return nil, nil, err
	}

	employeesResponse := new([]Employee)
	resp, err := s.client.Do(ctx, req, employeesResponse)
	if err != nil {
		return nil, resp, err
	}

	return employeesResponse, resp, nil
}
