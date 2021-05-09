package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	for i := len(arguments) - 1; i > 0; i-- {
		for _, k := range arguments[i] { //i c'est la phrase
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}