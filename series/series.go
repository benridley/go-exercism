package series

// All returns all contigious substrings from a string of a given length
func All(n int, s string) (subs []string) {
	for i := 0; i <= len(s)-n; i++ {
		subs = append(subs, s[i:(i+n)])
	}
	return subs
}

// UnsafeFirst returns the first substring of length n in a string
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
