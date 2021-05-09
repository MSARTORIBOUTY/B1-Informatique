package piscine

func Map(f func(int) bool, a []int) []bool {
	//On doit retourner un tableau de bool d'où le fait qu'on en crée un et on lui donne toutes les valeurs de a
	result := make([]bool, len(a)) // len(result)= len(a) avec result tableau type bool. Le premier est le type de tableau, le second la taille

	for i, val := range a {
		result[i] = f(val)
	}
	return result
}

//OU
/*var result []bool // result := []bool{} //donc on crée un tableau vide

for _,val := range a {
	result = append(result, f(val))
}
return val*/
