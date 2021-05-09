package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb >= 25 || nb < 0 { //*
		return 0

	} else if nb > 0 {
		for i := 1; i <= nb; i++ {
			result *= i

		}
	} else if nb == 1 {
		result = 1
	} else {
		result = 0
	}
	return result
}

//itération: répète un processus
//pour un chiffre on prend tous les chiffres inférieur et on les multiplie ensemble
/*result :=1 //parce que sinon ça affiche uniquement 0
for i := nb; i > 0; i-- {
	result *= nb
}
return result*/

//arg = veut dire le nombre de facteurs pour arriver à x soit  24 dans l'exemple. Du coup
//on ne peut pas prendre result = 0 car 0*1*2*3 nous donne 0 donc on commence par 1 obligatoirement.

/*result := 1
if nb >= 25 || nb < 0 {// 25 car overflow car ce sera supérieur à 2^64 (à vérifier)
	result = 0

} else if nb > 0 {
	for i := 1; i <= nb; i++ {
		result *= i
	}
	} else if nb == 0 {//avec les modifications du haut, y aurait plus besoins de ça.
		result = 1
	} else {
		result = 0
}
return result*/

//* dans un int on peut avoir 64 chiffres maxi, à partir de 25 caractère, il va tronquer les chiffres et donner n'importe quoi
//il va y a voir un overflow. c'est un peu quand on anvoie un texte et qu'on a un nombre définit de charactère autorisé. Au final, il saturer et n'envoyé qu'une partie du texte
