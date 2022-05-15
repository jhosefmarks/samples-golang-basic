//hello.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const numberMonitorings = 3
const delay = 5

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
	urls := loadUrlsByFile()

	fmt.Println()
	fmt.Println("Monitoring...")
	fmt.Println()

	for i := 0; i < numberMonitorings; i++ {
		for _, site := range urls {
			testSite(site)
		}

		time.Sleep(delay * time.Second)
	}
}

func testSite(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "has been loaded successfully!")
	} else {
		fmt.Println("Site", url, "has problems. Status Code:", resp.StatusCode)
	}
}

func loadUrlsByFile() []string {
	var urls []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return urls
}
