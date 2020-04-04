package wc

import "unicode"

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
