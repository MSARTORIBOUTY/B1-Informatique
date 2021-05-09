package piscine

func IsUpper(s string) bool {
	if len(s) > 0 {
		for _, i := range s {
			if !(i > 64 && i < 91) {
				return false
			}
		}
		return true
	} else {
		return true
	}
}
