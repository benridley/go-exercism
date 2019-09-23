package leap

// IsLeapYear return whether a given year is a leap year or not.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}
