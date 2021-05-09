package main

import "fmt"

type node struct {
	right *node
	left  *node
	data  int
}

func insert(root *node, elem int) *node { //root = racine c'est le n°4 sur l'arbre
	//recursive mais un loop fonctionne aussi
	if root == nil { //root est vide, donc il n'y a pas d'arbre
		return &node{data: elem} //on retourne l'adresse d'un nouveau node avec le data d'un nouvel élément. On va chercher ailleur pour créer un arbre
	}
	if elem < root.data { //si l'élément est inférieur à la donné de root, donc si l'enfant est plus petit.
		root.left = insert(root.left, elem) //On va se diriger à gauche (systématiquement), et il est le résultat de l'enfant de gauche de l'élément.
	} else {
		root.right = insert(root.right, elem) //l'inverse
	}
	return root // On affiche root

}

//On va printer l'arbre
func PrintTree(root *node) {
	if root == nil { //si la racine est vide, il ne peut pas y avoir d'arbre
		return
	}
	fmt.Println(root.data)
	PrintTree(root.left)
	PrintTree(root.right)

}
func main() {
	var root *node
	root = insert(root, 4)
	root = insert(root, 2)
	root = insert(root, 6)
	root = insert(root, 3)
	root = insert(root, 5)
	root = insert(root, 7)
	root = insert(root, 1)
	PrintTree(root)
}
