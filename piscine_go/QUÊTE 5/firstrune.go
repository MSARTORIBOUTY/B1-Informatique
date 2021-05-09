package piscine

func FirstRune(s string) rune { //s est définit comme un string et on veut qu'il devienne rune à la sortie
	str := []rune(s) //str renvoit à s au dessus, on veut le transformer s en rune d'où la suite
	return str[0]    //0 renvoit à la première lettre

}

//transformer un string en rune
