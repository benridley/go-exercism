package twelve

import (
	"fmt"
	"strings"
)

var numTable = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var items = map[int]string{
	1:  " a Partridge in a Pear Tree.",
	2:  " two Turtle Doves",
	3:  " three French Hens",
	4:  " four Calling Birds",
	5:  " five Gold Rings",
	6:  " six Geese-a-Laying",
	7:  " seven Swans-a-Swimming",
	8:  " eight Maids-a-Milking",
	9:  " nine Ladies Dancing",
	10: " ten Lords-a-Leaping",
	11: " eleven Pipers Piping",
	12: " twelve Drummers Drumming",
}

// Song returns the entire Twelve Days of Christmas Song
func Song() string {
	var lyrics strings.Builder
	for i := 1; i <= 12; i++ {
		lyrics.WriteString(Verse(i) + "\n")
	}
	return lyrics.String()[:len(lyrics.String())-1]
}

// Verse returns a particular 'day' of the Twelve Days of Christmas Song
func Verse(n int) string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me:", numTable[n]))

	for i := n; i > 1; i-- {
		str.WriteString(items[i] + ",")
	}

	if n > 1 {
		str.WriteString(" and")
	}

	str.WriteString(items[1])
	return str.String()
}
