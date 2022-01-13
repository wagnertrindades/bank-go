package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	myAccount := accounts.CheckingAccount{Owner: "Wagner", Balance: 1000.00}
	otherAccount := accounts.CheckingAccount{Owner: "Joao", Balance: 100.00}

	fmt.Println(myAccount)

	fmt.Println(myAccount.Withdraw(200))
	fmt.Println(myAccount.Balance)

	fmt.Println(myAccount.Withdraw(1200))
	fmt.Println(myAccount.Balance)

	fmt.Println(myAccount.Withdraw(-200))
	fmt.Println(myAccount.Balance)

	fmt.Println(myAccount.Deposit(-200))
	fmt.Println(myAccount.Balance)

	status, accountBalance := myAccount.Deposit(200)
	fmt.Println("Status:", status)
	fmt.Println("Account Balance:", accountBalance)

	fmt.Println(myAccount.Transfer(200, &otherAccount))
	fmt.Println(myAccount.Balance)
	fmt.Println(otherAccount.Balance)
}
