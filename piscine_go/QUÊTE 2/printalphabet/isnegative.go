package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {

	var aRune string = "TF" //type string: afficher un text/lettres

	if nb >= 0 { //si le nombre est positif
		z01.PrintRune(rune(aRune[1])) //afficher la variable aRune comme vrai. 1 en binaire = 'vrai'
	} else {
		z01.PrintRune(rune(aRune[0])) //sinon on affiche faux
	}
	a('\n') //saut à la ligne entre chaques réponses
}

func a(r rune) {
	z01.PrintRune(r)

}

// ou
//celui là est plus simple et plus compréhensible

/*if nb < 0 {
	z01.PrintRune('T')
} else {
	z01.PrintRune('F')
}
z01.PrintRune('\n')*/
