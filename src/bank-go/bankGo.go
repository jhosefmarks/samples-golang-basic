package main

import "fmt"

type CurrentAccount struct {
	holder    string
	branchNo  int
	accountNo int
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

func main() {
	showTitle()

	robertAccount := CurrentAccount{
		holder:    "Robert Griesemer",
		branchNo:  589,
		accountNo: 123456,
		balance:   2300.,
	}

	kenAccount := CurrentAccount{
		holder:    "Ken Thompson",
		branchNo:  324,
		accountNo: 341866,
		balance:   2500,
	}

	fmt.Println(robertAccount.holder, "bank statement")
	fmt.Println("= Balance.: $", robertAccount.balance)
	fmt.Println("- Withdraw: $ 1000 -", robertAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", robertAccount.balance)
	fmt.Println("- Withdraw: $ 1500 -", robertAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", robertAccount.balance)
	fmt.Println("- Deposit.: $ 500 -", robertAccount.Deposit(500))
	fmt.Println("= Balance.: $", robertAccount.balance)
	fmt.Println("- Deposit.: $ -50 -", robertAccount.Deposit(-50))
	fmt.Println("= Balance.: $", robertAccount.balance)

	fmt.Println()
	fmt.Println(kenAccount.holder, "bank statement")
	fmt.Println("= Balance.: $", kenAccount.balance)
	fmt.Println("- Withdraw: $ 1000 -", kenAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", kenAccount.balance)
	fmt.Println("- Deposit.: $ 700 -", kenAccount.Deposit(700))
	fmt.Println("= Balance.: $", kenAccount.balance)
	fmt.Println("- Transfer: $ 500 -", kenAccount.Tranfer(500, &robertAccount))
	fmt.Println("= Balance.: $", kenAccount.balance)
	fmt.Println("- Withdraw: $ 1500 -", kenAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", kenAccount.balance)

	fmt.Println()
	fmt.Println(robertAccount.holder, "bank statement")
	fmt.Println("= Balance.: $", robertAccount.balance)
}

func showTitle() {
	fmt.Println("*** Bank GO ***")
	fmt.Println("*** Welcome ***")
	fmt.Println()
}
