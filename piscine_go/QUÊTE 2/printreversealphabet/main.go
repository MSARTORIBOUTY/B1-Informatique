package main

import "github.com/01-edu/z01"

func main() {
	var i rune                      // pour afficher la lettre d'après le tableau ascii sinon ça donne des chiffres
	for i = 97 + 25; i >= 97; i-- { //97 c'est 'a' minuscule '+25' pour arriver à 122 qui est z minuscule. Donc i part de 122 et décroit jusqu'à 97 inclus. i-- pour décroissant
		z01.PrintRune(i) //afficher i
	}
	z01.PrintRune(10) // LF: marque une fin de ligne
}

// ou

/*for i := 'z'; i >= 'a'; i-- { //for intitialisation; condition; incrémentation
	z01.PrintRune(i)

}
z01.PrintRune('\n')
}*/
