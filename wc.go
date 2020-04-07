package wc

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
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

// AnalyzeLine ...
func (stats *Stats) AnalyzeLine(line string) {
	stats.LineCount++
	stats.WordCount += CountWords(line)
	stats.ByteCount += len(line)
	stats.RuneCount += utf8.RuneCountInString(line)
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

// AnalyzeReader ...
func AnalyzeReader(reader io.Reader) (Stats, error) {
	stats := Stats{}
	bufReader := bufio.NewReader(reader)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return Stats{}, err
		}

		stats.AnalyzeLine(line)
	}

	return stats, nil
}
