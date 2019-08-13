# go-zaim

go-zaim is a Go client library for the [Zaim API](https://dev.zaim.net/home)

## Install

```sh
$ go get github.com/Sa2Knight/go-zaim
```

## Usage

### Authenticate

You can use the [Zaim API](https://dev.zaim.net/home) by passing credentials to the NewClient function.

```go
package main

import (
	"fmt"
	. "github.com/Sa2Knight/go-zaim"
	"os"
)

func main() {
	client := NewClient(
		os.Getenv("ZAIM_CUSTOMER_KEY"),
		os.Getenv("ZAIM_CUSTOMER_SECRET"),
		os.Getenv("ZAIM_TOKEN"),
		os.Getenv("ZAIM_SECRET"),
	)
	fmt.Println(client)
}
```

### REST API

You can use the following APIs through client.

- GET /v2/home/user/verify
- GET /v2/home/money
- POST /v2/home/money/payment
- POST /v2/home/money/income
- PUT /v2/home/money/payment:id
- PUT /v2/home/money/income:id
- DELETE v2/home/money/payment/:id
- DELETE v2/home/money/income/:id
- GET /v2/home/category
- GET /v2/home/genre
- GET /v2/home/account
- GET /v2/account
- GET /v2/category
- GET /v2/genre
- GET /v2/currency

Refer to the [documentation](https://dev.zaim.net/home/api) for details on the API.


For example, the following code will fetch a list of payments within a specified date.

```go
package main

import (
	"fmt"
	. "github.com/Sa2Knight/go-zaim"
	"os"
	"net/url"
)

func main() {
	client := NewClient(
		os.Getenv("ZAIM_CUSTOMER_KEY"),
		os.Getenv("ZAIM_CUSTOMER_SECRET"),
		os.Getenv("ZAIM_TOKEN"),
		os.Getenv("ZAIM_SECRET"),
	)
	moneySlice, _ := client.FetchMoney(url.Values{
		"start_date": []string{"2019-08-12"},
		"end_date": []string{"2019-08-14"},
	})
	fmt.Println(moneySlice)
}
```

# Update/Delete multiple records

When you execute the FetchMoney method, you get a MoneSlice. This has methods to update or delete multiple records at once.

For example, the following code will delete all today's payment records.

```go
package main

import (
	. "github.com/Sa2Knight/go-zaim"
	"net/url"
	"os"
	"time"
)

func main() {
	today := time.Now().Format("2006-01-02")
	client := NewClient(
		os.Getenv("ZAIM_CUSTOMER_KEY"),
		os.Getenv("ZAIM_CUSTOMER_SECRET"),
		os.Getenv("ZAIM_TOKEN"),
		os.Getenv("ZAIM_SECRET"),
	)
	moneySlice, _ := client.FetchMoney(url.Values{
		"start_date": []string{today},
		"end_date":   []string{today},
	})
	moneySlice.DeleteAll(client)
}
```
