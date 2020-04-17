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
