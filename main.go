package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func (c *CheckingAccount) withdraw(value float64) string {
	if value >= 0 && value <= c.balance {
		c.balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient funds"
	}
}

func main() {
	myAccount := CheckingAccount{owner: "Wagner", agency: 123, accountNumber: 123456, balance: 1000.00}

	fmt.Println(myAccount)

	fmt.Println(myAccount.withdraw(200))
	fmt.Println(myAccount.balance)

	fmt.Println(myAccount.withdraw(1200))
	fmt.Println(myAccount.balance)

	fmt.Println(myAccount.withdraw(-200))
	fmt.Println(myAccount.balance)
}
