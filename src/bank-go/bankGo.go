package main

import "fmt"

type CurrentAccount struct {
	holder    string
	branchNo  int
	accountNo int
	balance   float64
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

	fmt.Println(robertAccount)
	fmt.Println(kenAccount)
}

func showTitle() {
	fmt.Println("*** Bank GO ***")
	fmt.Println("*** Welcome ***")
	fmt.Println()
}
