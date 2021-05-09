package piscine

func Compact(ptr *[]string) int {
	var tab []string

	for _, val := range *ptr {
		if val == "" {
		} else {
			tab = append(tab, val)
		}
	}
	*ptr = tab

	return len(tab)
}
