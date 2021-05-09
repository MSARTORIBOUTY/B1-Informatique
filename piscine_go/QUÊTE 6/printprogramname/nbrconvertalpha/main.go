/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	n := 1

	for i := 97; i <= 122; i++ {
		z01.PrintRune(rune(i + 48))
	}
	for _, k := range arg[n] {
		z01.PrintRune(k)
	}
	z01.PrintRune(' ')

}
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	isflag := false

	for _, arg := range arguments {
		if arg == "--upper" {
			isflag = true
		}
	}

	for _, arg := range arguments {
		numv := 0

		for _, val := range arg {
			numv = numv*10 + int(val-'0')
		}

		if numv >= 1 && numv <= 26 {
			if isflag == false {
				z01.PrintRune(rune(numv + 96))
			} else {
				z01.PrintRune(rune(numv + 64))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	x := len(args)
	var mult int
	var result int
	var i int

	if x == 0 {
	} else {
		if args[0] == "--upper" {
			i = 1
		} else {
			i = 0
		}

		u := i

		for i < x {
			mult = 1
			result = 0

			for j := len(args[i]) - 1; j >= 0; j-- {
				result += (int(args[i][j]) - 48) * mult
				mult *= 10
			}

			val := result

			if u == 1 {
				result += 64
			} else {
				result += 96
			}

			if 0 < val && val <= 26 {
				z01.PrintRune(rune(result))
			} else {
				z01.PrintRune(rune(' '))
			}
			i++
		}
		z01.PrintRune(rune('\n'))
	}
}
