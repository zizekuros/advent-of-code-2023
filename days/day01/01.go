package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve(part int, filePath string) {

	if part < 1 || part > 2 {
		fmt.Printf("Invalid part number: %v\n", part)
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if part == 2 {
			line = strings.ReplaceAll(line, "oneight", "18")
			line = strings.ReplaceAll(line, "twone", "21")
			line = strings.ReplaceAll(line, "threeight", "38")
			line = strings.ReplaceAll(line, "fiveight", "58")
			line = strings.ReplaceAll(line, "nineight", "98")
			line = strings.ReplaceAll(line, "sevenine", "79")
			line = strings.ReplaceAll(line, "eightwo", "82")
			line = strings.ReplaceAll(line, "eightree", "83")
			line = strings.ReplaceAll(line, "nineight", "98")
			line = strings.ReplaceAll(line, "one", "1")
			line = strings.ReplaceAll(line, "two", "2")
			line = strings.ReplaceAll(line, "three", "3")
			line = strings.ReplaceAll(line, "four", "4")
			line = strings.ReplaceAll(line, "five", "5")
			line = strings.ReplaceAll(line, "six", "6")
			line = strings.ReplaceAll(line, "seven", "7")
			line = strings.ReplaceAll(line, "eight", "8")
			line = strings.ReplaceAll(line, "nine", "9")
		}
		sum += calibrationValue(line)
	}

	fmt.Printf("%v", sum)
}

func calibrationValue(line string) int {

	val1 := -1
	val2 := -1

	for _, ch := range line {
		if unicode.IsDigit(ch) {
			intVal := int(ch - '0')
			if val1 == -1 {
				val1 = intVal
			}
			val2 = intVal
		}
	}

	lineVal, err := strconv.Atoi(fmt.Sprintf("%v%v", val1, val2))
	if err != nil {
		return 0
	}

	return lineVal
}
