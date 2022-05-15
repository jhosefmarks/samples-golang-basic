package accounts

import "bank-go/customers"

type CurrentAccount struct {
	Holder    customers.Holder
	BranchNo  int
	AccountNo int
	balance   float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	validOperation := value > 0 && value <= c.balance

	if validOperation {
		c.balance -= value
		return "Withdrawal successful."
	}

	return "Insufficient funds."
}

func (c *CurrentAccount) Deposit(value float64) string {
	if value <= 0 {
		return "Invalid value."
	}

	c.balance += value

	return "Deposit successfully."
}

func (c *CurrentAccount) Tranfer(value float64, destAccount *CurrentAccount) bool {
	validOperation := value < c.balance && value > 0

	if validOperation {
		c.balance -= value
		destAccount.Deposit(value)

		return true
	}

	return false
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
