package pangram

import "unicode"

// IsPangram returns true if the input string includes all characters in the alphabet
func IsPangram(in string) bool {
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alphabetCounter := make(map[rune]int)
	for _, e := range alphabet {
		alphabetCounter[e] = 0
	}

	for _, e := range in {
		alphabetCounter[unicode.ToLower(e)]++
	}

	for _, n := range alphabetCounter {
		if n == 0 {
			return false
		}
	}
	return true
}
