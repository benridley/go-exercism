package etl

import "strings"

// Transform takes an old format scrabble score table and returns a newly formatted one.
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for score, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}
	return out
}
