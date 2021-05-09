package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var a []int //a devient le tableau, on crée un tableau vide pour y mettre les valeurs
	valeur := n //n prend les valeurs 1, 2, 3
	count := 0  //nbr de chiffre qu'il y a dans n

	for n != 0 { //n ne prend pas la valeur 0
		valeur = n % 10 //pour récupérer le dernier caractère (on prend les dizaine seulement)
		n /= 10         // n divisé par 10 (donc si on divise 321, on garde 32 et ainsi de suite)
		a = append(a, valeur)
		count++
	} //pour trier:
	for i, j := 0, 1; j < count; i, j = i+1, j+1 { //i renvoit au premier chiffre, j au deuxième (j doit être inférieur au nombre de chiffre qu'il y a dans n, note: j corresponde à l'emplacement d'où l'importance de le mettre inf.)
		if a[i] > a[j] { //dans le tableau a, i > j car 3 > 2
			a[i], a[j] = a[j], a[i] //vue qu'on le met dans l'ordre, 32 ==> 23
			i = -1                  //On renvoit au for 0-1=-1, le calcul montre l'inverse en fait (on réinitialise)
			j = 0                   // 1-1=0
		}
	}
	for t := 0; t < count; t++ { //Il renvoit au n == 0. Puisqu'il n'y a que 0, il n'y a rien à trier, donc on l'affiche simplement
		z01.PrintRune(rune(a[t] + 48)) //On ajoute 0 à la variable t dans le tableau a (on met 48 pour le décalage car sur le tableau les chiffres commencent à 48 (0))
	}
}
