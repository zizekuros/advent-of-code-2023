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
	var wonCopies []int

	for i, line := range lines {

		lineSum := 0

		line = strings.ReplaceAll(line, "\n", " \n")

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
			if part == 2 {
				foundCopies := findCopies(wonCopies, (i + 1))
				if lineSum > 0 {

					for j := 1; j <= lineSum; j++ {
						wonCopy := (i + 1) + j
						for k := 1; k <= (foundCopies + 1); k++ {
							wonCopies = append(wonCopies, wonCopy)
						}
					}

				}
			}
		}

		if lineSum > 0 && part == 1 {
			linePow := int(math.Pow(2, float64((lineSum - 1))))
			sum += linePow
		}
	}

	if part == 1 {
		fmt.Println(sum)
	} else if part == 2 {
		fmt.Println(len(wonCopies) + (len(lines) - 1))
	}

}

func checkWinningCard(list []string, card string) bool {
	for _, item := range list {
		if card == item {
			return true
		}
	}
	return false
}

func findCopies(list []int, card int) int {
	sum := 0
	for _, item := range list {
		if card == item {
			sum++
		}
	}
	return sum
}
