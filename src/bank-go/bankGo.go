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
	fmt.Println("- Withdraw: $ 1000", robertAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", robertAccount.balance)
	fmt.Println("- Withdraw: $ 1500", robertAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", robertAccount.balance)

	fmt.Println()
	fmt.Println(kenAccount.holder, "bank statement")
	fmt.Println("= Balance.: $", kenAccount.balance)
	fmt.Println("- Withdraw: $ 1000", kenAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", kenAccount.balance)
	fmt.Println("- Withdraw: $ 1500", kenAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", kenAccount.balance)
}

func showTitle() {
	fmt.Println("*** Bank GO ***")
	fmt.Println("*** Welcome ***")
	fmt.Println()
}
