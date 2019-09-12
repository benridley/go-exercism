package sublist

import "reflect"

// Relation is an alias for a relationship between two lists
type Relation string

// Sublist returns a relationship between two lists, including superlist or equal lists
func Sublist(a, b []int) Relation {
	if reflect.DeepEqual(a, b) {
		return "equal"
	} else if isSublist(a, b) {
		return "sublist"
	} else if isSublist(b, a) {
		return "superlist"
	}
	return "unequal"
}

// Sublist strictly returns if a is a sublist of b
func isSublist(a, b []int) bool {
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(b); i++ {
		if b[i] == a[0] {
			for j := 0; j < len(a) && (i+j) < len(b); j++ {
				if b[i+j] != a[j] {
					break
				}
				if j == (len(a) - 1) {
					return true
				}
			}
		}
	}
	return false
}
