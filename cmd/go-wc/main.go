package main

import (
	"fmt"
	"log"
	"os"

	wc "github.com/irenicaa/go-wc"
)

func main() {
	stats, err := wc.AnalyzeReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stats)
}
