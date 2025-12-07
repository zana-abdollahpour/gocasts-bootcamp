package leap

// IsLeapYear determines if a year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}

	return false
}

// IsLeapYear2 determines if a year is a leap year.
func IsLeapYear2(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return true
}
