package accounts

import "bank-go/customers"

type SavingAccount struct {
	Holder    customers.Holder
	BranchNo  int
	AccountNo int
	Operation int
	balance   float64
}

func (c *SavingAccount) Withdraw(value float64) string {
	validOperation := value > 0 && value <= c.balance

	if validOperation {
		c.balance -= value
		return "Withdrawal successful."
	}

	return "Insufficient funds."
}

func (c *SavingAccount) Deposit(value float64) string {
	if value <= 0 {
		return "Invalid value."
	}

	c.balance += value

	return "Deposit successfully."
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
