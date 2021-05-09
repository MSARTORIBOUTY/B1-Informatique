package piscine

func BasicAtoi2(s string) int {
	result := 0
	// Méthode 2 : "Tout-en-un"
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

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
}
