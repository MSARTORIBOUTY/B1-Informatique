package piscine

func ConcatParams(args []string) string {

	result := "" // empty string pour ajouter autre arg

	for i := 0; i < len(args); i++ {
		if i < len(args)-1 { // til the last arg of args
			result = result + args[i] + "\n"
		} else {
			result = result + args[i]
		}
	}
	return result
}
