package cryptosquare

import (
	"log"
	"math"
	"regexp"
	"strings"
)

// Encode uses the 'crypto square' method to encode an input string.
func Encode(in string) string {
	// Normalize the input string by removing special characters and lowercasing the string
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	normalizedString := reg.ReplaceAllString(in, "")
	normalizedString = strings.ToLower(normalizedString)

	// Calculate the required cols/rows. Ideally we aim for a square, if not possible get the closest rectangle.
	sqrt := math.Sqrt(float64(len(normalizedString)))
	var rows, columns int
	if math.Mod(sqrt, math.Floor(sqrt)) == 0 {
		rows, columns = int(sqrt), int(sqrt)
	} else {
		rows, columns = int(math.Round(sqrt)), int(sqrt+1)
	}
	runes := []rune(normalizedString)

	// Pad the string to the end of the rectangle
	for len(runes) < (rows * columns) {
		runes = append(runes, ' ')
	}
	var encoded strings.Builder
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			encoded.WriteRune(runes[(j*(columns))+i])
		}
		encoded.WriteRune(' ')
	}
	encodedString := encoded.String()
	return encodedString[:len(encodedString)-1]
}
