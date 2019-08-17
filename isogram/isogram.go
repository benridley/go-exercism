package isogram

import "strings"

// IsIsogram returns wheterh a string contains no repeated letters, excluding spaces and hyphens.
func IsIsogram(in string) bool {
	letters := make(map[rune]bool)
	in = strings.ToLower(in)
	for _, letter := range in {
		if letter == ' ' || letter == '-' {
			continue
		}
		if letters[letter] == true {
			return false
		}
		letters[letter] = true
	}
	return true
}
