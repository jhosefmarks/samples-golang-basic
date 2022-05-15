//hello.go
package main

import "fmt"

func main() {
	showGreeting()

	fmt.Println()
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit program")

	fmt.Print("\nCommand: ")
	var comando int
	fmt.Scan(&comando)

	switch comando {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Leaving the program...")
	default:
		fmt.Println("Error: Invalid command!")
	}
}

func showGreeting() {
	nome := "Jhosef Marks"
	versao := 1.1

	fmt.Println("Hello, Mr.", nome)
	fmt.Println("Software version:", versao)
}
