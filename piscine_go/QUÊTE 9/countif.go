package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0

	for _, val := range tab {
		if f(val) { //La fonction retourne vrai. En gros on reprend le même système que any
			counter++ //Il compte chaques valeurs vrais
		}
	}

	return counter
}
