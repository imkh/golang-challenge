package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	mfxrecruitdev "example.com/go-mfx-recruit-dev"
)

func prettyPrint(i interface{}) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	encoder.Encode(i)
	fmt.Println(buffer)
}

func main() {
	client, err := mfxrecruitdev.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// List users
	fmt.Println("### List users")
	users, _, err := client.Users.ListUsers()
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(users)

	// Get user
	fmt.Println("### Get user 1")
	user, _, err := client.Users.GetUser(1)
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(user)

	// List user 1's accounts
	fmt.Println("### List user 1's accounts")
	userAccounts, _, err := client.Accounts.ListUserAccounts(1)
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(userAccounts)

	// Get account
	fmt.Println("### Get account 2")
	account, _, err := client.Accounts.GetAccount(2)
	if err != nil {
		log.Fatal(err)
	}
	prettyPrint(account)

	// User Not Found
	fmt.Println("### User Not Found")
	_, _, err = client.Users.GetUser(10)
	if err != nil {
		fmt.Println(err)
	}
}
