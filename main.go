package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func main() {
	myAccount := CheckingAccount{owner: "Wagner", agency: 123, accountNumber: 123456, balance: 1000.00}

	fmt.Println(myAccount)
}
