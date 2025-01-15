Application for Tron wallets address validation
===

[![Go Reference](https://pkg.go.dev/badge/github.com/sergesheff/tron-wallet-address-validator.svg)](https://pkg.go.dev/github.com/sergesheff/tron-wallet-address-validator)


This is a simple application for validating Tron wallets addresses with using of golang.

Sample of usage:
```golang
package main

import (
	"fmt"
	"github.com/sergesheff/tron-wallet-address-validator"
)

func main() {
	ok := tron.IsValidAddress("TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE")
	fmt.Println(ok)
}

```