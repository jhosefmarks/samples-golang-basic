package main

import (
	"fmt"

	"bank-go/accounts"
	"bank-go/customers"
)

func main() {
	showTitle()

	robertAccount := accounts.CurrentAccount{
		Holder:    customers.Holder{Name: "Robert Griesemer"},
		BranchNo:  589,
		AccountNo: 123456,
		Balance:   2300.,
	}

	kenAccount := accounts.CurrentAccount{
		Holder:    customers.Holder{Name: "Ken Thompson"},
		BranchNo:  324,
		AccountNo: 341866,
		Balance:   2500,
	}

	fmt.Println(robertAccount.Holder.Name, "bank statement -",
		robertAccount.BranchNo, "/", robertAccount.AccountNo)
	fmt.Println("= Balance.: $", robertAccount.Balance)
	fmt.Println("- Withdraw: $ 1000 -", robertAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", robertAccount.Balance)
	fmt.Println("- Withdraw: $ 1500 -", robertAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", robertAccount.Balance)
	fmt.Println("- Deposit.: $ 500 -", robertAccount.Deposit(500))
	fmt.Println("= Balance.: $", robertAccount.Balance)
	fmt.Println("- Deposit.: $ -50 -", robertAccount.Deposit(-50))
	fmt.Println("= Balance.: $", robertAccount.Balance)

	fmt.Println()
	fmt.Println(kenAccount.Holder.Name, "bank statement -",
		kenAccount.BranchNo, "/", kenAccount.AccountNo)
	fmt.Println("= Balance.: $", kenAccount.Balance)
	fmt.Println("- Withdraw: $ 1000 -", kenAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", kenAccount.Balance)
	fmt.Println("- Deposit.: $ 700 -", kenAccount.Deposit(700))
	fmt.Println("= Balance.: $", kenAccount.Balance)
	fmt.Println("- Transfer: $ 500 -", kenAccount.Tranfer(500, &robertAccount))
	fmt.Println("= Balance.: $", kenAccount.Balance)
	fmt.Println("- Withdraw: $ 1500 -", kenAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", kenAccount.Balance)

	fmt.Println()
	fmt.Println(robertAccount.Holder.Name, "bank statement -",
		robertAccount.BranchNo, "/", robertAccount.AccountNo)
	fmt.Println("= Balance.: $", robertAccount.Balance)
}

func showTitle() {
	fmt.Println("*** Bank GO ***")
	fmt.Println("*** Welcome ***")
	fmt.Println()
}
