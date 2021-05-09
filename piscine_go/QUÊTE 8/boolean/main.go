//Boolean
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	if isEven(len(args)) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}

}

//Cyprien
/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if isEven(len(args)) == true {
		printStr("EvenMsg")
	} else {
		printStr("OddMsg")
	}
}
func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func printStr(s string) {
	var texte string

	if s == "EvenMsg" {
		texte = "I have an even number of arguments\n"
	} else {
		texte = "I have an odd number of arguments\n"
	}

	for _, c := range texte {
		z01.PrintRune(rune(c))
	}
}*/
