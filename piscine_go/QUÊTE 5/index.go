package piscine

func Index(s string, toFind string) int {
	x := len(s)                 //Longueur de s
	y := len(toFind)            //Longueur de toFind
	for i := 0; i <= x-y; i++ { //On soustrait la longueur de s avec celle de toFind
		if toFind == s[i:i+y] {
			return i
		}
	}
	return -1 // celui qu'on veut afficher à la fin
}
