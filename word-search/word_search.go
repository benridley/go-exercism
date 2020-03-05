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
	foundTopBorrom, err := scanTopBottom(words, puzzle)
	for k, v := range foundLeftRight {
		found[k] = v
	}
	for k, v := range foundRightLeft {
		found[k] = v
	}
	for k, v := range foundTopBorrom {
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
	puzzle = TransposePuzzle(puzzle)
	for _, word := range words {
		for i, row := range puzzle {
			j := strings.Index(row, word)
			if j != -1 {
				found[word] = [2][2]int{{i, j}, {i - len(word) - 1, j}}
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

// TransposePuzzle converts columns of a puzzle to rows and vice versa
func TransposePuzzle(puzzle []string) []string {
	tr := make([][]byte, len(puzzle[0]))
	for i := range puzzle[0] {
		trRow := make([]byte, len(puzzle))
		tr[i] = trRow
	}
	for i, row := range puzzle {
		for j := range row {
			tr[j][i] = puzzle[i][j]
		}
	}
	trString := make([]string, len(puzzle[0]))
	for i := range tr {
		trString[i] = string(tr[i])
	}
	return trString
}

// PuzzleDiagonals gets the diagonals of the puzzle, bottom left to top right.
func PuzzleDiagonals(puzzle []string) []string {
	tr := make([][]byte, (len(puzzle) + len(puzzle[0]) - 1))
	// primary & lower diagonal
	for row := len(puzzle) - 1; row >= 0; row-- {
		diagLength := min((len(puzzle) - row), len(puzzle[0]))
		trRow := make([]byte, diagLength)
		for i := 0; i < diagLength; i++ {
			trRow[i] = puzzle[row+i][i]
		}
		tr[len(puzzle)-row-1] = trRow
	}
	// upper diagonal
	for col := 1; col < len(puzzle[0]); col++ {
		diagLength := min((len(puzzle[0]) - col), len(puzzle))
		trRow := make([]byte, diagLength)
		for i := 0; i < diagLength; i++ {
			trRow[i] = puzzle[i][col+i]
		}
		tr[len(puzzle)-1+col] = trRow
	}
	trString := make([]string, len(tr))
	for i := range tr {
		trString[i] = string(tr[i])
	}
	return trString
}

// GetDiagonalIndex translates a location found in a diagonal back to a location in the puzzle.
func GetDiagonalIndex(puzzle []string, i, j int) (int, int) {
	if i < (len(puzzle) - 1) {
		row := (len(puzzle) - 1) - i + j
		col := j
		return row, col
	}
	row := j
	col := i - (len(puzzle) - 1) + j
	return row, col
}

func reverseString(s string) string {
	o := make([]rune, len(s))
	for i, c := range []rune(s) {
		o[len(o)-1-i] = c
	}
	return string(o)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
