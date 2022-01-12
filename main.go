package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdraw(value float64) (string, bool) {
	if value > 0 && value <= c.balance {
		c.balance -= value
		return "Successful withdrawal", true
	}

	return "Insufficient funds", false
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Successful deposited", c.balance
	}

	return "The value should be higher than zero", c.balance
}

func (c *CheckingAccount) Transfer(value float64, destinyAccount *CheckingAccount) string {
	_, withdrawed := c.Withdraw(value)

	if withdrawed {
		destinyAccount.Deposit(value)
		return "Successful transfer"
	}

	return "Insufficient funds"
}

func main() {
	myAccount := CheckingAccount{owner: "Wagner", agency: 123, accountNumber: 123456, balance: 1000.00}
	otherAccount := CheckingAccount{owner: "Joao", balance: 100.00}

	fmt.Println(myAccount)

	fmt.Println(myAccount.Withdraw(200))
	fmt.Println(myAccount.balance)

	fmt.Println(myAccount.Withdraw(1200))
	fmt.Println(myAccount.balance)

	fmt.Println(myAccount.Withdraw(-200))
	fmt.Println(myAccount.balance)

	fmt.Println(myAccount.Deposit(-200))
	fmt.Println(myAccount.balance)

	status, accountBalance := myAccount.Deposit(200)
	fmt.Println("Status:", status)
	fmt.Println("Account Balance:", accountBalance)

	fmt.Println(myAccount.Transfer(200, &otherAccount))
	fmt.Println(myAccount.balance)
	fmt.Println(otherAccount.balance)
}
