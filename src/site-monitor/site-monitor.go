//hello.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const logFilename = "site-monitor.log"
const sitesFilename = "sites.txt"
const urlPattern = `^(?:https?:\/\/)(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`

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
			insertNewUrl()
		case 3:
			showLogs()
		case 4:
			clearLogs()
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
	versao := 1.2

	fmt.Println("Hello, Mr.", nome)
	fmt.Println("Software version:", versao)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Insert new url for monitoring")
	fmt.Println("3 - Show Logs")
	fmt.Println("4 - Clear Logs")
	fmt.Println("0 - Exit program")
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

func insertNewUrl() {
	var newUrl string

	fmt.Print("\nInsert new url: ")
	fmt.Scan(&newUrl)

	if validateUrl(newUrl) {
		saveUrl(newUrl)
	}
}

func validateUrl(url string) bool {
	regexpUrl := regexp.MustCompile(urlPattern)

	if !regexpUrl.MatchString(url) {
		fmt.Println(url, "is not a valid url!")
		return false
	}

	return true
}

func saveUrl(url string) {
	file, err := os.OpenFile(sitesFilename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString("\n" + url)

	file.Close()
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

	file, err := os.Open(sitesFilename)

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

func clearLogs() {
	fmt.Println()
	fmt.Println("Clearing Logs...")

	err := os.Remove(logFilename)
	if err != nil {
		fmt.Println("Error:", err)
	}
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
