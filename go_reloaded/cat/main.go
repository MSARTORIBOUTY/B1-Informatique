package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintText(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin) //Il va lire ce qu'il y a dedans et si il n'y a pas d'erreur, il retourne le data
		if err != nil {                       //Si il n'y a pas d'erreurs, err==nil
			PrintText(err.Error()) //car on a pas le droit au log.Fatal(err) donc o doit créer une fonction pour le print
			//Car PrintRune ou println ne supportent pas le type de err

		}
		PrintText(string(data)) //si il n'y a pas d'err alors on affiche le data à l'aide d'une autre fonction pour la même raison
	}
	for _, filename := range args {
		content, err := ioutil.ReadFile(filename) //ReadFile lit le contenue du fichier avec filename et retourne le contenue
		if err != nil {
			PrintText("ERROR: open " + filename + ": no such file or directory\n")
			//os.Exit(1)
		}
		//Pas d'errueur : err == nil
		PrintText(string(content)) //On print le contenu
	}
}
