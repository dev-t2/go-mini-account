package main

import (
	"fmt"
	"learn-go-method/accounts"
)

func main() {
	var account = accounts.CreateAccount(0)

	fmt.Println(account)
}