package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"

	wc "github.com/irenicaa/go-wc"
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

		lineCount++
		wordCount += wc.CountWords(line)
		byteCount += len(line)
		runeCount += utf8.RuneCountInString(line)
	}

	fmt.Println(lineCount, wordCount, byteCount, runeCount)
}
