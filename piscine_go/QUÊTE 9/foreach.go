package piscine

func ForEach(f func(int), a []int) {
	for _, val := range a { //range va lire chaque valeur contenue dans le tableau a
		f(val) //On affiche chaque valeur du tableau a avec la fonction f. (On applique la fonction à chaques éléments du tableau)
	}
}
