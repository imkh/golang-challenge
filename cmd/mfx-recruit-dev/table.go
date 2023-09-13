package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	mfxrecruitdev "example.com/go-mfx-recruit-dev"
)

const (
	LOW_BALANCE = 10000
)

func printTable(user *mfxrecruitdev.User, accounts []*mfxrecruitdev.Account) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Account Name", "Balance"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})

	// Used for printing localized human-readable account balances (¥120000 -> ¥120,000)
	printer := message.NewPrinter(language.English)

	var totalBalance int
	for _, account := range accounts {
		rows := []string{
			strconv.Itoa(account.ID),
			account.Name,
			printer.Sprintf("¥%d", account.Balance),
		}

		var color tablewriter.Colors
		if account.Balance < LOW_BALANCE {
			color = append(color, tablewriter.FgRedColor)
		}
		colors := []tablewriter.Colors{}
		for range rows {
			colors = append(colors, color)
		}

		table.Rich(rows, colors)
		totalBalance += account.Balance
	}
	if len(accounts) > 1 {
		table.SetFooter([]string{
			"",
			"TOTAL",
			printer.Sprintf("¥%d", totalBalance),
		})
		if totalBalance < LOW_BALANCE {
			table.SetFooterColor(
				tablewriter.Colors{tablewriter.FgRedColor},
				tablewriter.Colors{tablewriter.FgRedColor},
				tablewriter.Colors{tablewriter.FgRedColor},
			)
		}
	}

	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("User name: %s\n", user.Name)
	fmt.Printf("Number of accounts: %d\n", len(accounts))
	table.Render()
}
