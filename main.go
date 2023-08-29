package main

import (
	"fmt"
	"learn-go-method/accounts"
	"learn-go-method/dictionary"
)

func main() {
	fmt.Println("1. Account")

	var account = accounts.CreateAccount()

	fmt.Println(account)

	fmt.Println("Balance:", account.Balance())

	account.Deposit(10)

	fmt.Println("Balance:", account.Balance())

	var err = account.Withdraw(20)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance:", account.Balance())
	}

	fmt.Println()
	fmt.Println("2. Dictionary")

	var dict = dictionary.CreateDictionary("Hello", "World")

	value, err := dict.Search("Hello")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Key: Hello, Value:", value)
	}

	err = dict.Add("Go", "Programming")

	if err != nil {
		fmt.Println(err)
	}

	value, err = dict.Search("Go")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Key: Go, Value:", value)
	}

	err = dict.Update("Go", "Awesome")

	if err != nil {
		fmt.Println(err)
	}

	value, err = dict.Search("Go")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Key: Go, Updated Value:", value)
	}

	err = dict.Delete("Go")

	if err != nil {
		fmt.Println(err)
	}

	value, err = dict.Search("Go")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Key: Go, Value:", value)
	}
}