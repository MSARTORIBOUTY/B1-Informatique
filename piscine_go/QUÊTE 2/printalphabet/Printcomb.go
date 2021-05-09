package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var n rune //Création de trois variables car série de trois chiffres à chaques fois. En rune car on a pas le droit à fmt donc on va utiliser tableau ascii
	var m rune
	var o rune
	for n = 0; n <= 9; n++ { //n est le premier chiffre, toujours compris entre 0 et 9, on fait une boucle croissante
		for m = 0; m <= 9; m++ { // de même
			if m != n && m > n { //ils disent que le 1er et plus petit que le 2ème, on traduit!
				for o = 0; o <= 9; o++ { //on amène la troisième variable après (car sinon on ne sait pas que m>n)
					if o != m && o > m { //Le 3ème et plus grand que le 2ème, on traduit
						z01.PrintRune(n + 48) //Vue que ce sont des runes et qu'on veut afficher des lettres (/!\ rune c'est surtout pour les lettres du coup on utilise ascii) on rajoute la valeur 0 à la variable pour ne pas fausser le résultat mais surtout obtenir les chiffres
						z01.PrintRune(m + 48)
						z01.PrintRune(o + 48)
						if !(n == 7 && m == 8 && o == 9) { //Valeur final
							z01.PrintRune(',') //, entre chaques séries de chiffre (012, 013, 014,...)
							z01.PrintRune(' ') //espace après la virgule pour chaque série de chiffre
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
//ou

/*for a := '0'; a <= '9'; a++ {
	for b := '0'; b <= '9'; b++ {
		if a < b && a != b {
			for c := '0'; c <= '9'; c++ {
				if b < c && b != c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if !(a == 7 && b == 8 && c == 9) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
z01.PrintRune('\n')*/

