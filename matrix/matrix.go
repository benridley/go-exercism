package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a two dimensional array of integers.
type Matrix [][]int

// New returns a matrix createf from the input string.
func New(in string) (Matrix, error) {
	// Splitting on newlines provides the rows.
	rowStrings := strings.Split(in, "\n")

	// Convert the rows to numerical values.
	rows := make([][]int, len(rowStrings))
	for i, rowString := range rowStrings {

		// Splitting on spaces yields invidivual values in the row, e.g. "1 2 3" becomes ["1", "2", "3"]
		valueStrings := strings.Split(string(strings.Trim(rowString, " ")), " ")
		values := make([]int, len(valueStrings))

		for j, valString := range valueStrings {
			val, err := strconv.Atoi(valString)
			if err != nil {
				return nil, fmt.Errorf("Could not convert value %s into numeric value", valString)
			}
			values[j] = val
		}
		rows[i] = values
	}

	// Check rows have equal length
	for _, row := range rows {
		if len(row) != len(rows[0]) {
			return nil, errors.New("Rows must be of equal length")
		}
	}
	return rows, nil
}

// Rows returns an array of rows from the given matrix.
func (rows Matrix) Rows() [][]int {
	newRows := make([][]int, len(rows))
	for i, row := range rows {
		newRow := make([]int, len(row))
		copy(newRow, row)
		newRows[i] = newRow
	}
	return newRows
}

// Cols returns an array of columns from the given matrix
func (rows Matrix) Cols() [][]int {
	columns := make([][]int, len(rows[0]))
	for i := 0; i < len(rows[0]); i++ {
		columns[i] = make([]int, len(rows))
		for j := 0; j < len(rows); j++ {
			columns[i][j] = rows[j][i]
		}
	}
	return columns
}

// Set a matrix to the values provided by string
func (rows Matrix) Set(row, col, val int) bool {
	rows[row][col] = val
	return true
}
