# go-mfx-recruit-dev

A Go library for accessing the [MFX Recruit Dev API](https://mfx-recruit-dev.herokuapp.com/golang).

This library is strongly inspired by [go-gitlab](https://github.com/xanzy/go-gitlab) (which itself is inspired by [go-github](https://github.com/google/go-github)). I worked a lot with go-gitlab at my previous job, which made me quite familiar with its inner workings. I [contributed small changes](https://github.com/xanzy/go-gitlab/commits/master?author=imkh) a few times to the upstream repository whenever I needed it for my work, and I built a number of libraries in the same style both for personal use ([here](https://github.com/imkh/go-senscritique), [here](https://github.com/imkh/go-vgmdb) and [here](https://github.com/imkh/go-itunes-search)) and professional use (though the repos are private, as you can imagine).

## Table of Contents

- [Time to complete](#time-to-complete)
- [Maintainability](#maintainability)
- [Library](#library)
    - [Installation](#installation)
    - [Usage](#usage)
- [CLI application](#cli-application)
	- [Build](#build)
	- [Run](#run)

## Time to complete

* `mfxrecruitdev` library package in root + tests: 2~3 hours
* CLI app in `cmd/mfx-recruit-dev`: ~4 hours

## Maintainability

The library package contains code that is currently unused but allows for simple and straightforward support for eventual new GET, POST and PUT endpoints that need to send URL query parameters or a JSON body (`NewRequest` and `Do` functions in `mfx-recruit-dev.go`).

For the CLI application in `cmd/mfx-recruit-dev`, the standard library `flag` would be enough for the current use case, but I opted to use [urfave/cli](https://github.com/urfave/cli) as its declarative nature makes it extremely easy to read, understand, and extend the code to add new commands.

## Library

### Installation

Inside your project directory, run:

```console
$ go get example.com/go-mfx-recruit-dev
```

or import the module and run `go get` without parameters.

```go
import "example.com/go-mfx-recruit-dev"
```

### Usage

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
		log.Fatalf("Failed to create MFX Recruit Dev client: %v", err)
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
	_, _, err = client.Users.GetUser(15)
	if err != nil {
		fmt.Println(err)
	}
}
```

## CLI application

This `cmd/mfx-recruit-dev` directory contains source code for a simple CLI application that makes use of the library.

### Build

```console
$ go build -o mfx-recruit-dev cmd/mfx-recruit-dev/*.go
```

### Run

```console
$ ./mfx-recruit-dev
NAME:
   mfx-recruit-dev - A command-line tool for accessing the MFX Recruit Dev API

USAGE:
   mfx-recruit-dev [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   user, u     get a user and its accounts
   account, a  get an account
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

$ ./mfx-recruit-dev user --id 1
User ID: 1
User name: Alice
Number of accounts: 3
+----+--------------+---------+
| ID | ACCOUNT NAME | BALANCE |
+----+--------------+---------+
|  1 | A銀行         |  ¥20000 |
|  3 | C信用金庫      | ¥120000 |
|  5 | E銀行         |   ¥5000 |
+----+--------------+---------+
|         TOTAL     | ¥145000 |
+----+--------------+---------+

$ ./mfx-recruit-dev user --id 1 --json
{
        "id": 1,
        "name": "Alice",
        "accounts": [
                {
                        "id": 1,
                        "name": "A銀行",
                        "balance": 20000
                },
                {
                        "id": 3,
                        "name": "C信用金庫",
                        "balance": 120000
                },
                {
                        "id": 5,
                        "name": "E銀行",
                        "balance": 5000
                }
        ]
}

$ ./mfx-recruit-dev account --id 2
User ID: 2
User name: Bob
Number of accounts: 1
+----+--------------+---------+
| ID | ACCOUNT NAME | BALANCE |
+----+--------------+---------+
|  2 | Bカード       |    ¥200 |
+----+--------------+---------+

$ ./mfx-recruit-dev account --id 2 --json
{
        "id": 2,
        "user": {
                "id": 2,
                "name": "Bob"
        },
        "name": "Bカード",
        "balance": 200
}
```

---

### 技術課題 / Coding Challenge

English follows Japanese

以下の3つのWeb APIがあります(実際にアクセスできます)。
これらを利用して、あるユーザーIDをあたえたときのユーザー名および登録金融機関一覧と、登録金融機関の残高(balance)を出力するプログラムを作成してください。

実装はGolangで行ってください。
言語の特性に基づき、保守性を考慮した設計・実装を行なってください。
また、解答の公開は避けてください。[secret gist](https://help.github.com/articles/about-gists/#secret-gists) あるいはメールでの提出をお願いします。
プログラムの作成にかかった時間も回答をお願いいたします。


There are three APIs below. (These are actual API, so you can access them.)
Please create a program with some of these APIs. It should take a user ID as input and return the name, account list and balances of that user. In addition, the program should be designed and implemented with maintainability in mind.

You can use Golang to do this programming.
Please send your result and time taken your answer by [secret gist](https://help.github.com/articles/about-gists/#secret-gists) or email and refrain from releasing it to public so that everyone has equal opportunities.

```
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
