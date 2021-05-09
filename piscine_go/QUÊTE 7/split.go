package piscine

func Split(s, sep string) []string {
	var a []string
	toFind := string(sep)
	x := len(s)
	y := len(toFind)
	j := 0

	for i := 0; i <= x-y; i++ {
		if toFind[:] == s[i:i+y] {
			a = append(a, s[j:i])
			j = i + y
		}
		if i == x-y {
			a = append(a, s[j:x])
		}
	}
	return a
}
