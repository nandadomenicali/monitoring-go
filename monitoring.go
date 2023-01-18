package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 5

func main() {
	introduction()
}

func introduction() {
	version := 1.2
	fmt.Println("This program is specific to monitor site logs")
	fmt.Println("The version of this program is: ", version)
}

func displayMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)
	fmt.Println("")

	return command
}

func initialize() {
	fmt.Println("Monitoring...")
	sites := readSites()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing website", i, ":", site)
			test(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func test(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("website:", site, "Website loaded successfully")
		registerLog(site, true)
	} else {
		fmt.Println("website:", site, "It has problems. Status Code:", response.StatusCode)
		registerLog(site, false)
	}
}

func readSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
