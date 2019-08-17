package raindrops

import "fmt"

// Convert prints a string based on the factors of the input. Pling, plang and plong if the factors contain 3, 5, and 7 respectively.
func Convert(x int) string {
	str := ""
	if x%3 == 0 {
		str += ("Pling")
	}

	if x%5 == 0 {
		str += ("Plang")
	}

	if x%7 == 0 {
		str += ("Plong")
	}

	if str == "" {
		return fmt.Sprintf("%d", x)
	}

	return str
}
