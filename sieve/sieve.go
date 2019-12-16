package sieve

// Sieve returns all primes up to n using the sieve of Eratosthenes
func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	sieve := make([]int, n+1)
	primes := make([]int, 0, n+1)

	// Generate numbers 0 through n
	for i := 0; i < len(sieve); i++ {
		sieve[i] = i
	}

	// Apply the sieve, marking numbers as 0 if they're a multiple
	for i := 2; i < len(sieve); i++ {
		if sieve[i] != 0 {
			for j := (i * 2); j < len(sieve); j += i {
				sieve[j] = 0
			}
		}
	}

	// Put all nonzero numbers into primes list
	for i := 2; i < len(sieve); i++ {
		if sieve[i] != 0 {
			primes = append(primes, sieve[i])
		}
	}

	return primes
}
