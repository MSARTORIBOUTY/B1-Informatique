package piscine

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if letter > 64 && 91 > letter || letter > 96 && 123 > letter {
			count++
		}
	}
	return count
}
