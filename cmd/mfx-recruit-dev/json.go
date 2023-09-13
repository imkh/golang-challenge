package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	mfxrecruitdev "example.com/go-mfx-recruit-dev"
)

// The code in this file merges the User and Account structs to output a single JSON object for external use (e.g., to pass into jq).

type User struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Accounts []*Account `json:"accounts,omitempty"`
}

type Account struct {
	ID      int    `json:"id"`
	User    *User  `json:"user,omitempty"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func printUserJSON(user *mfxrecruitdev.User, accounts []*mfxrecruitdev.Account) {
	// Merge Accounts into User
	output := &User{
		ID:   user.ID,
		Name: user.Name,
	}
	for _, account := range accounts {
		output.Accounts = append(output.Accounts, &Account{
			ID:      account.ID,
			Name:    account.Name,
			Balance: account.Balance,
		})
	}

	prettyPrint(output)
}

func printAccountJSON(account *mfxrecruitdev.Account, user *mfxrecruitdev.User) {
	// Merge User into Account
	output := &Account{
		ID: account.ID,
		User: &User{
			ID:   user.ID,
			Name: user.Name,
		},
		Name:    account.Name,
		Balance: account.Balance,
	}

	prettyPrint(output)
}

// Pretty print JSON
func prettyPrint(i interface{}) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	encoder.Encode(i)
	fmt.Println(buffer)
}
