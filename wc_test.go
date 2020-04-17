package wc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords_withOneWord(test *testing.T) {
	count := CountWords("one")
	assert.Equal(test, 1, count)
}

func TestCountWords_withFewWords(test *testing.T) {
	count := CountWords("one two three")
	assert.Equal(test, 3, count)
}

func TestCountWords_withSpacesAtTheStart(test *testing.T) {
	count := CountWords("   one two three")
	assert.Equal(test, 3, count)
}

func TestCountWords_withSpacesAtTheEnd(test *testing.T) {
	count := CountWords("one two three   ")
	assert.Equal(test, 3, count)
}

func TestCountWords_empty(test *testing.T) {
	count := CountWords("")
	assert.Equal(test, 0, count)
}

func TestCountWords_withFewSpaces(test *testing.T) {
	count := CountWords("   ")
	assert.Equal(test, 0, count)
}

func TestCountWords_withDoubleSpaces(test *testing.T) {
	count := CountWords("one  two  three")
	assert.Equal(test, 3, count)
}

func TestCountWords_withTabs(test *testing.T) {
	count := CountWords("one\ttwo\tthree")
	assert.Equal(test, 3, count)
}

func TestCountWords_withUnicode(test *testing.T) {
	count := CountWords("один два три")
	assert.Equal(test, 3, count)
}

func TestStats_AnalyzeLine(test *testing.T) {
	stats := Stats{LineCount: 1, WordCount: 2, ByteCount: 3, RuneCount: 4}
	stats.AnalyzeLine("one two three")

	wantStats := Stats{LineCount: 2, WordCount: 5, ByteCount: 16, RuneCount: 17}
	assert.Equal(test, wantStats, stats)
}

func TestStats_AnalyzeLine_withUnicode(test *testing.T) {
	stats := Stats{LineCount: 1, WordCount: 2, ByteCount: 3, RuneCount: 4}
	stats.AnalyzeLine("один два три")

	wantStats := Stats{LineCount: 2, WordCount: 5, ByteCount: 25, RuneCount: 16}
	assert.Equal(test, wantStats, stats)
}
