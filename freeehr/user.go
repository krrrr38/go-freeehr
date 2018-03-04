package freeehr

import (
	"context"
)

// TODO
type UserService service

// TODO
type User struct {
	Id        *int          `json:"id,omitempty"`
	Companies []UserCompany `json:"companies,omitempty"`
}

// TODO
type UserCompany struct {
	Id         *int    `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Role       *string `json:"role,omitempty"`
	EmployeeId *int    `json:"employee_id,omitempty"`
}

// https://www.freee.co.jp/hr/api/#/%E3%83%AD%E3%82%B0%E3%82%A4%E3%83%B3%E3%83%A6%E3%83%BC%E3%82%B6/show1
func (s *UserService) GetMe(ctx context.Context) (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "/hr/api/v1/users/me", nil)
	if err != nil {
		return nil, nil, err
	}

	userResponse := new(User)
	resp, err := s.client.Do(ctx, req, userResponse)
	if err != nil {
		return nil, resp, err
	}

	return userResponse, resp, nil
}
