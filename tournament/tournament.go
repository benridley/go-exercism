package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

// ScoreCard represents the number of wins, losses, and draws a team has in a season.
type ScoreCard struct {
	Win, Loss, Draw int
}

// TeamPoints contains a team name and total number of points for easy sorting by points/alphabetical.
type TeamPoints struct {
	Name   string
	Points int
}

// TallyTable maps each team to its corresponding scorecard.
type TallyTable map[string]*ScoreCard

func (table TallyTable) addResultToTally(firstTeam, secondTeam, result string) (TallyTable, error) {
	// Initialize teams if they aren't already in the tally
	for _, team := range []string{firstTeam, secondTeam} {
		_, found := table[team]
		if !found {
			table[team] = &ScoreCard{
				Win:  0,
				Loss: 0,
				Draw: 0,
			}
		}
	}
	if firstTeam == secondTeam {
		return nil, fmt.Errorf("Cannot have two of the same teams in a match")
	}
	switch result {
	case "win":
		table[firstTeam].Win++
		table[secondTeam].Loss++
	case "loss":
		table[firstTeam].Loss++
		table[secondTeam].Win++
	case "draw":
		table[firstTeam].Draw++
		table[secondTeam].Draw++
	default:
		return nil, fmt.Errorf("Invalid match result: %s", result)
	}

	return table, nil
}

func (table TallyTable) toString() string {
	var tallyBuffer strings.Builder

	// Get team names sorted by score
	pointsList := []TeamPoints{}
	for teamName, scores := range table {
		p := TeamPoints{
			Name:   teamName,
			Points: (scores.Win*3 + scores.Draw),
		}
		pointsList = append(pointsList, p)
	}
	sort.Slice(pointsList, func(i, j int) bool {
		if pointsList[j].Points == pointsList[i].Points {
			return pointsList[i].Name < pointsList[j].Name
		}
		return pointsList[i].Points > pointsList[j].Points
	})

	tallyBuffer.WriteString(fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team"))
	for _, teamScores := range pointsList {
		scores := table[teamScores.Name]
		totalGames := scores.Draw + scores.Loss + scores.Win
		points := scores.Win*3 + scores.Draw
		tallyBuffer.WriteString(fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", teamScores.Name, totalGames, scores.Win, scores.Draw, scores.Loss, points))
	}
	return tallyBuffer.String()
}

// Tally returns tabulated match details from an input file
func Tally(reader io.Reader, buffer io.Writer) error {
	summary := TallyTable{}

	matchData, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	var firstTeam, secondTeam, result, strBuffer strings.Builder
	var seenChar, isCommented bool
	lineNumber := 0
	for _, ch := range matchData {
		switch {

		case (ch == '#'):
			isCommented = true

		case isCommented && ch == '\n':
			isCommented = false

		case isCommented:
			continue

		case (ch == '\n' || ch == ' ') && !seenChar:
			continue

		case (ch == ';' || ch == '\n') && strBuffer.Len() == 0:
			return fmt.Errorf("Missing value at line: %d", lineNumber)

		case ch == ';' && firstTeam.Len() == 0:
			firstTeam = strBuffer
			strBuffer.Reset()

		case ch == ';' && secondTeam.Len() == 0:
			secondTeam = strBuffer
			strBuffer.Reset()

		case ch == '\n' && result.Len() == 0 && firstTeam.Len() > 0 && secondTeam.Len() > 0:
			result = strBuffer
			summary, err = summary.addResultToTally(firstTeam.String(), secondTeam.String(), result.String())
			if err != nil {
				return fmt.Errorf("Error at line %d: %s", lineNumber, err.Error())
			}
			lineNumber++
			strBuffer.Reset()
			firstTeam.Reset()
			secondTeam.Reset()
			result.Reset()
			seenChar = false

		default:
			strBuffer.WriteByte(ch)
			seenChar = true
		}
	}

	if len(summary) == 0 {
		return fmt.Errorf("Input file contained no season data")
	}

	buffer.Write([]byte(summary.toString()))
	return nil
}
