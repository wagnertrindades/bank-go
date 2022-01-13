package accounts

import "bank/customers"

type CheckingAccount struct {
	Owner         customers.Customer
	Agency        int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) Withdraw(value float64) (string, bool) {
	if value > 0 && value <= c.Balance {
		c.Balance -= value
		return "Successful withdrawal", true
	}

	return "Insufficient funds", false
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "Successful deposited", c.Balance
	}

	return "The value should be higher than zero", c.Balance
}

func (c *CheckingAccount) Transfer(value float64, destinyAccount *CheckingAccount) string {
	_, withdrawed := c.Withdraw(value)

	if withdrawed {
		destinyAccount.Deposit(value)
		return "Successful transfer"
	}

	return "Insufficient funds"
}
