package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	args := a
	len := len(args)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if f(args[j], args[j+1]) == 1 {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
}
