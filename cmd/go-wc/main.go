package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	lineCount := 0
	wordCount := 0
	byteCount := 0
	runeCount := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		wordCountInLine := 0
		wasSpace := true
		for _, symbol := range line {
			if unicode.IsSpace(symbol) {
				if !wasSpace {
					wordCountInLine++
				}
				wasSpace = true
			} else {
				wasSpace = false
			}
		}

		lineCount++
		wordCount += wordCountInLine
		byteCount += len(line)
		runeCount += utf8.RuneCountInString(line)
	}

	fmt.Println(lineCount, wordCount, byteCount, runeCount)
}
