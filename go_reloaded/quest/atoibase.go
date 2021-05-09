package reloaded

func AtoiBase(s string, base string) int {
	//La base > 2 caractères
	//chaques caractères de la base doivent être uniques
	//La base ne doit contenir aucun signe

	len1 := len(s)
	len2 := len(base)
	arr1 := []rune(s) //On met chaques éléements dans un tableau de manière à pouvoir les manipuler plus facilement
	arr2 := []rune(base)
	result := 0
	exists := false

	//si s équivaut à 0 et que la base contient moins de 2 caractères, on retourne 0 car non valide
	if len1 == 0 || len2 < 2 {
		return 0
	}
	for i := 0; i < len2; i++ { //On ajoute les caractères de la base dans un tableau pour ensuite tous les balayers
		if arr2[i] == '-' || arr2[i] == '+' { //Si la base contient un signe on retourne 0
			return 0
		}
		for j := i + 1; j < len2; j++ {
			if arr2[i] == arr2[j] { //Si deux caractères de la base sont identiques, on retourne 0
				return 0
			}
		}
	}
	if arr1[0] == '-' { //Si le premier index de s est un - on retourne 0. Il doit être strictement positif
		return 0
	}
	pow := RecursivePower(len2, len1-1) //S étant tiré de la base, on utilise la puissance
	for i := 0; i < len1; i++ {
		exists = false
		for j := 0; j < len2; j++ {
			if arr1[i] == arr2[j] { //Si la valeur de l'index de s est = à une valeur de l'index de la base
				exists = true           //Alors c'est vrai
				result = result + j*pow //Et on retourne le résultat en ajoutant la valeur de j * base^s
				pow /= len2
				j = len2
			}
		}
		if !exists {
			return 0
		}
	}
	return result //On retourne le résultat obtenue
}
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}

}

/*package student

func AtoiBase(s string, base string) int {

    if !Base(base) {
        return 0
    }

    if !Number(s, base) {
        return 0
    }

    answer := 0
    lent := 0
    for range base {
        lent++
    }

    for _, char := range s {
        for index2, char2 := range base {
            if char == char2 {
                answer = answer*lent + index2
            }
        }
    }

    return answer
}

func Number(number, base string) bool {
    for _, n := range number {
        a := false
        for _, b := range base {
            if b == n {
                a = true
                break
            }
        }

        if !a {
            return false
        }
    }

    return true
}

func Base(base string) bool {
    if StringLen(base) < 2 {
        return false
    }
    return Unique(base)
}

func Unique(s string) bool {
    runes := []rune(s)
    for i, ch := range runes {
        for j := i + 1; j < StringLen(string(runes)); j++ {
            if ch == runes[j] && i != j {
                return false
            }
        }
    }
    return Signs(s)
}

func Signs(s string) bool {

    for _, ch := range s {
        if ch == '+' || ch == '-' {
            return false
        }
    }

    return true
}
func StringLen(str string) int {
    i := 0
    for range str {
        i++
    }
    return i

}*/
