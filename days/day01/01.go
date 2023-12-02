package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Solve() {

	filePath := "./days/day01/input.txt"

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
