package main

import "github.com/01-edu/z01"

func main() {
	var i rune
	for i = 48; i < 58; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}

//ou
/*for i := '0'; i <= '9'; i++ {//on ne met pas 10 car une rune c'est un caractère. (ascii)
	z01.PrintRune(i)

}
z01.PrintRune('\n')*/

// ce code est faux
/*for i := 0; i < 10; i++ {
	z01.PrintRune(rune(i)) //ça nous donnera des symboles de la table ascii en suivant le nbr indiqué

}
z01.PrintRune('\n')*/
