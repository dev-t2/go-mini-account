package main

import (
	"fmt"
	"learn-go-method/accounts"
)

func main() {
	var account = accounts.CreateAccount()

	fmt.Println("Balance:", account.Balance())

	account.Deposit(10)

	fmt.Println("Balance:", account.Balance())

	var err = account.Withdraw(20)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance:", account.Balance())
	}
}