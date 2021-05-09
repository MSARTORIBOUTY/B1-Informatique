package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0 // faux bool
	} else if nb == 1 {
		return 1 // vrai bool
	}
	i := 0
	for ; i <= nb/2; i++ { // Il faut que le nombre soit divisible par 2
		if i*i > nb {
			return 0
		} else if i*i == nb {
			return i
		}
	}
	return 0
}
