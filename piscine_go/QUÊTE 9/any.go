package piscine

func Any(f func(string) bool, a []string) bool {
	//La fonction Any retourne vrai pour un tableau de string
	for _, val := range a { //On cherche les valeurs du tableau
		if f(val) { //Si le string passe par la fonction f, au moins un élément retourne vrai
			return true //Donc il retourne vrai
		}
	}
	return false //Autrement il retourne faux

}

/*package piscine

func Any(f func(string) bool, a []string) bool {
	result := false

	for _, val := range a {
		if f(val) == true {
			result = true
		}
	}
	return result
}*/
