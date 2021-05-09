package main

import (
	"fmt"
	"io/ioutil" //pour lire les fichier ext. comme quest8
	"os"        //Boîte à outils des arguments
)

func main() {

	args := os.Args[1:] //chaque i dans le tableau équivaut à une ligne

	if len(args) > 1 {
		fmt.Print("Too many arguments")
		fmt.Print("\n")
	}

	if len(args) < 1 {
		fmt.Print("File name missing")
		fmt.Print("\n")
	}

	if len(args) == 1 {

		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
		}

		fmt.Printf("%s", content) // printf car on utilise %s pour afficher le contenu
	}
}

//cyprien
/*package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(file.Stat())

	arr := make([]byte, 14)//14 longueure de la phrase dans quest 8

	file.Read(arr)

	fmt.Println(arr)

	file.Close()
}*/
//Pour le lancer il faut faire go run main.go quest8
