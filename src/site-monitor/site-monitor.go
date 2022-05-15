//hello.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const logFilename = "site-monitor.log"

func main() {
	showGreeting()

	for {
		fmt.Println()
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Leaving the program...")
			os.Exit(0)
		default:
			fmt.Println("Error: Invalid command!")
			os.Exit(-1)
		}
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

	for _, site := range urls {
		testSite(site)
	}
}

func testSite(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "has been loaded successfully!")
		log(url, true)
	} else {
		fmt.Println("Site", url, "has problems. Status Code:", resp.StatusCode)
		log(url, false)
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

func log(url string, status bool) {
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url +
		" - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println()
	fmt.Println("Showing Logs...")

	file, err := ioutil.ReadFile(logFilename)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println()
	fmt.Println(string(file))
}
