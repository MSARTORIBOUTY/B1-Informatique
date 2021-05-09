package piscine

func CollatzCountdown(start int) int {
	result := start
	count := 0

	if start <= 0 {
		return -1
	}
	for {
		if result/2 != 0 {
			count++
			if result%2 == 0 {
				result /= 2
			} else {
				result *= 3
				result++
			}
			if result == 1 {
				break
			}
		} else {
			break
		}
	}
	return count
}
