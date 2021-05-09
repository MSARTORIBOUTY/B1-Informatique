package piscine

func IsNumeric(s string) bool {
	if len(s) > 0 {
		for _, i := range s {
			if !(i > 47 && i < 58) {
				return false
			}
		}
		return true
	} else {
		return true
	}
}
