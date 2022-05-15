//hello.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showGreeting()
	fmt.Println()
	showMenu()

	command := readCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Leaving the program...")
		os.Exit(0)
	default:
		fmt.Println("Error: Invalid command!")
		os.Exit(-1)
	}
}

func showGreeting() {
	nome := "Jhosef Marks"
	versao := 1.1

	fmt.Println("Hello, Mr.", nome)
	fmt.Println("Software version:", versao)
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit program")
}

func readCommand() int {
	var command int

	fmt.Print("\nCommand: ")
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	sites := []string{
		"https://go.dev/",
		"https://pkg.go.dev/",
		"https://random-status-code.herokuapp.com",
		"https://angular.io/",
	}

	fmt.Println()
	fmt.Println("Monitoring...")
	fmt.Println()

	for i, site := range sites {
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site", i+1, ":", site, "has been loaded successfully!")
		} else {
			fmt.Println("Site", i+1, ":", site, "has problems. Status Code:", resp.StatusCode)
		}
	}
}
