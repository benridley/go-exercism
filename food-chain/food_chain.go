package foodchain

import (
	"fmt"
	"strings"
)

var songLines = map[string]string{
	"fly":    "",
	"spider": "It wriggled and jiggled and tickled inside her.\n",
	"bird":   "How absurd to swallow a bird!\n",
	"cat":    "Imagine that, to swallow a cat!\n",
	"dog":    "What a hog, to swallow a dog!\n",
	"goat":   "Just opened her throat and swallowed a goat!\n",
	"cow":    "I don't know how she swallowed a cow!\n",
	"horse":  "She's dead, of course!",
}

var order = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

// Verse generates a given verse of the Old Lady who Swallowed a Fly
func Verse(n int) string {
	var verse strings.Builder
	verse.WriteString("I know an old lady who swallowed a ")
	verse.WriteString(order[n-1])
	verse.WriteString(".\n")
	verse.WriteString(songLines[order[n-1]])
	if n == 8 {
		return verse.String()
	}
	for i := n - 1; i > 0; i-- {
		if i == 2 {
			verse.WriteString(fmt.Sprintf("She swallowed the %s to catch the spider that wriggled and jiggled and tickled inside her.\n", order[i]))
		} else {
			verse.WriteString(fmt.Sprintf("She swallowed the %s to catch the %s.\n", order[i], order[i-1]))
		}
	}
	verse.WriteString("I don't know why she swallowed the fly. Perhaps she'll die.")
	return verse.String()
}

func Verses(x, y int) string {
	var verses strings.Builder
	for i := x; i < y; i++ {
		verses.WriteString(Verse(i))
		verses.WriteString("\n\n")
	}
	verses.WriteString(Verse(y))
	return verses.String()
}

func Song() string {
	return Verses(1, 8)
}
