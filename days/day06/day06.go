package day06

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	lines := strings.Split(string(content), "\n")
	timeValues := parseValues(lines[0])
	distanceValues := parseValues(lines[1])

	if part == 2 {
		timeValues = concatenateArray(timeValues)
		distanceValues = concatenateArray(distanceValues)

	}

	result := calculateWaysToBeatRecord(timeValues, distanceValues)
	fmt.Println(result)
}

func parseValues(input string) []int {
	values := []int{}
	parts := strings.Fields(input)
	for _, part := range parts[1:] {
		value, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			os.Exit(1)
		}
		values = append(values, value)
	}
	return values
}

func calculateWaysToBeatRecord(timeValues []int, distanceValues []int) int {
	totalWays := 1
	for i := 0; i < len(timeValues); i++ {
		recordDistance := distanceValues[i]
		ways := 0
		for j := 1; j <= timeValues[i]; j++ {
			distance := j * (timeValues[i] - j)
			if distance > recordDistance {
				ways++
			}
		}
		totalWays *= ways
	}
	return totalWays
}

func concatenateArray(arr []int) []int {
	strNums := make([]string, len(arr))
	for i, num := range arr {
		strNums[i] = strconv.Itoa(num)
	}
	result, _ := strconv.Atoi(strings.Join(strNums, ""))
	return []int{result}
}
