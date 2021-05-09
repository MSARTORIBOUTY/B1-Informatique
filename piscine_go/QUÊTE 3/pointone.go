package piscine

func PointOne(n *int) {
	//* avant le type ou juste avant n signifie que c'est un pointeur
	//pointeur: copie de l'adresse de la source sur laquelle on travaille
	*n = 1
}

/*Exemple pointeur avec la maison:
si on repeint les volets de la maison de quelqu'un, sans le pointeur on fait une copier de la maison pour repeindreles volets
mais, en soit la maison originelle n'est pas touché
Or, si on fait la copie de l'adresse source pour faire les volets, on modifie les volets de la maison originelle.
En gros, on garde en mémoire les modifications, d'une fonction à une autre
C'est sans doute pour ça, que le sudoku ne marchait pas, car le board n'était un pointeur
et par conséquents les modifications que faisaint chaque fonction les unes dans les autres n'étaient pas prise en compte
car on ne ravaillait que avec des copies et non pas la source originelle

D'où l'importance des pointeurs*/

//adresse de la variable se trouve dans la RAM. C'est un peu comme une adresse postale
//Quand on fait &n (esperluetten) ça dit que on va récupérer l'adresse de n
//On déclare un pointeur en mettant une * avant le type
//Avec plusieurs * ça fait un relais. C'est un pointeur de pointeur de pointeur,...
//Un tableau c'est un pointeur sur la première case
