package reloaded

func Atoi(s string) int {
	//1.Transformer un string en int
	//return 0 si la string n'est pas un nombre valide

	//si la longueur de s est nul
	if len(s) == 0 {
		return 0
	}
	result := 0
	signe := s[0] == '-'

	for i := 0; i < len(s); i++ {
		// les char définies dans le tableau ascii
		char := s[i] < '0' || s[i] > '9' // le char de la string à l'index i n'est pas compris entre 0 et 9
		// si ce sont des lettres et un caractère, on retourne 0
		if char && i != 0 { // Ref: char  et i non nul
			return 0 //On retourne quand même zéro car ce ne sera pas un chiffre
		}
		// si les charactères sont compris entre 0 et 9 inclus, alors on multiplie le résultat par 10 pour ajouter les nb succéssivements
		if !char {
			result *= 10
			// on caste en int
			nb := int(s[i] - '0')
			result += nb
		}
	}

	if signe { //Pas besoins de remettre == '-'. ça fait une erreur et puis il est déjà définit
		result *= -1 //On remet le moins devant
	}
	return result
}
