package day03

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Function to parse all lines and print out all found numbers in each line.
func ParseNumbers(engine [][]byte) {
	for i, line := range engine {
		numbersInLine := []string{}
		numStr := ""

		for _, char := range string(line) {
			if unicode.IsDigit(char) {
				numStr += string(char)
			} else if numStr != "" {
				numbersInLine = append(numbersInLine, numStr)
				numStr = ""
			}
		}

		if numStr != "" {
			numbersInLine = append(numbersInLine, numStr)
		}

		fmt.Printf("Line %d: Numbers in this line: [%s]\n", i+1, strings.Join(numbersInLine, ", "))
	}
}

func Solve(part int, filePath string) {
	if part < 1 || part > 2 {
		fmt.Printf("Invalid part number: %v\n", part)
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	engine := [][]byte{}
	lines := string(content)
	linesSplit := strings.Split(lines, "\n")

	for _, line := range linesSplit {
		engine = append(engine, []byte(line))
	}

	ParseNumbers(engine)
}
