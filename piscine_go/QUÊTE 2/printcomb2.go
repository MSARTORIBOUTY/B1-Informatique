package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	//00 01: n = 1; m = 2; o = 3; p = 4
	//faire une série de chiffre croissante : 00 01. 00 02. 00 03... 00 98. 00 99...

	for n := '0'; n <= '9'; n++ { //On fait 4 variables car série de 4 chiffres. Pas besoins de les définir comme rune si on utilise les single quote. (single quote are for rune; double quote are for fmt)
		for m := '0'; m <= '9'; m++ {
			for o := '0'; o <= '9'; o++ {
				for p := '0'; p <= '9'; p++ {
					if (n < o || (n <= o && m <= p)) && !(n == o && m == p) { //!(..) parce que c'est une série croissante
						z01.PrintRune(n) //on affiche n & m (pas de +48 car on a les sigle quote sui affiche dirrectement les chiffres donc on a pas besoins du tableau ascii)
						z01.PrintRune(m)
						z01.PrintRune(' ') // puis on met un espace entre eux parce que miskine y en a un

						z01.PrintRune(o) //et enfin on affiche o & p
						z01.PrintRune(p)
						if !(n == '9' && m == '8' && o == '9' && p == '9') { //dernière valeur. /!\ !(...) car c'est la dernière valeur et donc après la dernière valeur, on affiche pas ce qui suit
							z01.PrintRune(',') //On affiche virgule et espace entre chaque série de chiffres
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// ou

/*for a := 0; a <= 98; a++ {
	for b := a + 1; b <= 99; b++ {
		Dizainea := a / 10 // 10 car dans une rune on ne peut avoir qu'un chiffre et non 2 à la suite d'où le fait qu'on le divise par 10
		Unitea := a % 10
		Dizaineb := b / 10
		Uniteb := b % 10
		z01.PrintRune(rune(Dizainea + 48)) //On prend notre nombre et on l'encastre en rune avec la table ascii
		z01.PrintRune(rune(Unitea + 48))
		z01.PrintRune(' ')
		z01.PrintRune(rune(Dizaineb + 48))
		z01.PrintRune(rune(Uniteb + 48))
		if !(a == 98 && b == 99) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}

	}
}
z01.PrintRune('\n')*/
