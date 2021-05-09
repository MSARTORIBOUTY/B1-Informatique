/*package piscine

func BasicAtoi(s string) int {
	multiple *= 10
	}

	result = 0
	// Méthode 2 : "Tout-en-un"
	for i := 0; i < len(s); i++ {
		// On s'apprête à ajouter un chiffre à la fin,
		// on "décale" le nombre en multipliant par 10
		// ex : 4 => 40
		result *= 10
		chiffre := int(s[i] - '0')
		// On ajoute le chiffre des unités
		result += chiffre

		// Par ex avec "423", i va passer par
		// 4 => 40 + 2 (42) => 420 + 3 (423)
	}

	return result
}*/
/*package piscine

func BasicAtoi(s string) (i int) {
	str := 0
	r := []rune(s)
	i = 0

	}
	return str
}
*/

//celui de injin

package piscine

func BasicAtoi(s string) int {

	var result int
	var length_int int

	for i := 0; i < len(s); i++ {
		length_int = length_int + 1
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + (int(s[i]) - 48)
		}
	}

	if len(s) != 0 && s[0] == '-' {
		result = result * -1
	}

	return result
}

/*str := 0
for _, v := range s {
	a := 0
	for i := '1'; i <= v; i++ {
		a++
	}
	str = str*10 + a //10 line break
}
return str*/
