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

func TestCountWords_withSpacesAtTheEnd(test *testing.T) {
	count := CountWords("one two three ")
	assert.Equal(test, 3, count)
}
