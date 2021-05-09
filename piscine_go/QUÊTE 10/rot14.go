package piscine

func Rot14(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			continue
		}
		for j := 0; j < 14; j++ {
			if runes[i] == 'z' {
				runes[i] = 'a' - 1
			}
			if runes[i] == 'Z' {
				runes[i] = 'A' - 1
			}
			runes[i]++
		}
	}
	return string(runes)
}
