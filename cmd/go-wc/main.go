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
	stats := wc.Stats{}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		stats.LineCount++
		stats.WordCount += wc.CountWords(line)
		stats.ByteCount += len(line)
		stats.RuneCount += utf8.RuneCountInString(line)
	}

	fmt.Println(stats.LineCount, stats.WordCount, stats.ByteCount, stats.RuneCount)
}
