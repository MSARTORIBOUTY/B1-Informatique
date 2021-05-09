package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintString(s string) { //Parce que l'on a pas le droit au fmt, donc on le place transforme en rune
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 { //Si le programme n'a pas d'args
		data, err := ioutil.ReadAll(os.Stdin) //il prend le standar input, il prend l'information écrit et le fait passer par le programme pour le ressortir
		if err != nil {
			PrintString(err.Error())
		}
		PrintString(string(data))

	}

	//On balaye les arguments
	for _, filename := range args {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			PrintString("ERROR: open " + filename + ": no such file or directory\n")
			os.Exit(1)
		}
		PrintString(string(content))

	}
}

//LE MÊME UN PEU DIFFERENT AVEC DES ANOTATIONS
/*package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintString(s string) {
	random := []rune(s)
	for _, char := range random {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:] //On prend le premier au dernier argument, y compris le chemin de la provenance. Quand on fait un go run "nom fichier" puis abc par exemple, il va prendre l'argument abc. Mais on exclu le premier arg car il sert à rien

	//on va parcourir tous les args, donc tous les fichiers que l'on veut print en même temps
	for _, Filename := range args {
		//Filename := args[0] //car sinon erreur, on l'ajoute au tout premier index

		content, err := ioutil.ReadFile(Filename) //ioutil lire, afficher et manipuler des fichiers; () le nom du fichier que l'on veut lire
		if err != nil {
			PrintString("ERROR: open " + Filename +
				": no such file or directory")
			os.Exit(1) //On quitte avec un message d'erreur. 1 signifie erreur
			//log.Fatal(err)//Il retourne le programme
		}
		PrintString(string(content)) //Sinon juste fmt.Println(content) mais ça va afficher les charactères en ascii. Le %s le traite comme une string et l'affiche tel quel. Il faut qu'on le cast en string: fmt.Println(string(content))

		//pour tous les lires et pas se limiter à un seul
		if len(args) == 0 {
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				PrintString(err.Error())
			}
			fmt.Println(string(data))
		}
	}

}*/

//CELUI DE TRISTAN

//Cat abréviation de concaténer, mettre bout à bout
/*i := 0
for range args*/ // Ou faire for i := range args, il va lire le tableau du début à la fin avec i comme index. Il remplace le for avec initialisation, condition et incrémentation

/*package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func put(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
func putln(s string) {
	put(s)
	z01.PrintRune('\n')
}

func fatal(s string) {
	putln(s)
	os.Exit(1)
}

func main() {
	i := 0
	for range os.Args {
		i++
	}
	if i == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fatal(err.Error())
		}
		put(string(b))
	} else {
		for _, file := range os.Args[1:] {
			b, err := ioutil.ReadFile(file)
			if err != nil {
				fatal("ERROR: " + err.Error())
			}
			put(string(b))
		}
	}
}*/
