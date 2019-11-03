package pythagorean

import "math"

// Triplet represets a Pythagorean triplet where c ** 2 = a ** 2 + b ** 2
type Triplet [3]int

// Range returns all Pythagorean triplets with sides in the range min - max
func Range(min, max int) (triplets []Triplet) {
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			cSqr := (i * i) + (j * j)
			c := math.Sqrt(float64(cSqr))
			if math.Abs(c-math.Floor(c)) <= 0.000000001 && c <= float64(max) {
				triplets = append(triplets, Triplet{i, j, int(c)})
			}
		}
	}
	return triplets
}

// Sum returns all Pythagorean triplets where a + b + c = p
func Sum(p int) (triplets []Triplet) {
	for _, trip := range Range(1, p/2) {
		sumTrip := 0
		for _, side := range trip {
			sumTrip += side
		}
		if sumTrip == p {
			triplets = append(triplets, trip)
		}
	}
	return triplets
}
