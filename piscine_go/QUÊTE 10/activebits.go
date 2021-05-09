package piscine

func ActiveBits(n int) (count int) {
	count = 0

	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return
}
