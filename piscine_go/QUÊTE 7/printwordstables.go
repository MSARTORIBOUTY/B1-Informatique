package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(args []string) {
	var result string // empty string pour ajouter autre arg
	i := 0            // on déclare une variable qui renverra dans la longueur de len(args)

	for _, char := range args { //on doit balayer la longueur de args
		result += char //On ajoute le charactère au result (qui est vide soit dit en passant), du coup ça va afficher le charactère
		if i == len(args) {
		} else {
			result += "\n" //On affiche un retour à la ligne si il n'y a pas de charactère
		}
		i++ //on l'incrémente vue que on l'a utiliser pour i==
	}
	result += "\n" //on doit faire un retour à la ligne à la toute fin

	for _, char := range result { //On l'utilise pour pouvoir afficher la rune. Attention ce doit être un string
		z01.PrintRune(rune(char))
	}
}

/*
	var j int

	for i := 0; i < len(args); i++ {
		result = result + args[i] + "\n"
		j = i
	}
	fmt.Println(j)*/
