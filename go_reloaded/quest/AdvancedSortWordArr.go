package reloaded

func AdvancedSortWordArr(arr []string, f func(a, b string) int) {
	size := 0
	for range arr {
		size++
	}
	var temp string
	//Deux var i et j pour comparer les chiffres
	//On compare i au premier index avec j au deuxième
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if f(arr[j], arr[j+1]) > 0 {
				temp = arr[j]//On inverse les valeurs en fonction de leur ordre d'apparition
				arr[j] = arr[j+1]//dans le tableau ascii
				arr[j+1] = temp
			}
		}
	}
}

func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return 0
}


