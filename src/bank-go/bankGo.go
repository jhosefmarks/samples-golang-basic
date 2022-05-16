package main

import (
	"fmt"

	"bank-go/accounts"
	"bank-go/customers"
)

type verifyAccount interface {
	Withdraw(value float64) string
}

func PayBill(account verifyAccount, value float64) string {
	return account.Withdraw(value)
}

func main() {
	showTitle()

	robertAccount := accounts.CurrentAccount{
		Holder:    customers.Holder{Name: "Robert Griesemer"},
		BranchNo:  589,
		AccountNo: 123456,
	}

	kenAccount := accounts.CurrentAccount{
		Holder:    customers.Holder{Name: "Ken Thompson"},
		BranchNo:  324,
		AccountNo: 341866,
	}

	robAccount := accounts.SavingAccount{
		Holder:    customers.Holder{Name: "Rob Pike"},
		BranchNo:  421,
		AccountNo: 876312,
		Operation: 1,
	}

	fmt.Println(robertAccount.Holder.Name, "bank statement -",
		robertAccount.BranchNo, "/", robertAccount.AccountNo)
	fmt.Println("+ Init....: $ 2300 -", robertAccount.Deposit(2300))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())
	fmt.Println("- Withdraw: $ 1000 -", robertAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())
	fmt.Println("- Withdraw: $ 1500 -", robertAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())
	fmt.Println("+ Deposit.: $ 500 -", robertAccount.Deposit(500))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())
	fmt.Println("+ Deposit.: $ -50 -", robertAccount.Deposit(-50))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())
	fmt.Println("- Transfer: $ 2500 -", robertAccount.Tranfer(2500, &kenAccount))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())
	fmt.Println("- Pay bill: $ 190 -", PayBill(&robertAccount, 190))
	fmt.Println("= Balance.: $", robertAccount.GetBalance())

	fmt.Println()
	fmt.Println(kenAccount.Holder.Name, "bank statement -",
		kenAccount.BranchNo, "/", kenAccount.AccountNo)
	fmt.Println("+ Init....: $ 2500 -", kenAccount.Deposit(2500))
	fmt.Println("= Balance.: $", kenAccount.GetBalance())
	fmt.Println("- Withdraw: $ 1000 -", kenAccount.Withdraw(1000))
	fmt.Println("= Balance.: $", kenAccount.GetBalance())
	fmt.Println("+ Deposit.: $ 700 -", kenAccount.Deposit(700))
	fmt.Println("= Balance.: $", kenAccount.GetBalance())
	fmt.Println("- Transfer: $ 500 -", kenAccount.Tranfer(500, &robertAccount))
	fmt.Println("= Balance.: $", kenAccount.GetBalance())
	fmt.Println("- Withdraw: $ 1500 -", kenAccount.Withdraw(1500))
	fmt.Println("= Balance.: $", kenAccount.GetBalance())

	fmt.Println()
	fmt.Println(robertAccount.Holder.Name, "bank statement -",
		robertAccount.BranchNo, "/", robertAccount.AccountNo)
	fmt.Println("= Balance.: $", robertAccount.GetBalance())

	fmt.Println()
	fmt.Println(robAccount.Holder.Name, "bank statement -",
		robAccount.BranchNo, "/", robAccount.AccountNo, "-", robAccount.Operation)
	fmt.Println("+ Init....: $ 1000 -", robAccount.Deposit(1000))
	fmt.Println("= Balance.: $", robAccount.GetBalance())
	fmt.Println("- Withdraw: $ 1200 -", robAccount.Withdraw(1200))
	fmt.Println("= Balance.: $", robAccount.GetBalance())
	fmt.Println("+ Deposit.: $ 200 -", robAccount.Deposit(200))
	fmt.Println("= Balance.: $", robAccount.GetBalance())
	fmt.Println("- Pay bill: $ 1400 -", PayBill(&robAccount, 1400))
	fmt.Println("= Balance.: $", robAccount.GetBalance())
}

func showTitle() {
	fmt.Println("*** Bank GO ***")
	fmt.Println("*** Welcome ***")
	fmt.Println()
}
