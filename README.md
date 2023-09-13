# go-mfx-recruit-dev

A Go library for accessing the [MFX Recruit Dev API](https://mfx-recruit-dev.herokuapp.com/golang).

## Installation

Inside your project directory, run:

```console
$ go get example.com/go-mfx-recruit-dev
```

or import the module and run `go get` without parameters.

```go
import "example.com/go-mfx-recruit-dev"
```

## Usage

```go
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
```

---

## 技術課題 / Coding Challenge

English follows Japanese

以下の3つのWeb APIがあります(実際にアクセスできます)。
これらを利用して、あるユーザーIDをあたえたときのユーザー名および登録金融機関一覧と、登録金融機関の残高(balance)を出力するプログラムを作成してください。

実装はGolangで行ってください。
言語の特性に基づき、保守性を考慮した設計・実装を行なってください。
また、解答の公開は避けてください。[secret gist](https://help.github.com/articles/about-gists/#secret-gists) あるいはメールでの提出をお願いします。
プログラムの作成にかかった時間も回答をお願いいたします。

---

There are three APIs below. (These are actual API, so you can access them.)
Please create a program with some of these APIs. It should take a user ID as input and return the name, account list and balances of that user. In addition, the program should be designed and implemented with maintainability in mind.

You can use Golang to do this programming.
Please send your result and time taken your answer by [secret gist](https://help.github.com/articles/about-gists/#secret-gists) or email and refrain from releasing it to public so that everyone has equal opportunities.

```json
- Web API
  - https://mfx-recruit-dev.herokuapp.com/users/1
    - レスポンス例 / Sample response
    {
        "id": 1,
        "name": "Alice",
        "account_ids": [
          1,
          3,
          5
        ]
    }
- https://mfx-recruit-dev.herokuapp.com/users/1/accounts
    - レスポンス例 / Sample response
    [
        {
            "id": 1,
            "user_id": 1,
            "name": "A銀行",
            "balance": 20000
        },
        {
            "id": 3,
            "user_id": 1,
            "name": "C信用金庫",
            "balance": 120000
        },
        {
            "id": 5,
            "user_id": 1,
            "name": "E銀行",
            "balance": 5000
        }
    ]
  - https://mfx-recruit-dev.herokuapp.com/accounts/2
    - レスポンス例 / Sample response
    {
        "id": 2,
        "user_id": 2,
        "name": "Bカード",
        "balance": 200
    }
```
