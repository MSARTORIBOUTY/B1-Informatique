package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

//suite fibonacci est une suite d'entier, chaque terme est la somme des deux thermes précédents. 3 est la somme de 2+1

/*if index < 0 { // si l'index est négative ça renvoit à -1 (marqué dans ytrak)
	return -1
} else if index == 0 {
	return 0
} else if index == 1 {
	return 1
} else {
	return Fibonacci(index-1) + Fibonacci(index-2)
}*/
