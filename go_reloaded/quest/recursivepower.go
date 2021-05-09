package reloaded

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power-1) //-1 évite l'overflow, il permet de partir à l'enver
		//recursivepower appel la puissance du nb

	}
}
