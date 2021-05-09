package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i, letter := range s {
		if letter > 96 && letter < 123 {
			str[i] = letter - 32
		} else {
			str[i] = letter //On ne touche pas ceux qui sont déjà en maj
		}
	}
	return string(str)
}
