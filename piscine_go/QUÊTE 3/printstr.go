package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {

	for _, caractere := range s {

		z01.PrintRune(caractere) //caractere est une variable, il renvoit à s grace à =
	}
}

/*func PrintStr(s string) {
	// on doit afficher un par un les charactères de hello world!
	//str := []rune{s}
	for count := 0; count < len(s); count++ { //donc on doit fair une boucle qui va compter charactère par charactère. /!\ < len(s) car le dernier charactère c'est len(s)-1. Si on fait <= len(s) ça fonctionne pas
		z01.PrintRune(rune(s[count])) //On définit s comme rune et count va prendre séparément chaque caractère dans le tableau de s

	}
}*/
