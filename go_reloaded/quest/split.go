package reloaded

func Split(s, sep string) []string {
	g := 0

	table := make([]string, Size(s, sep))

	for i := 0; i < len(s); i++ { //i prend la longueur de s
		for j := 0; j < len(sep); j++ { //j prend la longueur de la sep
			//i+j<len(s), c'est pour le premier mot+separateur
			//s[i+j] == sep[j], si la partie de la string s est strictement égale au séparateur
			if i+j < len(s) && s[i+j] == sep[j] {
				if j == len(sep)-1 { //Si j est égale à la longueur du séparateur
					i += len(sep) //Je pense que ça sert à supprimer le séparateur car
					//il ajoute pas le caractère mais la longueur à la place où y a le séparateur
					//len(sep) est considéré comme 1
					g++
				}
			}
		}
		table[g] += string(s[i]) // On ajoute chaque lettre de la string s si c'est différent du séparateur

	}
	return table //affiche le tableau
}
func Size(s, sep string) int {
	count := 0
	for j := range s {
		//Lorsque il y a un espace, une newline, un tab, ou que c'est le dernier caractère on ajoute 1 à count
		for k := range sep {
			if j == k || s[j] == s[len(s)-1] {
				count++
			}
		}
	}
	return count
}
