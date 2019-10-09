package queenattack

import (
	"errors"
	"regexp"
)

// CanQueenAttack returns whether two queens on a given chess board are able to attack each other
func CanQueenAttack(w, b string) (bool, error) {
	var wLoc, bLoc [2]int
	var ok bool

	wLoc, ok = parseLocation(w)
	if !ok {
		return false, errors.New("Invalid location")
	}
	bLoc, ok = parseLocation(b)
	if !ok {
		return false, errors.New("Invalid location")
	}

	// Same square
	if wLoc[0] == bLoc[0] && wLoc[1] == bLoc[1] {
		return false, errors.New("The queens are on the same square")
	}

	// Same file
	if wLoc[0] == bLoc[0] {
		return true, nil
	}

	// Same rank
	if wLoc[1] == bLoc[1] {
		return true, nil
	}

	// Shared diagonal
	if abs(wLoc[0]-bLoc[0]) == abs(wLoc[1]-bLoc[1]) {
		return true, nil
	}

	return false, nil
}

func parseLocation(loc string) (out [2]int, ok bool) {
	validChessLocation := regexp.MustCompile("[a-h][1-8]")
	ok = validChessLocation.MatchString(loc)
	if !ok {
		return [2]int{0, 0}, ok
	}

	// Convert the alphabetical character to a location based on a = 1
	x := []rune(loc)[0]
	out[0] = int(x - ('a' - 1))
	// Convert the numerical character to an integer
	y := []rune(loc)[1]
	out[1] = int(y - ('1' - 1))
	return out, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
