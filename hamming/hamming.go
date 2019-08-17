/* 
Package hamming provides a method for calculating Hamming distance between two strings representing nucleotides.
*/
package hamming

import "errors"

// Distance returns the Hamming distance of two provided strings. The strings must be of equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Supplied strings must be of equal length to calculate Hamming distance.")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

