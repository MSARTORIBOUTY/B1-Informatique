package piscine

func SortIntegerTable(table []int) {
	//for pour parcourir de gauche à droite
	//table[0] -> plus petit chiffre
	//table[1] -> plus grand ou égal que table[0]
	//table[x] -> plus petit ou égale que table[x+1]
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ { //si on met len(s)-1 ici, le 0 dans table ne sera pas prit en compte
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]

			}
		}

	}
	return //ce n'est pas obligatoire, car quand on regarde func, il n'y a pas de type de retour
}
