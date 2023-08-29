package main

import (
	"fmt"
	"learn-go-method/accounts"
	"learn-go-method/dictionary"
)

func main() {
	var account = accounts.CreateAccount()

	fmt.Println(account)

	fmt.Println("Balance:", account.Balance())

	account.Deposit(10)

	fmt.Println("Balance:", account.Balance())

	var err = account.Withdraw(20)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance:", account.Balance())
	}

	fmt.Println()

	var dict = dictionary.CreateDictionary("Hello", "World")

	value, err := dict.Search("Helloo")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}