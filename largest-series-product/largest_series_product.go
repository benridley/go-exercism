package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digitString string, span int) (int, error) {
	digits := make([]int, len(digitString))
	for i := 0; i < len(digitString); i++ {
		if digitString[i] < '0' || digitString[i] > '9' {
			return 0, errors.New("Invalid digit")
		}
		digits[i] = int(digitString[i] - '0')
	}

	if span > len(digits) || span < 0 {
		return 0, errors.New("Invalid span length")
	}
	// Calculate initial result from first series of digits
	currentProduct := 1
	maxProduct := 0
	for _, n := range digits[:span] {
		currentProduct *= n
	}
	maxProduct = currentProduct
	for i := span; i < len(digits); i++ {
		if digits[i-span] == 0 {
			currentProduct = sliceProduct(digits[(i-span)+1 : (i + 1)])
		} else {
			currentProduct = (currentProduct / digits[i-span]) * digits[i]
		}
		if currentProduct > maxProduct {
			maxProduct = currentProduct
		}
	}
	return maxProduct, nil
}

func sliceProduct(in []int) int {
	if len(in) == 0 {
		return 0
	}
	product := 1
	for _, n := range in {
		product *= n
	}
	return product
}
