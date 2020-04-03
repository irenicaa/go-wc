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
	count := 0
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

		countInLine := 0
		wasSpace := true
		for _, symbol := range line {
			if unicode.IsSpace(symbol) {
				if !wasSpace {
					countInLine++
				}
				wasSpace = true
			} else {
				wasSpace = false
			}
		}

		fmt.Println(countInLine)

		count += countInLine
	}

	fmt.Println(count)
}
