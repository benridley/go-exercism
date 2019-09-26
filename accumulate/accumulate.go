package accumulate

// Accumulate applies a function to each element of a collection.
func Accumulate(n []string, fn func(string) string) []string {
	accumulated := make([]string, len(n))
	for i, val := range n {
		accumulated[i] = fn(val)
	}
	return accumulated
}
