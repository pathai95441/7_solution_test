package validate

func CheckConsecutiveLOrR(encoded string) bool {
	if len(encoded) >= 10 {
		if hasMoreThanTenConsecutive(encoded, 'L') || hasMoreThanTenConsecutive(encoded, 'R') {
			return true
		}
	}
	return false
}

func hasMoreThanTenConsecutive(encoded string, char rune) bool {
	count := 0
	for _, ch := range encoded {
		if ch == char {
			count++
			if count > 10 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
