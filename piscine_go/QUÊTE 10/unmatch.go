package piscine

func Unmatch(a []int) int {
	for _, val := range a {
		i := 0
		for _, l := range a {
			if val == l {
				i++
			}
		}
		if i%2 == 1 {
			return val
		}
	}
	return -1
}
