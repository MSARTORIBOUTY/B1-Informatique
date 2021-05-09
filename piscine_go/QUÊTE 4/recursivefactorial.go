package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb > 20 || nb < 0 {
		result = 0

	} else if nb > 0 {
		result = nb * RecursiveFactorial(nb-1) //*
	} else if nb == 1 {
		result = 1
	} else {
		result = 0
	}
	return result
}

//récursive :Qui peut être répété un nombre indéfini de fois par l'application de la même règle.
//* il s'agit d'une boucle. fac(1)= 1 ; fac(2)= fac(1)*2 ; fac(3)= fac(2)*3 donc fac(nb)= fac(nb-1)*nb
//c'est un système de rappel de fonction, on aurait très bien pu reprendre le système de la boucle for avec i,
//ça donnerait le même résultat, mais ce système prend moins de ligne. On fait la même chose mais écrit différement.
//Rappelons que pour cet exo, le for était interdit
