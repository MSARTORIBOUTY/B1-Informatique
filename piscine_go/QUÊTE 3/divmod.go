package piscine

func DivMod(a int, b int, div *int, mod *int) {
	//a & b sont deux variables
	//div et mod sont des pointeurs de type int (affiche chiffres)
	*div = a / b
	*mod = a % b
}
