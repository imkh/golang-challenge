package mfxrecruitdev

import (
	"fmt"
	"net/http"
)

// AccountsService handles communication with the accounts related methods
// of the MFX Recruit Dev API.
type AccountsService struct {
	client *Client
}

// Account represents an MFX Recruit Dev account.
type Account struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

// GetAccount gets a single account.
func (s *AccountsService) GetAccount(id int) (*Account, *http.Response, error) {
	path := fmt.Sprintf("accounts/%d", id)

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(Account)
	resp, err := s.client.Do(req, account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, err
}

// ListUserAccounts gets a list of accounts for the given user.
func (s *AccountsService) ListUserAccounts(userID int) ([]*Account, *http.Response, error) {
	path := fmt.Sprintf("users/%d/accounts", userID)

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var accounts []*Account
	resp, err := s.client.Do(req, &accounts)
	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, nil
}
