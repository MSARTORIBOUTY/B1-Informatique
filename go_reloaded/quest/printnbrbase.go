package reloaded

import (
	"github.com/01-edu/z01"
)

//1. écrire fonction int to string
//2. base non valide --> print NV
//règles de validité: base > 2 char/ Chaque char doit être différent/La base ne doit pas contenir + ou -/Elle doit supporter des chiffres négatifs

func PrintNbrBase(nbr int, base string) {
	str := ""

	//On crée une variable qui prend la longueur de la base car la base est une string, elle ne prend pas les chars séparément
	//lengthBase prend les chars de la base. Il est de type int
	lengthBase := 0
	for range base {
		lengthBase++
	}

	if lengthBase < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
	//si les chars de la base ne sont pas unique alors erreur
	for i := 0; i < lengthBase; i++ {
		for j := 1 + i; j < lengthBase; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	//Pour les nombres négagifs, on les rend positif.
	if nbr < 0 {
		nbr *= -1
		z01.PrintRune('-')
	} else if nbr == 0 {
		z01.PrintRune(rune(base[0]))
	}

	for nbr > 0 {
		str = str + string(base[nbr%lengthBase])//On prend le reste de la division du nombre et de la base 
		nbr = nbr / lengthBase
	}
	//Le for avant à mis la str à l'envers donc on la remet à l'endroit.
	for i := len(str) - 1; i > -1; i-- {
		z01.PrintRune(rune(str[i]))
	}
}
