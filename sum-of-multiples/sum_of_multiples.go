package summultiples

// SumMultiples calculates the sum of all multiples of given divisors up to a limit
func SumMultiples(limit int, divisors ...int) (sum int) {
	seen := make(map[int]bool)
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := 1; divisor*i < limit; i++ {
			n := divisor * i
			if seen[n] {
				continue
			}
			seen[n] = true
			sum += n
		}
	}
	return sum
}
