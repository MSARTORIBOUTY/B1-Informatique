package piscine

func StrLen(s string) (i int) { //S renvoit à "hello world". Du coup on rajout i pour après (on crée une variable i)
	for range s {
		i++ // ça rajoute 1 avec lecture gauche droite
	}
	return //fermeture loop mais à vérifier
}

/*package piscine

func StrLen(s string) int {
	count := 0    //on d'éclare count et on l'incrémante après dans le range s
	for range s { //range c'est une rune. Elle va lire chaques charactère et non pas octect par octect car elle a plus de mémoire. Elle a 4 octects
		count++
	}
	return count
	//return len([]rune(s))// à la place de return count. C'est une autre manière de faire, en créant un tableau de rune s

	//len(s) renvoit le nombre de bytes
	//For range parcourt rune par rune
}*/
