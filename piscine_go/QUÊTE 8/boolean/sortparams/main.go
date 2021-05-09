package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	x := len(args)
	i := 0
	var count int
	j := 1

	for {
		i++
		count = 0

		if args[i] > args[i+1] {
			args[i], args[i+1] = args[i+1], args[i]
		}
		j = 0
		for j < x-1 {
			if args[j] < args[j+1] {
				count++
			}
			j++
		}

		if count == x-2 {
			break
		}

		if i == x-2 {
			i = 0
		}
	}

	for i := 1; i < len(args); i++ {
		runes := []rune(args[i])
		for j := 0; j < len(args[i]); j++ {
			z01.PrintRune(rune(runes[j]))
		}
		z01.PrintRune(rune('\n'))
	}
}
