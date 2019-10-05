package house

import "strings"

var nouns = []string{
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn"}

var verbs = []string{
	"that lay in ",
	"that ate ",
	"that killed ",
	"that worried ",
	"that tossed ",
	"that milked ",
	"that kissed ",
	"that married ",
	"that woke ",
	"that kept ",
	"that belonged to ",
}

// Song returns the entire House that Jack Built song.
func Song() string {
	verses := make([]string, len(nouns))
	for i := 0; i < len(nouns); i++ {
		verses[i] = Verse((i + 1))
	}
	return strings.Join(verses, "\n\n")
}

// Verse returns a specific verse from the House that Jack Built song.
func Verse(n int) string {
	verse := make([]string, n)
	for j := (n - 1); j >= 0; j-- {
		var line string
		if j == (n - 1) {
			line = "This is " + nouns[j]
		} else {
			line = verbs[j] + nouns[j]
		}
		verse[(n-1)-j] = line
	}
	verseString := strings.Join(verse, "\n")
	return verseString
}
