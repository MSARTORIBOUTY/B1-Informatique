package main

import (
        "fmt"
        piscine ".."
)

func main() {
	fmt.Println(piscine.Sqrt(4))//ici 4 est le nbr que l'on va diviser. il va afficher 2 car on peut le diviser par 2. Si c'était 9, il afficherait 3. c'est des nbrs entiers!
	fmt.Println(piscine.Sqrt(3))//3 ne donne pas un reste entier donc il affichera 0. Cela vaut pour tous les autres comme 5 par exemple
}