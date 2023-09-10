package mfxrecruitdev

import (
	"fmt"
	"net/http"
)

// UsersService handles communication with the users related methods
// of the MFX Recruit Dev API.
type UsersService struct {
	client *Client
}

// User represents an MFX Recruit Dev user.
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AccountIds []int  `json:"account_ids"`
}

// ListUsers gets a list of users.
func (s *UsersService) ListUsers() ([]*User, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "users", nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// GetUser gets a single user.
func (s *UsersService) GetUser(id int) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%d", id)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}
