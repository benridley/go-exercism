package luhn

import (
	"fmt"
	"strings"
)

// Valid returns whether a given string is valid based on the Luhn algorithm
func Valid(in string) bool {
	// Strip whitespace
	in = strings.Join(strings.Fields(in), "")

	// Catch empty/single digit strings here and return.
	if len(in) < 2 {
		return false
	}

	inRunes := []rune(in)

	for i := len(in) - 2; i >= 0; i -= 2 {
		doubled := int(inRunes[i]-'0') * 2
		if doubled > 9 {
			doubled -= 9
		}

		inRunes[i] = rune(doubled + '0')
	}
	fmt.Println(string(inRunes))
	sum := 0
	for _, n := range inRunes {
		// Check for non-digit characters and fail if any are found
		if n < 48 || n > 57 {
			return false
		}

		sum += int(n - '0')
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
