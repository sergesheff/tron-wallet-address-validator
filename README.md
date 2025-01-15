Application for Tron wallets address validation
===

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