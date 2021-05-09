package reloaded

/*func Sep(s, sep string) []string {
	Space := []string{}
	empty := ""

	for range s {
		if sep == s {
			Space = append(Space, empty)
		}
	}
	return Space

}*/

func Split(s, sep string) []string {
	arr := make([]string, size(s))
	toFind := string(sep) //si on ne fait pas ça, pb de slice pour la suite
	x := len(s)
	y := len(toFind)
	//j := 0
	str := ""

	for _, char := range s {
		for i := 0; i <= x-y; i++ { //i est compris entre 0 et la string- sep compris.
			//càd qu'il prendra les lettres une à une jusqu'à tomber sur le séparateur qu'il va enlever (d'où le x-y: len(s)-len(y))
			if toFind[:] == s[i:i+y] { //On crée toFind car y ne peux pas comporter de [] et sans les [] on a une erreur de slice donc c'est important
				//Si sep == au mot + sep
				str += string(char)
				//arr = append(arr, s[j:i]) //On lui rajoute un espace (j = 0 est vue comme un espace) puis le mot
				//j = i + y                 //! cette étape est importante! c'est pour passer au mot suivant, sinon il affichera hello, un espace, helloHAhow,...)
				//C'est une réinitialisation de la valeur de j
			}
			//Sans cette étape, on ne possède pas "you?" car au final, après le dernier mot il n'y a pas de séparateur
			if i == x-y { // Si i est stricement égale à len(s)-len(sep) donc si c'est le mot sans séparateur
				arr[i] = str
				str = ""
				//arr = append(arr, s[j:x]) //On rajoute un espace et le mot
			}
		}
	}

	return arr //Afficher le tableau

}
func size(s string) int {
	count := 0
	for j := range s {
		//Lorsque il y a un espace, une newline, un tab, ou que c'est le dernier caractère on ajoute 1 à count
		if s[j] == 32 || s[j] == 10 || s[j] == 9 || s[j] == s[len(s)-1] {
			count++
			//Count va nous permettre de connaitre le nombre de mot
		}
	}
	return count
}

/*func Split(s, sep string) []string {
	table := []string{}
	lenSep := string(sep)
	x := len(s)
	y := len(lenSep)
	j := 0

	for i := 0; i <= x-y; i++ {
		if lenSep[:] == s[i:i+y] {
			table = append(table, s[j:i])
			j = i + y
		}
		if i == x-y {
			table = append(table, s[j:x])
		}
	}
	return table

}*/

/*func Split(s, sep string) []string {
	table := []string{}
	lettre := ""                           //On crée une string vide pour faire les mots
	sep = strings.ReplaceAll(s, "HA", " ") //On remplace chaque HA en espace dans s

	for i, char := range sep {
		if char != ' ' || char != '\t' || char != '\n' {
			lettre += string(sep[i]) //Quand c'est une lettre, on l'ajoute l'une après l'autre. Si c'est vide ce sera un espace. 9a affiche la phrase désirée

		} else if lettre != "" { //Si c'est l'inverse de char
			table = append(table, lettre) //On l'ajoute au tableau
			lettre = ""                   //Et one le réinitialise pour former un nouveau mot
		}
	}
	table = append(table, lettre) //On ne peut pas retourner char car il n'existe que dans la boucle for. On rajoutte les valeur juste avan le dernier return sinon il nous retourne deux fois le tableau
	return table
}*/
