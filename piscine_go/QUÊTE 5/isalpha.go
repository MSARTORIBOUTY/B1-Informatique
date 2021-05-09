package piscine

func IsAlpha(s string) bool {
	if len(s) > 0 {
		for _, i := range s {
			if !(i > 96 && i < 123) && !(i > 47 && i < 58) && !(i > 64 && i < 91) {
				return false
			}
		}
		return true
	} else {
		return true
	}
}
