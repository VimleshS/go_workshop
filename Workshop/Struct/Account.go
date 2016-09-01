package main

import (
	"fmt"
)

type Account struct {
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) WithDraw(amount float64) {
	a.Balance -= amount
}

func main() {
	acc := Account{"User1", 0}
	acc.Deposit(200)
	fmt.Println(acc.Balance)
}
