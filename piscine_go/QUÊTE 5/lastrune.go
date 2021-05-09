package piscine

func LastRune(s string) rune {
	str := []rune(s)     //C'est un cast
	return str[len(s)-1] //-1 car on part de la fin
}

//affiche dernier rune d'un string
