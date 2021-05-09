package student

func SplitWhiteSpaces(s string) []string {
	table := make([]string, size(s))
	//table est un tableau qui à son premier argument à 0
	var a string
	var i int
	for o, char := range s {
		if char != rune(32) { //espace avec ascii
			//On ajoute la char à la variable a
			a += string(char)
		} else {
			//Pour les espaces
			//On ajoute a au tableau à l'argument i
			table[i] = a
			//On réinitialise a
			a = ""
			//On ajoute 1 a i pour passer à l'argument suivant du tableau
			i++
		}
		//Pour faire apparaître le dernier argument
		//sinon il n'est pas prit en compte
		//On ajoute a au tableau
		if s[o] == s[len(s)-1] {
			table[i] = a
			//a = ""
		}
	}
	return table
}

//size va compter le mot et len prend seulemet chaque caractère individuellement
//En comptant le mot, il va fermet le crocher directement à la fin et non pas en rajoutant un espace

func size(s string) int {
	count := 0
	for j := range s {
		//Lorsque il y a un espace, une newline, un tab, ou que c'est le dernier caractère on ajoute 1 à count
		if s[j] == 32 || s[j] == 10 || s[j] == 9 || s[j] == s[len(s)-1] {
			count++
			//Count va nous permettre de connaitre le nombre de a
		}
	}
	return count
}
