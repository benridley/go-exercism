package flatten

// Flatten unpacks nested lists into a single list without nil values.
func Flatten(arr interface{}) []interface{} {
	a := arr.([]interface{})
	flat := []interface{}{}
	for _, n := range a {
		if n == nil {
			continue
		} else if l, ok := n.([]interface{}); ok {
			flat = append(flat, Flatten(l)...)
		} else {
			flat = append(flat, n)
		}
	}
	return flat
}
