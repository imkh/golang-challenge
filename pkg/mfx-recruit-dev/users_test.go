package mfxrecruitdev

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	mux, client := setup(t)

	path := "/users"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_users.json")
	})

	users, _, err := client.Users.ListUsers()
	require.NoError(t, err)

	want := []*User{
		{
			ID:         1,
			Name:       "Alice",
			AccountIds: []int{1, 3, 5},
		},
		{
			ID:         2,
			Name:       "Bob",
			AccountIds: []int{2, 4},
		},
		{
			ID:         3,
			Name:       "Carol",
			AccountIds: []int{6},
		},
	}
	require.Equal(t, want, users)
}

func TestGetUser(t *testing.T) {
	mux, client := setup(t)

	path := "/users/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user.json")
	})

	user, _, err := client.Users.GetUser(1)
	require.NoError(t, err)

	want := &User{
		ID:         1,
		Name:       "Alice",
		AccountIds: []int{1, 3, 5},
	}
	require.Equal(t, want, user)
}
