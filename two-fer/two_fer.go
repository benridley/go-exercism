/*
Package twofer provides a single method that returns a 'share with' string according to user input.
*/
package twofer

import "fmt"

// ShareWith returns a string "One for X, one for me" where X is either a supplied string, or 'you' when a string is absent.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
