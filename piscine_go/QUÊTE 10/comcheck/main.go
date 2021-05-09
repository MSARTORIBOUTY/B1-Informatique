package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	alert := false

	for _, val := range args {
		if val == "01" || val == "galaxy" || val == "galaxy 01" {
			alert = true
		}
	}
	if alert == true {
		fmt.Println("Alert!!!")
	}
}
