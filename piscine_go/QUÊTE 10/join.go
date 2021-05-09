package piscine

func Join(strs []string, sep string) (result string) {
	for i, val := range strs {
		result += val
		if i != len(strs)-1 {
			result += sep
		}
	}
	return
}
