package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	d := *b

	*a = c / d
	*b = c % d
}

/*func UltimateDivMod(a *int, b *int) {
	var div int = *a
	*a = *a / *b
	*b = div % *b
}*/
