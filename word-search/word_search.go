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
	foundTopBottom, err := scanTopBottom(words, puzzle)
	foundBottomTop, err := scanBottomTop(words, puzzle)
	foundTopLeftBottomRight, err := scanTopLeftBottomRight(words, puzzle)
	foundBottomRightTopLeft, err := scanBottomRightTopLeft(words, puzzle)
	foundBottomLeftTopRight, err := scanBottomLeftTopRight(words, puzzle)
	foundTopRightBottomLeft, err := scanTopRightBottomLeft(words, puzzle)
	for k, v := range foundLeftRight {
		found[k] = v
	}
	for k, v := range foundRightLeft {
		found[k] = v
	}
	for k, v := range foundTopBottom {
		found[k] = v
	}
	for k, v := range foundBottomTop {
		found[k] = v
	}
	for k, v := range foundTopLeftBottomRight {
		found[k] = v
	}
	for k, v := range foundBottomRightTopLeft {
		found[k] = v
	}
	for k, v := range foundBottomLeftTopRight {
		found[k] = v
	}
	for k, v := range foundTopRightBottomLeft {
		found[k] = v
	}
	if len(found) == 0 {
		return nil, errors.New("No words found in input puzzle")
	}
	if len(found) != len(words) {
		return nil, errors.New("Not all words are present in the puzzle")
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
		for i, col := range puzzle {
			j := strings.Index(col, word)
			if j != -1 {
				found[word] = [2][2]int{{i, j}, {i, j + len(word) - 1}}
			}
		}
	}
	return found, nil
}

func scanBottomTop(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	puzzle = reversePuzzle(TransposePuzzle(puzzle))
	for _, word := range words {
		for i, col := range puzzle {
			j := strings.Index(col, word)
			if j != -1 {
				// un-reverse j's index
				j = len(col) - j - 1
				found[word] = [2][2]int{{i, j}, {i, j - (len(word) - 1)}}
			}
		}
	}
	return found, nil
}

func scanTopLeftBottomRight(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	diagonals := PuzzleDiagonals(puzzle)
	for _, word := range words {
		for i, row := range diagonals {
			j := strings.Index(row, word)
			if j != -1 {
				xi, xj := GetDiagonalIndex(puzzle, i, j)
				yi, yj := GetDiagonalIndex(puzzle, i, j+(len(word)-1))
				found[word] = [2][2]int{{xi, xj}, {yi, yj}}
			}
		}
	}
	return found, nil
}

func scanBottomRightTopLeft(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	diagonals := PuzzleDiagonals(puzzle)
	for _, word := range words {
		for i, row := range diagonals {
			j := strings.Index(reverseString(row), word)
			if j != -1 {
				// un-reverse j's index
				j = len(row) - j - 1
				xi, xj := GetDiagonalIndex(puzzle, i, j)
				yi, yj := GetDiagonalIndex(puzzle, i, j-len(word)+1)
				found[word] = [2][2]int{{xj, xi}, {yj, yi}}
			}
		}
	}
	return found, nil
}

// Reverse puzzle, then reverse row to scan this direction.
func scanBottomLeftTopRight(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	diagonals := PuzzleDiagonals(reversePuzzle(puzzle))
	for _, word := range words {
		for i, row := range diagonals {
			j := strings.Index(reverseString(row), word)
			if j != -1 {
				// un-reverse j's index
				j = len(row) - j - 1
				xi, xj := GetDiagonalIndex(puzzle, i, j)
				xj = len(puzzle[xi]) - xj - 1
				yi, yj := GetDiagonalIndex(puzzle, i, j-len(word)+1)
				yj = len(puzzle[yi]) - yj - 1
				found[word] = [2][2]int{{xj, xi}, {yj, yi}}
			}
		}
	}
	return found, nil
}

// Reverse puzzle to scan this direction.
func scanTopRightBottomLeft(words, puzzle []string) (map[string][2][2]int, error) {
	found := map[string][2][2]int{}
	diagonals := PuzzleDiagonals(reversePuzzle(puzzle))
	for _, word := range words {
		for i, row := range diagonals {
			j := strings.Index(row, word)
			if j != -1 {
				xi, xj := GetDiagonalIndex(puzzle, i, j)
				xj = len(puzzle[xi]) - xj - 1
				yi, yj := GetDiagonalIndex(puzzle, i, j+len(word)-1)
				yj = len(puzzle[yi]) - yj - 1
				found[word] = [2][2]int{{xj, xi}, {yj, yi}}
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

// PuzzleDiagonals gets the diagonals of the puzzle, bottom left to top right in \ direction.
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
