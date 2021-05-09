package main

import (
	"github.com/01-edu/z01"
)

func main() {
	s := "x = 42, y = 21"
	for r := 0; r < len(s); r++ {
		z01.PrintRune(rune(s[r]))
	}
	z01.PrintRune('\n')
}

//cyprien
/*package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x string
	y string
}

func main() {
	points := point{}

	setPoint(&points)

	x := "x = "
	y := ", y = "

	for _, c := range x {
		z01.PrintRune(rune(c))
	}

	nb1 := []rune(points.x)
	nb2 := []rune(points.y)

	for _, c := range nb1 {
		z01.PrintRune(rune(c))
	}

	for _, c := range y {
		z01.PrintRune(rune(c))
	}

	for _, c := range nb2 {
		z01.PrintRune(rune(c))
	}
	z01.PrintRune(rune('\n'))
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}*/
