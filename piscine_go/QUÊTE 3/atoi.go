package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := 0

	// TODO : factoriser (et renommer lol ?) ces conditions
	commenceParUnMoins := s[0] == '-'

	// Méthode 2 : "Tout-en-un"
	for i := 0; i < len(s); i++ {
		// isNotNumber boolean par ce que > et < retournent des bool
		// Et || compare des bool
		isNotNumber := s[i] < '0' || s[i] > '9'

		if isNotNumber && i != 0 {
			return 0
		}

		if !isNotNumber {
			// On s'apprête à ajouter un chiffre à la fin,
			// on "décale" le nombre en multipliant par 10
			// ex : 4 => 40
			result *= 10
			chiffre := int(s[i] - '0')
			// fmt.Println("Chiffre", i, " : ", chiffre)
			// On ajoute le chiffre des unités
			result += chiffre

			// Par ex avec "423", i va passer par
			// 4 => 40 + 2 (42) => 420 + 3 (423)
			// fmt.Println(result)
		}
	}

	if commenceParUnMoins {
		result *= -1
	}

	return result
}
