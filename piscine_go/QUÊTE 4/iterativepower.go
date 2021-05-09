package piscine

func IterativePower(nb int, power int) int {
	if power < 0 { //si la puissance est négative ça retourne 0
		return 0
	} else if power == 1 {
		return 1
	}

	resultat := 1
	for i := 0; i < power; i++ { //on donne à i la valeur de la puissance
		resultat *= nb
	}
	return resultat

}

//power = x^y (c'est la puissance)

/*if power < 0 {
	return 0
} else if power == 0 {
	return 1
}
result := nb
for i := 0; i < power-1; i++ {
	result *= nb
}
return result*/
