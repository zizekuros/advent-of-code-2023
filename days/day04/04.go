package day04

import (
	"fmt"
	"math"
	"os"
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
	sum := 0

	for _, line := range lines {

		lineSum := 0

		index := strings.Index(line, ":")
		trimmedText := line[index+1:]
		parts := strings.Split(trimmedText, "|")

		if len(parts) == 2 {
			winning, my := strings.Split(parts[0], " "), strings.Split(parts[1], " ")
			for _, card := range my {
				if card != " " && card != "" && checkWinningCard(winning, card) {
					lineSum += 1
				}
			}
		}

		if lineSum > 0 {
			linePow := int(math.Pow(2, float64((lineSum - 1))))
			sum += linePow
		}
	}

	fmt.Println(sum)

}

func checkWinningCard(list []string, card string) bool {
	for _, item := range list {
		if card == item {
			return true
		}
	}
	return false
}
