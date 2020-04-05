package wc

import (
	"fmt"
	"unicode"
)

// Stats ...
type Stats struct {
	LineCount int
	WordCount int
	ByteCount int
	RuneCount int
}

func (stats Stats) String() string {
	return fmt.Sprint(
		stats.LineCount,
		stats.WordCount,
		stats.ByteCount,
		stats.RuneCount,
	)
}

// CountWords ...
func CountWords(text string) int {
	wordCountInLine := 0
	wasSpace := true
	for _, symbol := range text {
		if unicode.IsSpace(symbol) {
			if !wasSpace {
				wordCountInLine++
			}
			wasSpace = true
		} else {
			wasSpace = false
		}
	}

	return wordCountInLine
}
