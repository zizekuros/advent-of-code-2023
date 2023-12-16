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

// Function to parse all lines and print out all found numbers in each line.
func ParseNumbers(engine [][]byte) {

	for i, line := range engine {
		numbersInLine := []string{}
		numStr := ""
		numStartIndex, numEndIndex := -1, -1

		for j, char := range string(line) {
			if unicode.IsDigit(char) {
				if numStartIndex == -1 {
					numStartIndex = j
				}
				numEndIndex = j
				numStr += string(char)
			} else if numStr != "" {
				numbersInLine = append(numbersInLine, numStr)

				addIt := false

				// check left and right
				if (numStartIndex > 0 && isSymbol(string(line[numStartIndex-1]))) || (numEndIndex < (len(line)-1) && isSymbol(string(line[numEndIndex+1]))) {
					fmt.Printf("adding left/right: %v\n", numStr)
					addIt = true
				}

				// check above
				if i > 0 {
					for k := numStartIndex; k < numEndIndex; k++ {
						if isSymbol(string(engine[i-1][k])) {
							fmt.Printf("adding up: %v\n", numStr)
							addIt = true
							break
						}
					}
				}

				// check under
				if i < (len(engine) - 1) {
					for k := numStartIndex; k < numEndIndex; k++ {
						if len(engine[i+1]) > k {
							if isSymbol(string(engine[i+1][k])) {
								fmt.Printf("adding up: %v\n", numStr)
								addIt = true
								break
							}
						}
					}
				}

				if addIt {
					addToSum(numStr)

				}
				numStr, numStartIndex, numEndIndex = "", -1, -1
			}
		}

		if numStr != "" {
			numbersInLine = append(numbersInLine, numStr)
		}
		if i < 332 {
			fmt.Printf("Line %d: Numbers in this line: [%s]\n", i+1, strings.Join(numbersInLine, ", "))
		}
	}

	fmt.Println(sum)
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

func isSymbol(ch string) bool {
	return !regexp.MustCompile(`^[0-9a-zA-Z\.]$`).MatchString(ch)
}

func addToSum(num string) {
	add, err := strconv.Atoi(num)
	if err == nil {
		sum += add
	}
}
