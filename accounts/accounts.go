package accounts

import (
	"errors"
	"fmt"
)

type account struct {
	balance int
}

var errWithdraw = errors.New("Withdraw Method Error")

func CreateAccount() *account {
	return &account{balance: 0}
}

func (a account) String() string{
	return fmt.Sprintf("Created Account Balance: %d", a.balance) 
}

func (a account) Balance() int {
	return a.balance
}

func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errWithdraw
	}

	a.balance -= amount

	return nil
}