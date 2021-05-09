package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Valid(a int, op rune, b int) {

	/*if a >= 9223372036854775807 { //Pour que l'overflow print 0
		z01.PrintRune('0')
		return //Pour arrêter l'algo
	}*/
	/*if op != '+' || op != '-' || op != '/' || op != '*' || op != '%' {
		z01.PrintRune('0')
		return
	}*/

	//On détermine la règle des calculs. Quel signe fait quoi
	/*switch os.Args[2] {
	case "+":
		result := a + b
		print(result)
	case "-":
		result := a - b
		print(result)
	case "*":
		result := a * b
		print(result)
	case "/":
		result := a / b
		print(result)
	case "%":
		result := a % b
		print(result)
	}*/
	result := 0
	if op == '+' {
		result = a + b
		print(result)
	} else if op == '-' {
		result = a - b
		print(result)
	} else if op == '*' {
		result = a * b
		print(result)
	} else if op == '/' {
		result = a / b
		print(result)
	} else if op == '%' {
		result = a % b
		print(result)
	} else {
		z01.PrintRune('0')
		return
	}
	/*for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == "+" && b[i] == "+"{
				if result == "-" {
					z01.PrintRune('0')
					return
				}
			} else if a[i] == "-" && b[i] == "-" {
				if result == "+" {
					z01.PrintRune('0')
					return
				}
			}
		}
	}*/

}
func InvalidNbr(a string, op rune, b string) {
	DivError := "No division by 0"
	ModError := "No modulo by 0"
	table := []rune(a) //On met la valeur de a dans un tableau de rune pour qu'il comprenne que c'est l'alphabet que l'on veut sinon ça ne marche pas
	table2 := []rune(b)

	if b == "0" {
		if op == '/' {
			print(DivError)
			return //On retourne une string qui va afficher le message désiré
		} else if op == '%' {
			print(ModError)
			return
		}
	}

	//Si a ou b sont équivalent à un mot. Leur type est définit en string dans la func sinon ça marche pas
	for i := 0; i < len(a); i++ {
		if table[i] != '-' && (table[i] >= 58 && table[i] <= 127 || table[i] >= 0 && table[i] <= 47) { // ATTENTION!! Bien dire que le permier index doit être différent de - sinon tous les calculs avce un - marqueront 0
			z01.PrintRune('0')
			return //Pour arrêter l'algorithme sinon il allait dans la func valid
		}
	}
	for i := 0; i < len(b); i++ {
		if table2[i] != '-' && (table2[i] >= 58 && table2[i] <= 127 || table2[i] >= 0 && table2[i] <= 47) {
			z01.PrintRune('0')
			return
		}
	}
	Valid(StringToInt(a), op, StringToInt(b)) //On appel la fonction valide si il n'y a aucun erreur
	return                                    //Un retour car j'ai mis une string enretour dans la func
}
func StringToInt(s string) int {

	copie := s
	if copie[0] == '-' { //Si le chiffre commence par -
		copie = s[1:] //Alors la copie commence par le premier chiffre donc l'index 1
	}
	//fonction result converti s qui est une string en byte
	result := 0
	for _, xrune := range copie { //On créer une variable xrune qui parcourt s
		result = result*10 + int(xrune-48) //fonction charToInt avec xrune que l'on converti en byte
		//On multiplie par 10 pour pouvoir décaler les dizaines
	}
	if s[0] == '-' { //si s commence par un - alors on multiplie le résultat par -1 pour le rendre positif et faire le calcul.
		result *= -1
	}
	return result

}
func main() {
	//3 arguments: une valeur, un opérateur et une autre valeur
	//utilisation int64
	//opérateur valide: + - * / %
	//invalide opérateur --> 0
	//invalide nombre d'arg --> print rien du tout
	//%0 et /0 print "No modulo by 0" v
	a := os.Args[1]
	b := os.Args[3]
	var op rune

	for _, char := range os.Args[2] {
		op = char
	}
	InvalidNbr(a, op, b) //On appel la func nbr invalid pour toutes erreurs
}
