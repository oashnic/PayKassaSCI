# Paykassa SCI
[![GoDoc](https://godoc.org/github.com/oashnic/PayKassaSCI?status.svg)](https://godoc.org/github.com/oashnic/PayKassaSCI)
[![Go Report Card](https://goreportcard.com/badge/github.com/oashnic/PayKassaSCI)](https://goreportcard.com/report/github.com/oashnic/PayKassaSCI)

Golang wrapper for PayKassa SCI v0.4

## Install

Install the package with:

```bash
go get github.com/oashnic/PayKassaSCI
```

Import it with:

```go
import "github.com/oashnic/PayKassaSCI"
```

and use `PayKassaSCI` as the package name inside the code.

## Example

```go
package main

import (
	"fmt"
	"gitlab.com/oashnic/PayKassa"
	"log"
)

func main() {
	var merchantID = 1234
	var merchantKey = "QQQQQQQQQQQQQQQ"
	var testPay = false
	sci := PayKassaSCI.InitSCI(merchantID, merchantKey, testPay)

	var orderID = 1234
	var amount = 2
	var comment = "test"
	var phone = false
	var paidCommission = "shop"

	resp, err := sci.GetLinkForDeposit(orderID, amount, PayKassaSCI.BitCoin, comment, phone, paidCommission)

	if err != nil{
		log.Fatalln(err)
	}

	if resp.Error {
		log.Fatalln(resp.Message)
		return
	}

	fmt.Println("URL paying process: " + resp.Data.URL)
}
```