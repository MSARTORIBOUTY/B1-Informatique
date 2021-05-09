package piscine

func IsLower(s string) bool {
	if len(s) > 0 {
		for _, i := range s {
			if !(i > 96 && i < 123) {
				return false
			}
		}
		return true
	} else {
		return true
	}
}
