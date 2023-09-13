package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	mfxrecruitdev "example.com/go-mfx-recruit-dev"
)

func main() {
	client, err := mfxrecruitdev.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create MFX Recruit Dev client: %v", err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:    "mfx-recruit-dev",
		Usage:   "A command-line tool for accessing the MFX Recruit Dev API",
		Version: "v1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "get a user and its accounts",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "id",
						Usage:    "ID of the user",
						Required: true,
					},
					&cli.BoolFlag{
						Name:  "json",
						Usage: "Print JSON output",
					},
				},
				Action: func(cCtx *cli.Context) error {
					userID := cCtx.Int("id")
					user, _, err := client.Users.GetUser(userID)
					if err != nil {
						return fmt.Errorf("failed to get user %d: %v", userID, err)
					}
					userAccounts, _, err := client.Accounts.ListUserAccounts(userID)
					if err != nil {
						return fmt.Errorf("failed to get user %d's accounts: %v", userID, err)
					}

					if cCtx.Bool("json") {
						printUserJSON(user, userAccounts)
					} else {
						printTable(user, userAccounts)
					}
					return nil
				},
			},
			{
				Name:    "account",
				Aliases: []string{"a"},
				Usage:   "get an account",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "id",
						Usage:    "ID of the account",
						Required: true,
					},
					&cli.BoolFlag{
						Name:  "json",
						Usage: "Print JSON output",
					},
				},
				Action: func(cCtx *cli.Context) error {
					accountID := cCtx.Int("id")
					account, _, err := client.Accounts.GetAccount(accountID)
					if err != nil {
						return fmt.Errorf("failed to get account %d: %v", accountID, err)
					}
					user, _, err := client.Users.GetUser(account.UserID)
					if err != nil {
						return fmt.Errorf("failed to get user %d: %v", account.UserID, err)
					}

					if cCtx.Bool("json") {
						printAccountJSON(account, user)
					} else {
						printTable(user, []*mfxrecruitdev.Account{account})
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
