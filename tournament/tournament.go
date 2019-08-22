package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ScoreCard contains a grop of season results.
type ScoreCard struct {
	Won  int
	Drew int
	Lost int
}

// TallyResult provides a ScoreCard for each team.
type TallyResult map[string]*ScoreCard

func (T TallyResult) addResultToScoreCard(firstTeam, secondTeam, result string) (TallyResult, error) {
	if _, ok := T[firstTeam]; !ok {
		T[firstTeam] = &ScoreCard{
			Won:  0,
			Drew: 0,
			Lost: 0,
		}
	}
	if _, ok := T[secondTeam]; !ok {
		T[firstTeam] = &ScoreCard{
			Won:  0,
			Drew: 0,
			Lost: 0,
		}
	}

	if result == "win" {
		T[firstTeam].Won++
		T[secondTeam].Lost++
	} else if result == "loss" {
		T[firstTeam].Lost++
		T[secondTeam].Won++
	} else if result == "draw" {
		T[firstTeam].Drew++
		T[secondTeam].Drew++
	} else {
		return T, errors.New("Invalid result: " + result)
	}
	return T, nil
}

// Tally returns tabulated match details from an input file
func Tally(reader *strings.Reader, buffer *bytes.Buffer) error {
	bufferString := ""
	tallyResult := TallyResult{}

	firstTeam := ""
	secondTeam := ""
	result := ""
	lineCount := 1

	for char, _, err := reader.ReadRune(); err != io.EOF; char, _, err = reader.ReadRune() {
		if err != nil {
			return fmt.Errorf("Failed to read from stream at line: %d", lineCount)
		}
		switch {
		case char == ';' && firstTeam == "":
			firstTeam = bufferString
			bufferString = ""
			continue

		case char == ';' && secondTeam == "":
			secondTeam = bufferString
			bufferString = ""
			continue

		case char == '\n' && result == "":
			result = bufferString
			bufferString = ""
			lineCount++

			tallyResult, err = tallyResult.addResultToScoreCard(firstTeam, secondTeam, result)
			firstTeam, secondTeam, result = "", "", ""
			if err != nil {
				return fmt.Errorf("Bad result at line %d: %s", lineCount, err.Error())
			}
			continue

		case char == '\n' && (firstTeam == "" || secondTeam == ""):
			return fmt.Errorf("Syntax error in input at line: %d", lineCount)

		default:
			bufferString += string(char)
		}
	}
	// Write header
	buffer.WriteString(fmt.Sprintf("%30s|%4s|%4s|%4s|%4s|%4s|\n", "Team", "MP", "W", "D", "L", "P"))
	for teamName, results := range tallyResult {
		totalPlayed := results.Won + results.Lost + results.Drew
		points := 3*results.Won + results.Drew
		summary := fmt.Sprintf("%30s|%4d|%4d|%4d|%4d|%4d|\n", teamName, totalPlayed, results.Won, results.Drew, results.Lost, points)
		buffer.WriteString(summary)
	}
	return nil
}
