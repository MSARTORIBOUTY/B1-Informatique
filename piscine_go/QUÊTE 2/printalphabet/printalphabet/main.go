package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 'a'; i <= 'z'; i++ { //si on fait avec les valeurs ascii (a=97; z=122) j'aurais une erreur car i serait de type int et non rune
		z01.PrintRune(i) //et on peut pas avoir un type int avec un type rune
	}
	z01.PrintRune('\n') //On le met après le for car sinon on a un retour à la ligne après chaque lettres
}

//les single quote sont pour les runes et ça affichera ce qu'il y a entre les singles quotes et du coup c'est directement le charactère et non pas la valeur ascii

// 97 -> code ascii de a
//'a' -> rune(97)
//'a' + 1 -> 'b'
