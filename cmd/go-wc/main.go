package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	wc "github.com/irenicaa/go-wc"
)

func main() {
	showLines := flag.Bool("l", false, "show line count")
	showWords := flag.Bool("w", false, "show word count")
	showBytes := flag.Bool("c", false, "show byte count")
	showSymbols := flag.Bool("m", false, "show symbol count")
	flag.Parse()

	stats, err := wc.AnalyzeReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	if *showLines {
		fmt.Printf(" %6d", stats.LineCount)
	}
	if *showWords {
		fmt.Printf(" %6d", stats.WordCount)
	}
	if *showBytes {
		fmt.Printf(" %6d", stats.ByteCount)
	}
	if *showSymbols {
		fmt.Printf(" %6d", stats.RuneCount)
	}
	fmt.Println()
}
