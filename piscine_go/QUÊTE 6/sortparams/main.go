package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] //bibliothèque des arguments, [1:]:1 on commence par l'argument 1; :jusqu'à la fin. Par exemple choumi is the best le premier arg est choumi

	for i := 0; i < len(args); i++ { //on ne met pas de -1 arprès len(arguments) car sinon les lettres ne sont pas triés
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for k := 0; k < len(args); k++ {
		for _, letter := range args[k] { // on liste
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
