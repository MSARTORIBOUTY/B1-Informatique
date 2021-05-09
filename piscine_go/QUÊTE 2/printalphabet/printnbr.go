package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	//Afficher le signe de n

	if n < 0 { //Si n est négatif alors
		z01.PrintRune('-')             //On affiche le - comme signe négatif avant la valeur de n
		if n == -9223372036854775808 { //Peut importe ce que l'on met ça ne changera pas le résultat. Il faut juste que n soit négatif
			z01.PrintRune('9')
		} else { //Pour afficher le premier (chiffre qui est négatif) mais aussi pour que la suite soit positive (on multiplie n par -1)
			n *= -1
		}
	}
	//Afficher a valeur de n
	if n >= 0 {
		if n > 9 {
			PrintNbr(n / 10) //Je pense qu'on divise n par 10 pour avoir chaque chiffre séparément (ex: 123 ==> 1 2 3)
		}
		r := '0' //r type rune, si r affiche 0
		for i := 0; i < n%10; i++ {
			r++
		}
		z01.PrintRune(r)
	}
}

/*package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	//for n := 0; n >= 10; n++ //on ne peut pas faire ça car on redéclare un n comme variable != du n int plus haut et on veut pas

	var board []int
	if n < 0 { //Si n est négatif on va le rendre positif
		z01.PrintRune('-') //Mais on affiche quand même le - sinon ça marche pas pour afficher un nbr négatif
		n *= -1
	}

	for { //On peut faire un for avec seulement une condition ou rien du tout
		unit := n % 10 //Car en prenant le chiffre que l'on veut et le diviser par dix, en prenant le reste donc modulo 10 on retrouve le chiffre
		n /= 10        // ET on divise n par 10
		//z01.PrintRune(rune('0' + unit)) //on affiche le reste de la division de n avec 10 donc on affiche unit
		board = append(board, unit) //J'ajoute à ma tableau, la valeur n et le résultat du tableau ce sera la valeur n (parce que sinon on a pas les valeurs complètes)
		if n == 0 {                 //Si n = 0 on sort de la boucle
			break
		} //ça nous affiche le résultat à l'envers
	}
	//for qui affiche chaque caractère du tableau dans l'ordre, qu'on peut afficher avec fmt.Print
	for i := len(board) - 1; i >= 0; i-- {
		chiffre := board[i] //quand on a un nombre trop grand et qui est négatif, ça fait un overflow mais l'algorithme ne le prenait pas en compte
		if chiffre < 0 {
			chiffre *= -1
		}
		z01.PrintRune(rune(chiffre + 48)) //On peut aussi remplacer 48 par '0'
	}
}*/
