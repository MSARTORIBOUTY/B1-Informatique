package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	len := 0

	for i := range arguments {
		len = i
	}

	for i := 1; i <= len; i++ {
		for _, k := range arguments[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune(10)
	}
}
