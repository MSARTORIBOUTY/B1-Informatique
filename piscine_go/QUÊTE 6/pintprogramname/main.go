/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args                //os c'est une bibliothèque
	programname := []rune(argument[0]) //on déclare la fonction comme étant une rune

	for i := range programname { // ça fait la liste ducon! ça liste tout le fichier
		z01.PrintRune(programname[i]) //Tu affiche le programme dans i
	}
}
*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[0]
	slashindex := -1
	for i, r := range s {
		if r == '/' {
			slashindex = i
		}
	}
	for _, letter := range s[slashindex+1:] {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

/* s := []rune(os.Args[0])//index 0 is the name of the folder

for _, letter := range s {
	if letter != '.' && letter != '/' {//on enlève les signes pour les fichiers exécutable
		z01.PrintRune(letter)
	}
}
z01.PrintRune('\n')