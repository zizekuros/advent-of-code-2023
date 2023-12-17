package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var sum int = 0

func Solve(part int, filePath string) {

	if part < 1 || part > 2 {
		fmt.Printf("Invalid part number: %v\n", part)
		os.Exit(1)
	}

	if part == 2 {
		fmt.Println("Part 2 not implemented yet")
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

	for i, line := range engine {
		numStr := ""
		numStartIndex, numEndIndex := -1, -1

		for j, char := range string(line) + "." {
			if unicode.IsDigit(char) {
				if numStartIndex == -1 {
					numStartIndex = j
				}
				numEndIndex = j
				numStr += string(char)
			} else if numStr != "" {

				addIt := false

				// check left and right
				if (numStartIndex > 0 && isSymbol(string(line[numStartIndex-1]))) || (numEndIndex < (len(line)-1) && isSymbol(string(line[numEndIndex+1]))) {
					addIt = true
				}

				// check up
				if i > 0 {
					for k := numStartIndex; k <= numEndIndex; k++ {
						if isSymbol(string(engine[i-1][k])) {
							addIt = true
							break
						}
					}
				}

				// check bottom
				if i < (len(engine) - 1) {
					for k := numStartIndex; k <= numEndIndex; k++ {
						if len(engine[i+1]) > k {
							if isSymbol(string(engine[i+1][k])) {
								addIt = true
								break
							}
						}
					}
				}

				// check up-left corner
				if i > 0 && numStartIndex > 0 {
					if isSymbol(string(engine[i-1][numStartIndex-1])) {
						addIt = true
					}
				}

				// check up-right corner
				if i > 0 && numEndIndex < (len(engine[i-1])-2) {
					if isSymbol(string(engine[i-1][numEndIndex+1])) {
						addIt = true
					}
				}

				// check bottom-left corner
				if i < (len(engine)-2) && numStartIndex > 0 && len(engine[i+1]) > (numStartIndex-1) {
					if isSymbol(string(engine[i+1][numStartIndex-1])) {
						addIt = true
					}
				}

				// check bottom-right corner
				if i < (len(engine)-2) && (numEndIndex+1) < len(engine[i+1]) {
					if isSymbol(string(engine[i+1][numEndIndex+1])) {
						addIt = true
					}
				}

				if addIt {
					addToSum(numStr)

				}
				numStr, numStartIndex, numEndIndex = "", -1, -1
			}
		}
	}

	fmt.Println(sum)
}

func isSymbol(ch string) bool {
	return !regexp.MustCompile(`^[0-9a-zA-Z\.]$`).MatchString(ch)
}

func addToSum(num string) {
	add, err := strconv.Atoi(num)
	if err == nil {
		sum += add
	}
}
