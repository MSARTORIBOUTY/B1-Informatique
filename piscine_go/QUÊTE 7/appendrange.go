package piscine

func AppendRange(min, max int) []int {
	var answer []int

	for i := min; i < max; i++ {
		if i < max {
			answer = append(answer, i) // ajouter des grandeurs. Une sorte de min et max dans les intervalles.
		} else {
			return nil
		}
	}
	return answer
}
