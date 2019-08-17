package grains

import "errors"

// Square returns the number of grains on a given 'square', if the first square has a single grain
// and each square that follows has double the previous square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid value")
	}

	return 1 << uint64(n-1), nil
}

// Total returns the sum of the series 1 + 1^2 + 1^2... + 1^63 which is known ahead of time.
func Total() uint64 {
	return 18446744073709551615
}
