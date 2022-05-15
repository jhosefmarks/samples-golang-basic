package accounts

import "bank-go/customers"

type CurrentAccount struct {
	Holder    customers.Holder
	BranchNo  int
	AccountNo int
	Balance   float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	validOperation := value > 0 && value <= c.Balance

	if validOperation {
		c.Balance -= value
		return "Withdrawal successful."
	}

	return "Insufficient funds."
}

func (c *CurrentAccount) Deposit(value float64) string {
	if value <= 0 {
		return "Invalid value."
	}

	c.Balance += value

	return "Deposit successfully."
}

func (c *CurrentAccount) Tranfer(value float64, destAccount *CurrentAccount) bool {
	validOperation := value < c.Balance && value > 0

	if validOperation {
		c.Balance -= value
		destAccount.Deposit(value)

		return true
	}

	return false
}
