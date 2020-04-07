package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

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

		stats.AnalyzeLine(line)
	}

	fmt.Println(stats)
}
