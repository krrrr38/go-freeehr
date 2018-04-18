package freeehr

import (
	"fmt"
)

// CompanyService manage company related resources
type CompanyService service

// GetEmployees returns company employee list
// https://www.freee.co.jp/hr/api/#/%E5%BE%93%E6%A5%AD%E5%93%A1/index
func (s *CompanyService) GetEmployees(companyID int, pagingOption *PagingOption) (*[]Employee, *Response, error) {
	path := AddPagingQueryParam(fmt.Sprintf("/hr/api/v1/companies/%v/employees", companyID), pagingOption)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	employeesResponse := new([]Employee)
	resp, err := s.client.Do(req, employeesResponse)
	if err != nil {
		return nil, resp, err
	}

	return employeesResponse, resp, nil
}
