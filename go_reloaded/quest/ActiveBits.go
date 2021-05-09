package reloaded

func ActiveBits(n int) uint {
	//Doit afficher le nombre de bits actifs (valeur 1)
	//Dans la représentation binaire d'un nombre en int
	var count uint //On retourne un utin donc on déclare count comme uint et que c'est le nbr bits actifs

	//On transforme n de base 10 en base 2!

	if n == 0 { //Si le chiffre = 0
		return 0
	}
	for n > 0 {
		if n%2 == 1 { //Si le résultat du reste de la division de n par 2 = 1
			count++ //ça veut dire que le bit est actif donc count prend un nombre en plus
		}
		n /= 2 //Voir PrintnbrBase si besoins
	}
	return count
}
