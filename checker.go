package main

import (
	"fmt"
	"os"
	"bufio"
	"net/http"
	"github.com/TwiN/go-color"
)

func main() {
	file := os.Args[1]
	read, err := os.Open(file)
	
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		res, err := http.Get("https://linktr.ee/"+line)
		if err != nil {
			fmt.Println(err)
		}
		if res.StatusCode != 404 {
			fmt.Println(color.Red + "[+] " + color.Reset + "linktr.ee/"+line + " isn't available")
		} 
		if res.StatusCode != 200 {
			fmt.Println(color.Green + "[+] " + color.Reset + "linktr.ee/"+line + " is available")
		}
	}
}