package wordsearch

import (
	"errors"
	"strings"
)

// Solve solves a word square by identifying the location of input words in the given puzzle
func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	foundLeftRight, err := scanLeftRight(words, puzzle)
	foundRightLeft, err := scanRightLeft(words, puzzle)
	for k, v := range foundLeftRight {
		found[k] = v
	}
	for k, v := range foundRightLeft {
		found[k] = v
	}
	if len(found) == 0 {
		return nil, errors.New("No words found in input puzzle")
	}
	return found, err
}

func scanLeftRight(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	for _, word := range words {
		for i, row := range puzzle {
			j := strings.Index(row, word)
			if j != -1 {
				found[word] = [2][2]int{{j, i}, {j + len(word) - 1, i}}
			}
		}
	}
	return found, nil
}

func scanRightLeft(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	for _, word := range words {
		for i, row := range puzzle {
			j := strings.Index(reverseString(row), word)
			if j != -1 {
				// un-reverse j's index
				j = len(row) - j - 1
				found[word] = [2][2]int{{j, i}, {j - len(word) + 1, i}}
			}
		}
	}
	return found, nil
}

func scanTopBottom(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	for _, word := range words {
		for i, row := range puzzle {
			j := strings.Index(reverseString(row), word)
			if j != -1 {
				// un-reverse j's index
				j = len(row) - j - 1
				found[word] = [2][2]int{{j, i}, {j - len(word) + 1, i}}
			}
		}
	}
	return found, nil
}

func reversePuzzle(puzzle []string) []string {
	rev := make([]string, len(puzzle))
	for i, row := range puzzle {
		rev[i] = reverseString(row)
	}
	return rev
}

func transposePuzzle(puzzle []string) []string {
	tr := make([]string, len(puzzle[0]))
	for i, row := range puzzle {
		for j, col := range row {
			tr[j][i] = puzzle[i][j]
		}
	}
	return tr
}

func reverseString(s string) string {
	o := make([]rune, len(s))
	for i, c := range []rune(s) {
		o[len(o)-1-i] = c
	}
	return string(o)
}
