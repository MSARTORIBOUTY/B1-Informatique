package reloaded

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}
func NbrBase(nbr int, baseFrom string) string {
	indx := 0
	for _, res := range baseFrom {
		if string(res) == "-" || string(res) == "+" { //On ignore le signe
			indx = 1 //et on commence à l'index 1
			break
		}
	}
	if indx == 1 || len(baseFrom) < 2 {
		return "0"
	} else if 2147483647 < nbr || -2147483648 > nbr {
		return string(nbr)
	} else {
		i := 0
		smt := ""
		for nbr >= Len(baseFrom) {
			if nbr >= Len(baseFrom) {
				smt += string(baseFrom[nbr%Len(baseFrom)]) //On rajoute dans la string le reste de la division du nbr avec la base)
				nbr = nbr / Len(baseFrom)                  //Et le nbr prend la valeur de la vision du nbr avec la longueur de la base
				//BaseFrom, car le nbr vient de baseFrom
				i++
			}
		}
		smt += string(baseFrom[nbr])
		return Reverse(smt)
	}
}

//balaye la longueur
func Len(d string) int {
	len := 0
	for range d {
		len++
	}
	return len
}

//on remet les élémets à l'endroit dans le tableau
func Reverse(s string) string {
	var reverse string
	for i := Len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func AtoiBase(s string, base string) int {

	if !Base(base) {
		return 0
	}

	if !Number(s, base) {
		return 0
	}

	answer := 0
	lent := 0 //On balaye la longueur de la base
	for range base {
		lent++
	}

	for _, char := range s {
		for index2, char2 := range base {
			if char == char2 { //Si le char du nbr est équivalent au char contenu dans la base
				answer = answer*lent + index2 //Alors on multiplie la réponse avec la longueur de la base et on ajoute le caractère à l'index i
			}
		}
	}

	return answer
}

func Number(number, base string) bool {
	for _, n := range number {
		a := false
		for _, b := range base {
			if b == n { //Si le caractère de nbr et de la base son téquivalent alors c'est vrai
				a = true
				break
			}
		}

		if !a { //sinon c'est faux
			return false
		}
	}

	return true //Et on retourne la valeur si c'est vrai
}

//longueur de la base pour qu'elle soit valide
func Base(base string) bool {
	if StringLen(base) < 2 { //Si la len de la base est inf à 2 caractères
		return false
	}
	return Unique(base)
}

//La base doit avoir des caractères unique
func Unique(s string) bool {
	runes := []rune(s)
	for i, ch := range runes {
		for j := i + 1; j < StringLen(string(runes)); j++ {
			if ch == runes[j] && i != j { //Si on retrouve plusieurs fois un char dans la base et à un placement différent,
				return false //On retourne faux
			}
		}
	}
	return Signs(s)
}

//si la base comporte un signe
func Signs(s string) bool {

	for _, ch := range s {
		if ch == '+' || ch == '-' {
			return false
		}
	}

	return true
}

//remplace len, il balaye la longueur
func StringLen(str string) int {
	i := 0
	for range str {
		i++
	}
	return i

}
