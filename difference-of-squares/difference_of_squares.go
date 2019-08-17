package diffsquares

// SumOfSquares returns the sum of the squares of the numbers from 0 to the input number
func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += (i * i)
	}
	return sum
}

// SquareOfSum returns the square of the sum of the numbers from 0 to the input number
func SquareOfSum(num int) int {
	sum := (num * (num + 1) / 2)
	return sum * sum
}

// Difference returns the difference of the sum of squares and the square of the sums from 0 up to the input number
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
