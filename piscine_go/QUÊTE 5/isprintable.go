package piscine

func IsPrintable(s string) bool {
	if len(s) > 0 {
		for _, i := range s {
			if i > 0 && i < 32 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
