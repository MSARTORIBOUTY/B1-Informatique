package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

/*package piscine

func Swap(a *int, b *int) {
	//permuter l'affiche de a avec b
	//cette syntax est plus autorisé que l'autre
	temp := *b
	*b = *a
	*a = temp

}*/
