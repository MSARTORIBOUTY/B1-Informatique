package piscine

func Max(a []int) int {

	args := a
	x := len(args)
	i := 0
	var count int
	j := 1

	for {
		i++
		count = 0

		if args[i] > args[i+1] {
			args[i], args[i+1] = args[i+1], args[i]
		}
		j = 0
		for j < x-1 {
			if args[j] < args[j+1] {
				count++
			}
			j++
		}

		if count == x-2 {
			break
		}

		if i == x-2 {
			i = 0
		}
	}
	return args[len(args)-1]
}
