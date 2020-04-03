package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	lineCount := 0
	wordCount := 0
	symbolCount := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Print(line)

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

		fmt.Println(wordCountInLine)

		lineCount++
		wordCount += wordCountInLine
		symbolCount += len(line)
	}

	fmt.Println(lineCount, wordCount, symbolCount)
}
