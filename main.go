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

		count := 0
		wasSpace := true
		for _, symbol := range line {
			if unicode.IsSpace(symbol) {
				if !wasSpace {
					count++
				}
				wasSpace = true
			} else {
				wasSpace = false
			}
		}

		fmt.Println(count)
	}
}
