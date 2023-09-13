package mfxrecruitdev

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	mux, client := setup(t)

	path := "/accounts/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_account.json")
	})

	account, _, err := client.Accounts.GetAccount(2)
	require.NoError(t, err)

	want := &Account{
		ID:      2,
		UserID:  2,
		Name:    "Bカード",
		Balance: 200,
	}
	require.Equal(t, want, account)
}

func TestListUserAccounts(t *testing.T) {
	mux, client := setup(t)

	path := "/users/1/accounts"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_user_accounts.json")
	})

	accounts, _, err := client.Accounts.ListUserAccounts(1)
	require.NoError(t, err)

	want := []*Account{
		{
			ID:      1,
			UserID:  1,
			Name:    "A銀行",
			Balance: 20000,
		},
		{
			ID:      3,
			UserID:  1,
			Name:    "C信用金庫",
			Balance: 120000,
		},
		{
			ID:      5,
			UserID:  1,
			Name:    "E銀行",
			Balance: 5000,
		},
	}
	require.Equal(t, want, accounts)
}
