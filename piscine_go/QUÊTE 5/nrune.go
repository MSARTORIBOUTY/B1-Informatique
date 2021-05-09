package piscine

func NRune(s string, n int) rune {
	str := []rune(s)          //On renvoit n dans le string
	if n > 0 && n <= len(s) { //n est compris entre 0 et la longueur de s
		return str[n-1]
	} else {
		return 0
	}
}
