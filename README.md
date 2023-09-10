# money-forward-challenge

Website: https://mfx-recruit-dev.herokuapp.com/golang

### 技術課題 / Coding Challenge

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
