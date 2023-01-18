package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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
